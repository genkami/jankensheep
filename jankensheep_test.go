package jankensheep_test

import (
	"testing"

	"github.com/genkami/jankensheep"
)

func TestSignGoString(t *testing.T) {
	assertSign := func(s jankensheep.Sign, expected string) {
		actual := s.GoString()
		if actual != expected {
			t.Errorf("Sign %d: expected %s but got %s", s, expected, actual)
		}
	}

	assertSign(jankensheep.Rock, "Rock")
	assertSign(jankensheep.Paper, "Paper")
	assertSign(jankensheep.Scissors, "Scissors")
}
