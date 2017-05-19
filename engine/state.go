package engine

import (
	"errors"
	"fmt"
	"math"
	"sync"
	"time"
)

type Engine struct {
	Neighborhoods []Neighborhood
	Players       map[string]*Player

	moveTo     map[string]*Position
	playerLock *sync.Mutex
}

func NewEngine() *Engine {
	return &Engine{
		Neighborhoods: NewNeighborhoods(),
		Players:       make(map[string]*Player, 0),
		playerLock:    new(sync.Mutex),
		moveTo:        make(map[string]*Position),
	}
}

func (e *Engine) Start() {
	go e.loop()
}

func (e *Engine) AddPlayer(player *Player) {
	e.playerLock.Lock()
	defer e.playerLock.Unlock()
	e.Players[player.Id] = player
}

func (e *Engine) Move(charId string, pos *Position) error {
	_, exists := e.moveTo[charId]
	if !exists {
		e.moveTo[charId] = pos
		return nil
	} else {
		return errors.New(fmt.Sprintf("Character %s already moving", charId))
	}
}

// private

func (e *Engine) loop() {
	for {
		// move players
		for k, _ := range e.moveTo {
			e.moveChar(k)
		}
		time.Sleep(time.Millisecond * 50)
	}
}

func (e *Engine) moveChar(charId string) {
	player := e.Players[charId]
	curX := player.CurrentPosition.X
	curY := player.CurrentPosition.Y
	toX := e.moveTo[charId].X
	toY := e.moveTo[charId].Y

	inc := float64(player.Speed) * (float64(player.Level) * .01)
	incX := inc * (diff(curX, toX) / diff(curY, toY))
	incY := inc * (diff(curY, toY) / diff(curX, toX))

	moved := false

	if diff(toX, curX) > .02 {
		if toX > curX {
			player.CurrentPosition.X = curX + incX
		} else {
			player.CurrentPosition.X = curX - incX
		}
		moved = true
	}

	if diff(toY, curY) > .02 {
		if toY > curY {
			player.CurrentPosition.Y = curY + incY
		} else {
			player.CurrentPosition.Y = curY - incY
		}
		moved = true
	}

	if !moved {
		player.CurrentPosition.X = math.Floor(player.CurrentPosition.X + 0.5)
		player.CurrentPosition.Y = math.Floor(player.CurrentPosition.Y + 0.5)
		delete(e.moveTo, charId)
	} else {
		fmt.Printf("Move %s: %v\n", player.Name, player.CurrentPosition)
	}
}

func diff(x1 float64, x2 float64) float64 {
	if x1 > x2 {
		return x1 - x2
	} else {
		return x2 - x1
	}
}
