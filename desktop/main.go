package main

import (
	"github.com/docopt/docopt-go"
	"github.com/hajimehoshi/ebiten"
	"github.com/reconquest/mrdig/game/engine"
)

var version = "[manual build]"

var usage = `mrdig - run desktop version of Mr. Dig.

Usage:
  mrdig -h | --help
  mrdig <width> <height>

Options:
  -h --help  Show this help.
`

var opts struct {
	Width  int
	Height int
}

func main() {
	args, err := docopt.ParseArgs(usage, nil, "mrdig "+version)
	if err != nil {
		panic(err)
	}

	err = args.Bind(&opts)
	if err != nil {
		panic(err)
	}

	engine := engine.New()

	err = ebiten.Run(engine.Update, opts.Width, opts.Height, 1, "Mr. Dig")
	if err != nil {
		panic(err)
	}
}
