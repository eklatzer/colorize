package main

import (
	"bufio"
	"flag"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"

	"colorize/colorprinter"
	"colorize/config"
)

var cfgFile string
var cfg config.Config

func main() {
	if cfgFile == "" {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalf("%v: failed to detect homedir, please provide configfile with flag --config", err)
		}
		cfgFile = filepath.Join(home, ".colorized.yaml")
	}

	var err error
	cfg, err = config.FromFile(cfgFile)
	if err != nil {
		cfg = config.DefaultConfig
	}
	err = cfg.Init()
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	var scanner *bufio.Scanner
	if isInputFromPipe() {
		scanner = bufio.NewScanner(bufio.NewReader(os.Stdin))
	} else {
		scanner, err = getScannerForFile(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
	}

	for scanner.Scan() {
		colorprinter.PrintLineColorized(scanner.Text(), cfg.Colors, cfg.Default.ColorValue)
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
