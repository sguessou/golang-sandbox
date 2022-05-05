//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once
package main

import "fmt"

type Player struct {
	name              string
	health, maxHealth uint
	energy, maxEnergy uint
}

func (p *Player) addHealth(amount uint) {
	p.health += amount
	if p.health > p.maxHealth {
		p.health = p.maxHealth
	}
	fmt.Println(p.name, "Add", amount, "health ->", p.health)
}

func (p *Player) applyDamage(amount uint) {
	if p.health-amount > p.health {
		p.health = 0
	} else {
		p.health -= amount
	}

	fmt.Println(p.name, "Damage", amount, "health ->", p.health)
}

func (p *Player) addEnergy(amount uint) {
	p.energy += amount
	if p.energy > p.maxEnergy {
		p.energy = p.maxEnergy
	}
	fmt.Println(p.name, "Add", amount, "energy ->", p.energy)
}

func (p *Player) consumeEnergy(amount uint) {
	if p.energy-amount > p.energy {
		p.energy = 0
	} else {
		p.energy -= amount
	}

	fmt.Println(p.name, "Consume", amount, "energy ->", p.energy)
}

func main() {
	player := Player{
		name:      "Joe",
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
	}

	player.applyDamage(99)
	player.addHealth(10)
	player.consumeEnergy(20)
	player.addEnergy(10)

	player.consumeEnergy(9999)
}
