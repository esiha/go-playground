package controller

type Newer interface {
	New() (Controller, error)
}
