package main

import (
	"os"

	"github.com/ejuju/poc-go-tty-art/internal/algolight"
	"github.com/ejuju/poc-go-tty-art/internal/gameoflife"
	"github.com/ejuju/poc-go-tty-art/internal/markode"
)

const (
	cmdGameOfLife = "game-of-life"
	cmdMarkode    = "markode"
	cmdAlgolight  = "algolight"
)

func main() {
	if len(os.Args) <= 1 {
		os.Args = append(os.Args, cmdGameOfLife)
	}
	var run func() error
	switch os.Args[1] {
	default:
		panic("unknown command")
	case cmdGameOfLife:
		run = gameoflife.Run
	case cmdMarkode:
		run = markode.Run
	case cmdAlgolight:
		run = algolight.Run
	}
	err := run()
	if err != nil {
		panic(err)
	}
}
