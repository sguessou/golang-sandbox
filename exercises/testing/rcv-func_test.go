//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests
package main

import "testing"

func newPlayer() Player {
	return Player{
		name:      "Joe",
		health:    100,
		maxHealth: 200,
		energy:    500,
		maxEnergy: 600,
	}
}

func TestEnergyFunctions(t *testing.T) {
	player := newPlayer()

	player.addEnergy(10)
	expected := uint(510)
	if player.energy != expected {
		t.Errorf("expected %v, got %v", expected, player.energy)
	}

	player.addEnergy(800)
	if player.energy != player.maxEnergy {
		t.Errorf("expected %v, got %v", player.maxEnergy, player.energy)
	}

	player.consumeEnergy(200)
	expected = 400
	if player.energy != expected {
		t.Errorf("expected %v, got %v", expected, player.energy)
	}

	player.consumeEnergy(1000)
	expected = 0
	if player.energy != expected {
		t.Errorf("expected %v, got %v", expected, player.energy)
	}

}
