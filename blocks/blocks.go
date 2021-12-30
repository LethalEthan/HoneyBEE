package blocks

type Block struct {
	Name         string
	currentState int
	minState     int
	maxState     int
}
