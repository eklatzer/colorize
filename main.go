package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"

	"colorize/config"
)

const defaultConfigFile = ".colorized.yaml"

var flagValues flags

func init() {
	home, _ := homedir.Dir()
	flag.StringVar(&flagValues.configFile, "config", filepath.Join(home, defaultConfigFile), "Path to the config file")
	flag.BoolVar(&flagValues.printDefaultConfig, "print-default-config", false, "Whether to print the default-config")
	flag.Parse()

	if flagValues.printDefaultConfig {
		out, err := yaml.Marshal(config.DefaultConfig)
		if err != nil {
			log.Fatalf("failed to marshal config: %v", err)
		}
		fmt.Println(string(out))
		os.Exit(0)
	}
}

func main() {
	cfg, err := config.FromFile(flagValues.configFile)
	if err != nil {
		cfg = config.DefaultConfig
	}

	var scanner *bufio.Scanner
	if isInputFromPipe() {
		scanner = bufio.NewScanner(bufio.NewReader(os.Stdin))
	} else {
		scanner, err = getScannerForFile(flag.Arg(0))
		if err != nil {
			log.Fatalf("failed to read file: %v\n", err)
		}
	}

	for scanner.Scan() {
		printLineColorized(scanner.Text(), cfg)
	}
}

func getScannerForFile(path string) (*bufio.Scanner, error) {
	readFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	return fileScanner, nil
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func printLineColorized(text string, cfg config.Config) {
	for _, rule := range cfg.Ruleset {
		if rule.MatchString(text) {
			rule.ColorScheme.PrintlnColored(text)
			return
		}
	}

	cfg.Default.PrintlnColored(text)
}
