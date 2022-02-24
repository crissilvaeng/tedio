package misc

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/crissilvaeng/tedio/internal/codes"
)

var (
	fgregex  = regexp.MustCompile(`^\033\[(3[0-7]|9[0-7])m$`)
	bgregex  = regexp.MustCompile(`^\033\[(4[0-7]|10[0-7])m$`)
	styregex = regexp.MustCompile(`^\033\[([0-7])m$`)
)

type RuneColorizer struct {
	format string
}

func (c *RuneColorizer) Colorize(r rune) string {
	return fmt.Sprintf(c.format, unicode.ToUpper(r))
}

type RuneColorizerBuilder struct {
	foreground string
	background string
	styles     []string
}

func NewRuneColorizerBuilder() *RuneColorizerBuilder {
	return &RuneColorizerBuilder{styles: make([]string, 0)}
}

func (b *RuneColorizerBuilder) Foreground(foreground string) *RuneColorizerBuilder {
	if fgregex.MatchString(foreground) && len(b.foreground) == 0 {
		b.foreground = foreground
	}
	return b
}

func (b *RuneColorizerBuilder) Background(background string) *RuneColorizerBuilder {
	if bgregex.MatchString(background) && len(b.background) == 0 {
		b.background = background
	}
	return b
}

func (b *RuneColorizerBuilder) Styles(styles ...string) *RuneColorizerBuilder {
	for _, style := range styles {
		if styregex.MatchString(style) {
			b.styles = append(b.styles, style)
		}
	}
	return b
}

func (b *RuneColorizerBuilder) ToBuilder() *RuneColorizerBuilder {
	return &RuneColorizerBuilder{
		foreground: b.foreground,
		background: b.background,
		styles:     b.styles,
	}
}

func (b *RuneColorizerBuilder) Build() *RuneColorizer {
	sb := &strings.Builder{}

	sb.WriteString(b.foreground)
	sb.WriteString(b.background)

	for _, style := range b.styles {
		sb.WriteString(style)
	}

	sb.WriteString(" %c ")
	sb.WriteString(codes.Reset)

	return &RuneColorizer{
		format: sb.String(),
	}
}
