package trailforge

import "fmt"

type Difficulty string

const (
	Easy Difficulty = "Easy"
	Medium Difficulty = "Medium"
	Hard Difficulty = "Hard"
	Beast Difficulty = "Beast"
)

type Hike struct {
	Name string
	Location string
	DistanceMi float64
	ElevGainFt int
	Difficulty Difficulty
	Solo bool
}

func (h Hike) String() string {
	return fmt.Sprintf("%s (%.1f mi, %d ft gain, %s)",
		h.Name, h.DistanceMi, h.ElevGainFt, h.Difficulty)
}
