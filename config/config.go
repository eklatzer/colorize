package config

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/fatih/color"
	"gopkg.in/yaml.v3"

	"colorize/colorhelper"
)

var DefaultConfig = Config{
	Colors: []*ColorForLevel{
		{Expression: "(?i).*fatal.*", Color: "FgHiRed"},
		{Expression: "(?i).*error.*", Color: "FgRed"},
		{Expression: "(?i).*warn.*", Color: "FgYellow"},
		{Expression: "(?i).*info.*", Color: "FgBlue"},
		{Expression: "(?i).*debug.*", Color: "FgGreen"},
		{Expression: "(?i).*trace.*", Color: "FgHiCyan"},
	},
	Default: &ColorForLevel{Color: "FgWhite"},
}

func FromFile(path string) (Config, error) {
	var config Config
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	return config, err
}

type Config struct {
	Colors  []*ColorForLevel
	Default *ColorForLevel
}

type ColorForLevel struct {
	Expression string
	Color      string
	ColorValue *color.Color
	Regex      *regexp.Regexp
}

func (c *Config) Init() error {
	for _, colorizing := range c.Colors {
		val, err := colorhelper.StringToColor(colorizing.Color)
		if err != nil {
			return err
		}
		colorizing.ColorValue = color.New(val)

		colorizing.Regex, err = regexp.Compile(colorizing.Expression)
		if err != nil {
			return fmt.Errorf("failed to compile regex: %v", err)
		}
	}

	val, err := colorhelper.StringToColor(c.Default.Color)
	if err != nil {
		return err
	}
	c.Default.ColorValue = color.New(val)

	return nil
}
