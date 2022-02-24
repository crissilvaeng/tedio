package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/crissilvaeng/tedio/internal/codes"
	"github.com/crissilvaeng/tedio/internal/misc"
)

func main() {
	builder := misc.NewRuneColorizerBuilder().
		Foreground(codes.FgBrightWhite).Styles(codes.Bold)

	green := builder.ToBuilder().Background(codes.BgGreen).Build()
	yellow := builder.ToBuilder().Background(codes.BgYellow).Build()
	gray := builder.ToBuilder().Background(codes.BgBlack).Build()

	colors := []*misc.RuneColorizer{green, yellow, gray}

	seed := rand.NewSource(time.Now().Unix())
	rand := rand.New(seed)
	for _, r := range "paralelepipedo" {
		color := colors[rand.Intn(len(colors))]
		fmt.Print(color.Colorize(r))
	}
	fmt.Println()
}
