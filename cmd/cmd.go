package cmd

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"colorize/colorprinter"
	"colorize/config"
)

var cfgFile string
var cfg config.Config

func init() {
	cobra.OnInitialize(onInitCobra)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file with path and extension (default: $HOME/.colorized.yaml)")
}

var RootCmd = &cobra.Command{
	Use:   "",
	Short: "Colorize is a command to print logs colorized",
	Run: func(cmd *cobra.Command, args []string) {
		var scanner *bufio.Scanner
		if isInputFromPipe() {
			scanner = bufio.NewScanner(bufio.NewReader(os.Stdin))
		} else {
			err := cobra.ExactArgs(1)(cmd, args)
			if err != nil {
				cmd.Help()
				return
			}
			scanner, err = getScannerForFile(args[0])
			if err != nil {
				log.Fatal(err)
			}
		}

		for scanner.Scan() {
			colorprinter.PrintLineColorized(scanner.Text(), cfg.Colorizing.Colors, cfg.Colorizing.Default.ColorValue)
		}
	},
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

func onInitCobra() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalf("%v: failed to detect homedir, please provide configfile with flag --config", err)
		}
		viper.SetConfigFile(filepath.Join(home, ".colorized.yaml"))
	}
	err := readConfigInto(&cfg)
	if err != nil {
		log.Debug("failed to load config (%v), using default values", err)
		cfg = config.DefaultConfig
	}
	err = cfg.Init()
	if err != nil {
		log.Fatalf("failed to parse color: %v", err)
	}
}

func readConfigInto(config interface{}) error {
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&config)
	return err
}
