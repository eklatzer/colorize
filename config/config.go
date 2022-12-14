package config

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"gopkg.in/yaml.v3"
)

var DefaultConfig = Config{
	Colors: []*ColorForLevel{
		{Expression: "(?i).*fatal.*", Color: "#FF0000"},
		{Expression: "(?i).*error.*", Color: "#FFA500"},
		{Expression: "(?i).*warn.*", Color: "#FFFF00"},
	},
	Default: &ColorForLevel{Color: "#ADD8E6"},
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
	Regex      *regexp.Regexp
}

func (c *Config) CompileRegexes() error {
	for _, colorizing := range c.Colors {
		regex, err := regexp.Compile(colorizing.Expression)
		if err != nil {
			return fmt.Errorf("failed to compile regex: %v", err)
		}
		colorizing.Regex = regex
	}
	return nil
}
