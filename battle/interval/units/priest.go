package units

import (
	"designer/battle/interval/core"
	"designer/battle/interval/skills"
)

type Priest struct{ BaseUnit }

func NewPriest(id, team string, ai core.AIStrategy) *Priest {
	p := &Priest{}
	p.id, p.name, p.team = id, "Priest("+id+")", team
	p.hp, p.max = 45, 45
	p.skill = skills.Heal{}
	p.ai = ai
	return p
}
