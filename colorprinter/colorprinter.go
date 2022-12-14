package colorprinter

import (
	"github.com/fatih/color"

	"colorize/config"
)

func PrintLineColorized(text string, colorMappings []*config.ColorForLevel, defaultColor *color.Color) {
	for _, colorizing := range colorMappings {
		if colorizing.Regex.MatchString(text) {
			colorizing.ColorValue.Println(text)
			return
		}
	}

	defaultColor.Println(text)
}
