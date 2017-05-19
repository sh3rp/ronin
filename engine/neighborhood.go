package engine

var NEIGHBORHOOD_SIZE_X = 5
var NEIGHBORHOOD_SIZE_Y = 5

var NEIGHBORHOOD_COUNT_X = 5
var NEIGHBORHOOD_COUNT_Y = 5

var names = []string{
	"Waldorf",
	"Everleaf",
	"Shinetop",
	"Pants",
	"Shoes",
	"Shirt",
	"New Campbell",
	"Fairfield",
	"Bellenew",
}

type Neighborhood struct {
	Name string
	X    int
	Y    int
}

func NewNeighborhoods() []Neighborhood {
	var neighborhoods []Neighborhood

	for idx, name := range names {
		var y = 0
		if idx > 0 {
			y = NEIGHBORHOOD_COUNT_Y / idx
		}
		neighborhoods = append(neighborhoods, Neighborhood{
			Name: name,
			X:    NEIGHBORHOOD_COUNT_X + idx,
			Y:    y,
		})
	}
	return neighborhoods
}

func (n *Neighborhood) Center() *Position {
	return &Position{X: float64(n.X + (NEIGHBORHOOD_SIZE_X / 2)), Y: float64(n.Y + (NEIGHBORHOOD_SIZE_Y / 2))}
}
