package units

import (
	"designer/battle/interval/core"
	"designer/battle/interval/skills"
)

type Mage struct{ BaseUnit }

func NewMage(id, team string, ai core.AIStrategy) *Mage {
	m := &Mage{}
	m.id, m.name, m.team = id, "Mage("+id+")", team
	m.hp, m.max = 40, 40
	m.skill = skills.Fireball{}
	m.ai = ai
	return m
}
