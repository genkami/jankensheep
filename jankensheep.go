package jankensheep

import "fmt"

// Sign is a hand sign in rock-paper-scissors game.
type Sign int

// All hand signs.
// Note that we internally use modulo calculation to implement the game, so they are defined explcitly instead of using `iota`.
const (
	Rock     Sign = 0
	Paper    Sign = 1
	Scissors Sign = 2
)

func (s Sign) GoString() string {
	switch s {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	default:
		panic(errInvalidSign(s))
	}
}

func errInvalidSign(s Sign) error {
	return fmt.Errorf("jankensheep: invalid sign: %d", s)
}
