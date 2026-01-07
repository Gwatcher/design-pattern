package units

import (
	"designer/battle/interval/core"
	"designer/battle/interval/skills"
)

type Warrior struct{ BaseUnit }

func NewWarrior(id, team string, ai core.AIStrategy) *Warrior {
	w := &Warrior{}
	w.id, w.name, w.team = id, "Warrior("+id+")", team
	w.hp, w.max = 60, 60
	w.skill = skills.Slash{}
	w.ai = ai
	return w
}
