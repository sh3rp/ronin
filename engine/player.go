package engine

import "github.com/twinj/uuid"

type Player struct {
	Id              string
	Name            string
	CurrentPosition *Position

	Level int // 1-50
	Speed int // 1-10, based on character roll
}

type Position struct {
	X float64
	Y float64
}

func NewPlayer(name string, speed int) *Player {
	return &Player{
		Id:    uuid.NewV4().String(),
		Name:  name,
		Level: 1,
		Speed: speed,
		CurrentPosition: &Position{
			X: 1,
			Y: 1,
		},
	}
}
