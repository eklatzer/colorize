package colorhelper

import (
	"fmt"

	"github.com/fatih/color"
)

type unknownColorError struct {
	color string
}

func (u unknownColorError) Error() string {
	return fmt.Sprintf("unknown color: %q", u.color)
}

func StringToColor(in string) (color.Attribute, error) {
	color, exisits := stringToColor[in]
	if !exisits {
		return color, unknownColorError{in}
	}

	return color, nil
}

var stringToColor = map[string]color.Attribute{
	"FgBlack":     color.FgBlack,
	"FgRed":       color.FgRed,
	"FgGreen":     color.FgGreen,
	"FgYellow":    color.FgYellow,
	"FgBlue":      color.FgBlue,
	"FgMagenta":   color.FgMagenta,
	"FgCyan":      color.FgCyan,
	"FgWhite":     color.FgWhite,
	"FgHiBlack":   color.FgHiBlack,
	"FgHiRed":     color.FgHiRed,
	"FgHiGreen":   color.FgHiGreen,
	"FgHiYellow":  color.FgHiYellow,
	"FgHiBlue":    color.FgHiBlue,
	"FgHiMagenta": color.FgHiMagenta,
	"FgHiCyan":    color.FgHiCyan,
	"FgHiWhite":   color.FgHiWhite,
}
