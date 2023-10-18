package main

import (
	"fmt"
	ent "gogame/entity"
)

func main() {
	var pl *ent.Player = ent.NewPlayer()
	var mn *ent.Monster = ent.NewMonster()
	for {
		result1 := pl.Attack(mn)
		result2 := mn.Attack(pl)
		if pl.GetHP() < int(float32(pl.GetHP())*0.5) {
			pl.Heal()
		}
		if result1 == "Противник мертв" {
			fmt.Println("Победа игрока")
			break
		}
		if result2 == "Противник мертв" {
			fmt.Println("Победа монстра")
			break
		}
	}
}
