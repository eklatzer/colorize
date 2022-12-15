package config

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/charmbracelet/lipgloss"
	"gopkg.in/yaml.v3"

	"colorize/config/regex"
)

var DefaultConfig = Config{
	Ruleset: []Rule{
		{Expression: regex.Expression{Regexp: regexp.MustCompile("(?i).*fatal.*")}, ColorScheme: ColorScheme{Foreground: "#FF0000"}},
		{Expression: regex.Expression{Regexp: regexp.MustCompile("(?i).*error.*")}, ColorScheme: ColorScheme{Foreground: "#FFA500"}},
		{Expression: regex.Expression{Regexp: regexp.MustCompile("(?i).*warn.*")}, ColorScheme: ColorScheme{Foreground: "#FFFF00"}},
	},
	Default: ColorScheme{
		Foreground: "#ADD8E6",
	},
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
	Ruleset []Rule
	Default ColorScheme
}

type Rule struct {
	Expression  regex.Expression
	ColorScheme ColorScheme
}

func (r *Rule) MatchString(m string) bool {
	return r.Expression.Regexp.MatchString(m)
}

type ColorScheme struct {
	Foreground string
	Background string
}

func (c *ColorScheme) PrintlnColored(text string) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color(c.Foreground)).Background(lipgloss.Color(c.Background))
	fmt.Println(style.Render(text))
}
