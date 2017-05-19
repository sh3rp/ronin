package engine

type Contract struct {
	Target       *NPC
	Price        int
	Neighborhood *Neighborhood
}

type NPC struct {
	Name  string
	Level int
	Speed int
}
