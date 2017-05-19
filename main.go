package main

import (
	"math/rand"
	"time"

	"github.com/sh3rp/ronin/engine"
)

func main() {
	e := engine.NewEngine()
	e.Start()
	player := engine.NewPlayer("Pants", 1)
	e.AddPlayer(player)
	for {
		randNeighborhood := e.Neighborhoods[rand.Intn(len(e.Neighborhoods))]
		e.Move(player.Id, randNeighborhood.Center())
		time.Sleep(time.Second * 1)
	}
}
