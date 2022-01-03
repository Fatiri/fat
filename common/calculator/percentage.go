package calculator

import (
	"math"
)

type PayloadPercentage struct {
	OldNumber float64
	NewNumber float64
}

func (c *PayloadPercentage) CalculatePercenatgeIncreaseDecrease() (float64, float64, error) {
	increase := float64(c.OldNumber-c.NewNumber) / float64(c.NewNumber) * 100
	decrease := float64(c.OldNumber-c.NewNumber) / float64(c.OldNumber) * 100
	return increase, decrease, nil
}

func (c *PayloadPercentage) PercentageChange() (delta float64) {
	diff := float64(c.NewNumber - c.OldNumber)
	delta = (diff / float64(c.OldNumber)) * 100

	if math.IsNaN(delta) {
		delta = 0
	}
	return
}
