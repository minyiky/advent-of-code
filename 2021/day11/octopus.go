package day11

type octopus struct {
	energy  int
	flashed bool
}

func (o *octopus) increaseEnergy() {
	o.energy++
}

func (o *octopus) reset() {
	o.energy = 0
	o.flashed = false
}

func (o *octopus) flash() {
	o.flashed = true
}
