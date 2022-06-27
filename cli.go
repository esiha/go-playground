package main

import (
	"errors"
	"fmt"
	"go-playground/domain/controller"
	"go-playground/domain/direction"
	"go-playground/domain/mower"
	"go-playground/driven/file"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		failed(errors.New("invalid number of arguments"))
	}

	if newer, err := file.NewControllerNewer(os.Args[1]); err != nil {
		failed(err)
	} else if positions, err := controller.Run(&newer); err != nil {
		failed(err)
	} else {
		printPositions(positions)
	}
}

func failed(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}

func printPositions(positions []mower.Position) {
	for _, p := range positions {
		fmt.Printf("%d %d %v\n", p.Point().X(), p.Point().Y(), news(p.Direction()))
	}
}

func news(d direction.Direction) string {
	switch d {
	case direction.North:
		return "N"
	case direction.East:
		return "E"
	case direction.South:
		return "S"
	case direction.West:
		return "W"
	default:
		panic("unhandled position")
	}
}
