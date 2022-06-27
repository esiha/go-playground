package controller

import "go-playground/domain/mower"

func Run(newer Newer) ([]mower.Position, error) {
	if controller, err := newer.New(); err != nil {
		return nil, err
	} else {
		controller.Run()
		return controller.MowersPositions(), nil
	}
}
