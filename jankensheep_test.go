package jankensheep_test

import (
	"testing"

	. "github.com/genkami/jankensheep"
)

func TestSignGoString(t *testing.T) {
	assertSign := func(s Sign, expected string) {
		actual := s.GoString()
		if actual != expected {
			t.Errorf("Sign %d: expected %s but got %s", s, expected, actual)
		}
	}

	assertSign(Rock, "Rock")
	assertSign(Paper, "Paper")
	assertSign(Scissors, "Scissors")
}

func TestBeats(t *testing.T) {
	assertBeats := func(x, y Sign) {
		if !x.Beats(y) {
			t.Errorf("expected %#v to beat %#v, but it doesn't", x, y)
		}
	}
	assertNotBeats := func(x, y Sign) {
		if x.Beats(y) {
			t.Errorf("expected %#v to not beat %#v, but it does", x, y)
		}
	}

	assertNotBeats(Rock, Rock)
	assertNotBeats(Rock, Paper)
	assertBeats(Rock, Scissors)

	assertBeats(Paper, Rock)
	assertNotBeats(Paper, Paper)
	assertNotBeats(Paper, Scissors)

	assertNotBeats(Scissors, Rock)
	assertBeats(Scissors, Paper)
	assertNotBeats(Scissors, Scissors)
}

func TestPlay_TwoPlayers(t *testing.T) {
	assertResultIs(t, []Sign{Rock, Rock}, []int{})
	assertResultIs(t, []Sign{Rock, Paper}, []int{1})
	assertResultIs(t, []Sign{Rock, Scissors}, []int{0})

	assertResultIs(t, []Sign{Paper, Rock}, []int{0})
	assertResultIs(t, []Sign{Paper, Paper}, []int{})
	assertResultIs(t, []Sign{Paper, Scissors}, []int{1})

	assertResultIs(t, []Sign{Scissors, Rock}, []int{1})
	assertResultIs(t, []Sign{Scissors, Paper}, []int{0})
	assertResultIs(t, []Sign{Scissors, Scissors}, []int{})
}

func TestPlay_ManyPlayers_SinglePlayerWins(t *testing.T) {
	assertResultIs(t, []Sign{Rock, Scissors, Scissors}, []int{0})
	assertResultIs(t, []Sign{Rock, Paper, Rock}, []int{1})
	assertResultIs(t, []Sign{Paper, Paper, Scissors}, []int{2})

	assertResultIs(t, []Sign{Rock, Scissors, Scissors, Scissors}, []int{0})
	assertResultIs(t, []Sign{Rock, Scissors, Scissors, Scissors, Scissors}, []int{0})
	assertResultIs(t, []Sign{Rock, Scissors, Scissors, Scissors, Scissors, Scissors}, []int{0})
}

func TestPlay_ManyPlayers_MultiplePlayersWin(t *testing.T) {
	assertResultIs(t, []Sign{Rock, Rock, Scissors}, []int{0, 1})
	assertResultIs(t, []Sign{Rock, Paper, Paper}, []int{1, 2})
	assertResultIs(t, []Sign{Scissors, Paper, Scissors}, []int{0, 2})

	assertResultIs(t, []Sign{Scissors, Paper, Scissors, Scissors}, []int{0, 2, 3})
	assertResultIs(t, []Sign{Scissors, Paper, Scissors, Scissors, Scissors}, []int{0, 2, 3, 4})

	assertResultIs(t, []Sign{Scissors, Paper, Paper, Scissors, Scissors}, []int{0, 3, 4})
	assertResultIs(t, []Sign{Scissors, Paper, Paper, Paper, Scissors}, []int{0, 4})
}

func TestPlay_AllSignsAreTheSame(t *testing.T) {
	assertResultIs(t, []Sign{Rock, Rock}, []int{})
	assertResultIs(t, []Sign{Rock, Rock, Rock}, []int{})
	assertResultIs(t, []Sign{Rock, Rock, Rock, Rock}, []int{})
	assertResultIs(t, []Sign{Rock, Rock, Rock, Rock, Rock}, []int{})

	assertResultIs(t, []Sign{Paper, Paper}, []int{})
	assertResultIs(t, []Sign{Paper, Paper, Paper}, []int{})
	assertResultIs(t, []Sign{Paper, Paper, Paper, Paper}, []int{})
	assertResultIs(t, []Sign{Paper, Paper, Paper, Paper, Paper}, []int{})

	assertResultIs(t, []Sign{Scissors, Scissors}, []int{})
	assertResultIs(t, []Sign{Scissors, Scissors, Scissors}, []int{})
	assertResultIs(t, []Sign{Scissors, Scissors, Scissors, Scissors}, []int{})
	assertResultIs(t, []Sign{Scissors, Scissors, Scissors, Scissors, Scissors}, []int{})
}

func TestPlay_AllKindsOfSignsArePresent(t *testing.T) {
	assertResultIs(t, []Sign{Rock, Paper, Scissors}, []int{})
	assertResultIs(t, []Sign{Paper, Scissors, Rock}, []int{})
	assertResultIs(t, []Sign{Scissors, Rock, Paper}, []int{})

	assertResultIs(t, []Sign{Rock, Paper, Scissors, Rock}, []int{})
	assertResultIs(t, []Sign{Rock, Paper, Scissors, Paper}, []int{})
	assertResultIs(t, []Sign{Rock, Paper, Scissors, Scissors}, []int{})
}

func TestPlay_InsufficientNumberOfPlayers(t *testing.T) {
	assertPanics := func(name string, thunk func()) {
		panicked := false

		func() {
			defer func() {
				if err := recover(); err != nil {
					panicked = true
				}
			}()
			thunk()
		}()

		if !panicked {
			t.Errorf("%s: expected the function to panic, but it didn't", name)
		}
	}

	assertPanics("empty", func() { Play([]Sign{}) })
	assertPanics("single", func() { Play([]Sign{Rock}) })
}

func assertResultIs(t *testing.T, signs []Sign, expected []int) {
	actual := Play(signs)
	if !equalInts(expected, actual) {
		t.Errorf("expected Play(%#v) == %#v, but got %#v", signs, expected, actual)
	}
}

func equalInts(xs, ys []int) bool {
	if len(xs) != len(ys) {
		return false
	}
	for i, x := range xs {
		y := ys[i]
		if x != y {
			return false
		}
	}
	return true
}
