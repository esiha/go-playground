package main

import (
	"bufio"
	"fmt"
	"go-playground/domain/controller"
	"go-playground/domain/direction"
	"go-playground/domain/instruction"
	"go-playground/domain/lawn"
	"go-playground/domain/mower"
	"go-playground/domain/point"
	"io"
	"os"
	"strconv"
	"strings"
)

type ConfigReader struct {
	reader io.Reader
}

func NewConfigReader(path string) (ConfigReader, error) {
	if f, err := os.Open(path); err != nil {
		return ConfigReader{}, err
	} else {
		return newFromReader(f), nil
	}
}

func newFromReader(reader io.Reader) ConfigReader {
	return ConfigReader{reader}
}

func (n *ConfigReader) ReadConfig() (controller.Controller, error) {
	scanner := bufio.NewScanner(n.reader)
	if !scanner.Scan() {
		return controller.Controller{}, fmt.Errorf("empty file")
	} else if l, err := parseLawn(scanner.Text()); err != nil {
		return controller.Controller{}, err
	} else if mowers, err := parseMowers(scanner); err != nil {
		return controller.Controller{}, err
	} else {
		return controller.New(mowers, l), nil
	}
}

func parseLawn(s string) (lawn.Lawn, error) {
	if topRightCorner, err := parsePoint(s); err != nil {
		return lawn.Lawn{}, err
	} else {
		return lawn.Rectangular(topRightCorner), nil
	}
}

func parseMowers(scanner *bufio.Scanner) ([]mower.Mower, error) {
	mowers := make([]mower.Mower, 0)
	for scanner.Scan() {
		startLine := scanner.Text()
		if !scanner.Scan() {
			return nil, fmt.Errorf("missing instructions line for mower")
		} else if m, err := parseMower(startLine, scanner.Text()); err != nil {
			return nil, err
		} else {
			mowers = append(mowers, m)
		}
	}
	return mowers, nil
}

func parseMower(startLine string, instructionsLine string) (mower.Mower, error) {
	if p, err := parsePosition(startLine); err != nil {
		return mower.Mower{}, err
	} else if instructions, err := parseInstructions(instructionsLine); err != nil {
		return mower.Mower{}, err
	} else {
		return mower.New(p, instructions), nil
	}
}

func parseInstructions(ss string) ([]instruction.Instruction, error) {
	instructions := make([]instruction.Instruction, len(ss))
	for i, s := range strings.Split(ss, "") {
		if inst, err := parseInstruction(s); err != nil {
			return nil, err
		} else {
			instructions[i] = inst
		}
	}
	return instructions, nil
}

func parseInstruction(s string) (instruction.Instruction, error) {
	switch s {
	case "A":
		return instruction.Advance, nil
	case "D":
		return instruction.TurnRight, nil
	case "G":
		return instruction.TurnLeft, nil
	default:
		return -1, fmt.Errorf("unknown instruction '%v'", s)
	}
}

func parsePosition(s string) (mower.Position, error) {
	if t := strings.Split(s, " "); len(t) != 3 {
		return mower.Position{}, fmt.Errorf("expected 3 parts, got %v", len(t))
	} else if x, err := strconv.Atoi(t[0]); err != nil {
		return mower.Position{}, err
	} else if y, err := strconv.Atoi(t[1]); err != nil {
		return mower.Position{}, err
	} else if d, err := parseDirection(t[2]); err != nil {
		return mower.Position{}, err
	} else {
		return mower.NewPosition(point.New(x, y), d), nil
	}
}

func parseDirection(s string) (direction.Direction, error) {
	switch s {
	case "N":
		return direction.North, nil
	case "E":
		return direction.East, nil
	case "W":
		return direction.West, nil
	case "S":
		return direction.South, nil
	default:
		return -1, fmt.Errorf("unknown direction '%v'", s)
	}
}

func parsePoint(s string) (point.Point, error) {
	if t := strings.Split(s, " "); len(t) != 2 {
		return point.Point{}, fmt.Errorf("expected 2 parts, got %v", len(t))
	} else if x, err := strconv.Atoi(t[0]); err != nil {
		return point.Point{}, err
	} else if y, err := strconv.Atoi(t[1]); err != nil {
		return point.Point{}, err
	} else {
		return point.New(x, y), nil
	}
}
