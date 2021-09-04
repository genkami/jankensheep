package jankensheep

import "fmt"

// Sign is a hand sign in rock-paper-scissors game.
type Sign int

// All hand signs.
// Note that we internally use modulo calculation to implement the game rule, so they are defined explcitly instead of using `iota`.
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

// Beats returns `true` if `s` beats `t`; otherwise it returns `false`.
func (s Sign) Beats(t Sign) bool {
	return ((t + 1) % 3) == s
}

// Play returns indices of winners of this game in `O(N)` time complexity.
// It panics when `len(signs) < 2`.
func Play(signs []Sign) []int {
	if len(signs) < 2 {
		panic("jankensheep: at least two players are needed to play the game")
	}

	// First, we decide which sign wins the game.
	var winningSign = Sign(0)
	someoneWins := false
outer:
	for _, winningSign = range []Sign{Rock, Paper, Scissors} {
		someoneBeated := false
		for _, s := range signs {
			if s.Beats(winningSign) {
				continue outer
			} else if winningSign.Beats(s) {
				someoneBeated = true
			}
		}
		if someoneBeated {
			someoneWins = true
			break
		}
	}
	if !someoneWins {
		return []int{}
	}

	// Then, we decide which players are the winners.
	winnerIndices := make([]int, 0, len(signs))
	for i, s := range signs {
		if s == winningSign {
			winnerIndices = append(winnerIndices, i)
		}
	}
	return winnerIndices
}

func errInvalidSign(s Sign) error {
	return fmt.Errorf("jankensheep: invalid sign: %d", s)
}
