package internal

import (
	"fmt"
	"strings"
)

type Color string

func (c Color) String() string {
	return string(c)
}

const (
	white Color = "\x1b[0;37m"

	BgGreen  Color = "\033[42m"
	BgYellow Color = "\033[43m"
	BgGray   Color = "\033[47m"

	reset Color = "\033[0m"
)

type colorizer struct {
	options map[Color]bool
}

func NewColorizer() *colorizer {
	return &colorizer{
		options: map[Color]bool{BgGreen: true, BgYellow: true, BgGray: true},
	}
}

func (c *colorizer) Colorize(text string, bgcolor Color) string {
	if strings.TrimSpace(text) == "" {
		return text
	}
	if !c.options[bgcolor] {
		return text
	}
	return fmt.Sprintf("%s%s%s%s", bgcolor, white, text, reset)
}
