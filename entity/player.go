package entity

import (
	"fmt"
	"strconv"
)

type Player struct {
	entity
	maxHealth   int
	healCounter int8
}

func NewPlayer() *Player {
	entity := newEntity()
	return &Player{
		entity:      *entity,
		maxHealth:   entity.health,
		healCounter: 0,
	}
}

func (pl *Player) Heal() string {
	if pl.healCounter == 4 {
		return "Исчерпан лимит"
	}
	if pl.health == pl.maxHealth {
		return "Лечение не требуется"
	}
	if float32(pl.health)+float32(pl.maxHealth)*0.3 > float32(pl.maxHealth) {
		pl.health = pl.maxHealth
	} else {
		pl.health = pl.health + int(float32(pl.maxHealth)*0.3)
	}
	pl.healCounter++
	return "Выполено лечение. Текущее состояние здоровья: " + strconv.FormatUint(uint64(pl.health), 10)
}

func (pl Player) Attack(m *Monster) string {
	msg := attack(pl.entity, &m.entity)
	if msg == "Успешная атака" {
		fmt.Println("Совершена успешная атака")
	}
	return msg
}
