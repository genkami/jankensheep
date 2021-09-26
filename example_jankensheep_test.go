package jankensheep_test

import (
	"fmt"

	"github.com/genkami/jankensheep"
)

func ExampleSign_Beats() {
	myHand := jankensheep.Paper
	opponentsHand := jankensheep.Rock
	if myHand.Beats(opponentsHand) {
		fmt.Println("win")
	}
	// Output: win
}

func ExamplePlay() {
	players := []string{"Watame", "Miko", "Roboco"}
	signs := []jankensheep.Sign{jankensheep.Rock, jankensheep.Scissors, jankensheep.Rock}
	winners := jankensheep.Play(signs)
	fmt.Println("Winners:")
	for _, i := range winners {
		fmt.Println(players[i])
	}
	// Output: Winners:
	// Watame
	// Roboco
}

func ExamplePlay_Draw() {
	signs := []jankensheep.Sign{jankensheep.Rock, jankensheep.Paper, jankensheep.Scissors}
	winners := jankensheep.Play(signs)
	if len(winners) == 0 {
		fmt.Println("draw")
	}
	// Output: draw
}
