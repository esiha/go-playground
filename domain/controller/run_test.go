package controller

import (
	"errors"
	"go-playground/domain/direction"
	"go-playground/domain/instruction"
	"go-playground/domain/lawn"
	"go-playground/domain/mower"
	"go-playground/domain/point"
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name    string
		newer   Newer
		want    []mower.Position
		wantErr bool
	}{
		{
			"Failure returns error",
			&failingNewer{},
			nil,
			true,
		},
		{
			"Success returns positions",
			&succeedingNewer{},
			[]mower.Position{mower.NewPosition(point.New(3, 2), direction.East)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Run(tt.newer)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type failingNewer struct{}

func (f *failingNewer) New() (Controller, error) {
	return Controller{}, errors.New("failed")
}

type succeedingNewer struct{}

func (s *succeedingNewer) New() (Controller, error) {
	return New(
		[]mower.Mower{
			mower.New(
				mower.NewPosition(point.New(2, 2), direction.East),
				[]instruction.Instruction{instruction.Advance},
			),
		},
		lawn.Rectangular(point.New(3, 3)),
	), nil
}
