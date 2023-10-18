package entity

import (
	"math/rand"
	"time"
)

type Monster struct {
	entity
}

func NewMonster() *Monster {
	rand.Seed(time.Now().UnixNano())
	entity := newEntity()
	return &Monster{entity: *entity}
}

func (m Monster) Attack(pl *Player) string {
	return attack(m.entity, &pl.entity)
}
