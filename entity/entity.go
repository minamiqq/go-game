package entity

import (
	"math/rand"
	"strconv"
	"time"
)

type entity struct {
	attack     int
	damage     [2]int
	protection int
	health     int
}

func newEntity() *entity {
	rand.Seed(time.Now().UnixNano())
	return &entity{
		attack:     rand.Intn(30) + 1,
		damage:     [2]int{(rand.Intn(2) + 1), rand.Intn(15-3) + 3},
		protection: rand.Intn(30) + 1,
		health:     100,
	}
}

func (e entity) GetHP() int {
	return e.health
}

func attack(e1 entity, e2 *entity) string {
	if e2.health <= 0 {
		return "Противник мертв"
	}
	var mod int
	if e1.attack < e2.protection {
		mod = 1
	} else {
		mod = int(e1.attack-e2.protection) + 1
	}
	rand.Seed(time.Now().UnixNano())
	var num int
	for i := 0; i <= mod; i++ {
		num = rand.Intn(6) + 1
		if num == 5 || num == 6 {
			damage := rand.Intn(e1.damage[1]-e1.damage[0]) + e1.damage[0]
			e2.health = e2.health - damage
			return "Совершена успешная атака, урон составил " + strconv.Itoa(damage) + "\nОстаток здоровья " + strconv.Itoa(e2.health)
		}
	}
	return "Неудачная атака"
}
