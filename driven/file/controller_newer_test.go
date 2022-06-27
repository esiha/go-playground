package file

import (
	"go-playground/domain/controller"
	"go-playground/domain/direction"
	"go-playground/domain/instruction"
	"go-playground/domain/lawn"
	"go-playground/domain/mower"
	"go-playground/domain/point"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestNewer_New(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    controller.Controller
		wantErr bool
	}{
		{
			"empty file",
			[]string{},
			controller.Controller{},
			true,
		},
		{
			"invalid lawn X",
			[]string{
				"x 5",
			},
			controller.Controller{},
			true,
		},
		{
			"invalid lawn Y",
			[]string{
				"5 x",
			},
			controller.Controller{},
			true,
		},
		{
			"invalid mower x",
			[]string{
				"5 5",
				"x 3 N",
			},
			controller.Controller{},
			true,
		},
		{
			"invalid mower y",
			[]string{
				"5 5",
				"3 x N",
			},
			controller.Controller{},
			true,
		},
		{
			"invalid mower direction",
			[]string{
				"5 5",
				"3 3 X",
			},
			controller.Controller{},
			true,
		},
		{
			"incomplete mower",
			[]string{
				"5 5",
				"3 3 N",
			},
			controller.Controller{},
			true,
		},
		{
			"invalid mower instruction",
			[]string{
				"5 5",
				"3 3 N",
				"X",
			},
			controller.Controller{},
			true,
		},
		{
			"valid configuration",
			[]string{
				"4 5",
				"3 3 N",
				"ADG",
				"1 1 E",
				"GG",
				"4 4 W",
				"DD",
				"2 3 S",
				"AAA",
			},
			controller.New(
				[]mower.Mower{
					mower.New(
						mower.NewPosition(point.New(3, 3), direction.North),
						[]instruction.Instruction{instruction.Advance, instruction.TurnRight, instruction.TurnLeft},
					),
					mower.New(
						mower.NewPosition(point.New(1, 1), direction.East),
						[]instruction.Instruction{instruction.TurnLeft, instruction.TurnLeft},
					),
					mower.New(
						mower.NewPosition(point.New(4, 4), direction.West),
						[]instruction.Instruction{instruction.TurnRight, instruction.TurnRight},
					),
					mower.New(
						mower.NewPosition(point.New(2, 3), direction.South),
						[]instruction.Instruction{instruction.Advance, instruction.Advance, instruction.Advance},
					),
				},
				lawn.Rectangular(point.New(4, 5))),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newer := newFromReader(reader(tt.input))

			got, err := newer.New()
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func reader(lines []string) io.Reader {
	return strings.NewReader(strings.Join(lines, "\n"))
}
