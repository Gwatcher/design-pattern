package units

import "designer/battle/interval/core"

type BaseUnit struct {
	id   string
	name string
	team string
	hp   int
	max  int

	skill core.Skill
	ai    core.AIStrategy
}

func (u *BaseUnit) ID() string          { return u.id }
func (u *BaseUnit) Name() string        { return u.name }
func (u *BaseUnit) Team() string        { return u.team }
func (u *BaseUnit) HP() int             { return u.hp }
func (u *BaseUnit) MaxHP() int          { return u.max }
func (u *BaseUnit) IsAlive() bool       { return u.hp > 0 }
func (u *BaseUnit) Skill() core.Skill   { return u.skill }
func (u *BaseUnit) AI() core.AIStrategy { return u.ai }

func (u *BaseUnit) TakeDamage(n int) (died bool) {
	if n < 0 {
		n = 0
	}
	if u.hp <= 0 {
		return false
	}
	u.hp -= n
	if u.hp <= 0 {
		u.hp = 0
		return true
	}
	return false
}

func (u *BaseUnit) Heal(n int) {
	if n < 0 {
		n = 0
	}
	if u.hp <= 0 {
		return
	}
	u.hp += n
	if u.hp > u.max {
		u.hp = u.max
	}
}
