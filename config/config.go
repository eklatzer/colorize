package config

import (
	"regexp"

	"github.com/fatih/color"

	"colorize/colorhelper"
)

var DefaultConfig = Config{
	Colorizing: Colorizing{
		Colors: []*ColorForLevel{
			{Expression: "(?i).*fatal.*", Color: "FgHiRed"},
			{Expression: "(?i).*error.*", Color: "FgRed"},
			{Expression: "(?i).*warn.*", Color: "FgYellow"},
			{Expression: "(?i).*info.*", Color: "FgBlue"},
			{Expression: "(?i).*debug.*", Color: "FgGreen"},
			{Expression: "(?i).*trace.*", Color: "FgHiCyan"},
		},
		Default: &ColorForLevel{Color: "FgWhite"},
	},
}

type Config struct {
	Colorizing Colorizing
}

type Colorizing struct {
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
	for _, colorizing := range c.Colorizing.Colors {
		val, err := colorhelper.StringToColor(colorizing.Color)
		if err != nil {
			return err
		}
		colorizing.ColorValue = color.New(val)

		colorizing.Regex, err = regexp.Compile(colorizing.Expression)
		if err != nil {
			return err
		}
	}

	val, err := colorhelper.StringToColor(c.Colorizing.Default.Color)
	if err != nil {
		return err
	}
	c.Colorizing.Default.ColorValue = color.New(val)

	return nil
}
