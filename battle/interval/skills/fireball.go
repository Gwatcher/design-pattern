package skills

import "designer/battle/interval/core"

type Fireball struct{}

func (Fireball) Name() string { return "Fireball" }

func (Fireball) Use(ctx *core.BattleContext, caster core.Unit, all []core.Unit) {
	enemy := pickLowestHPEnemy(caster, all)
	if enemy == nil {
		return
	}
	dmg := 18
	ctx.Bus.Publish(core.EventCast, caster.Name()+" casts Fireball on "+enemy.Name())
	died := enemy.TakeDamage(dmg)
	ctx.Bus.Publish(core.EventDamage, enemy.Name()+" takes "+itoa(dmg)+" dmg (HP "+itoa(enemy.HP())+"/"+itoa(enemy.MaxHP())+")")
	if died {
		ctx.Bus.Publish(core.EventDeath, enemy.Name()+" died")
	}
}
