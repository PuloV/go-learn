package main

const (
	START int = iota
	UP
	LEFT
	DOWN
	RIGHT
)

type DragonFractal struct {
	history       []int
	step          int
	iterationStep int
}

func (d *DragonFractal) Next() string {
	if d.step == 0 {
		d.history = make([]int, 1)
		d.history = append(d.history, START)
		d.step = 1
	}
	if d.step/2 <= d.iterationStep {
		d.iterationStep = 0
	}
	direction := d.history[d.step-(d.iterationStep*2)] + 1
	if direction == 5 {
		direction = 1
	}
	d.step += 1
	d.iterationStep += 1
	d.history = append(d.history, direction)
	return d.directionAsString(direction)
}

func (d *DragonFractal) directionAsString(direction int) string {
	switch direction {
	case UP:
		return "up"
	case LEFT:
		return "left"
	case DOWN:
		return "down"
	case RIGHT:
		return "right"
	}
	return ""
}
