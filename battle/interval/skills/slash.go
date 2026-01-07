package skills

import "designer/battle/interval/core"

type Slash struct{}

func (Slash) Name() string { return "Slash" }

func (Slash) Use(ctx *core.BattleContext, caster core.Unit, all []core.Unit) {
	enemy := pickLowestHPEnemy(caster, all)
	if enemy == nil {
		return
	}
	dmg := 12
	ctx.Bus.Publish(core.EventCast, caster.Name()+" uses Slash on "+enemy.Name())
	died := enemy.TakeDamage(dmg)
	ctx.Bus.Publish(core.EventDamage, enemy.Name()+" takes "+itoa(dmg)+" dmg (HP "+itoa(enemy.HP())+"/"+itoa(enemy.MaxHP())+")")
	if died {
		ctx.Bus.Publish(core.EventDeath, enemy.Name()+" died")
	}
}
