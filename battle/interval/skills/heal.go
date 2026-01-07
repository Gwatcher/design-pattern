package skills

import "designer/battle/interval/core"

type Heal struct{}

func (Heal) Name() string { return "Heal" }

func (Heal) Use(ctx *core.BattleContext, caster core.Unit, all []core.Unit) {
	ally := pickLowestHPAlly(caster, all)
	if ally == nil {
		return
	}
	amt := 15
	ctx.Bus.Publish(core.EventCast, caster.Name()+" casts Heal on "+ally.Name())
	ally.Heal(amt)
	ctx.Bus.Publish(core.EventHeal, ally.Name()+" heals "+itoa(amt)+" (HP "+itoa(ally.HP())+"/"+itoa(ally.MaxHP())+")")
}
