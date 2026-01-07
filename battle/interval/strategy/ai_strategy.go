package strategy

import "designer/battle/interval/core"

type AggressiveAI struct{}

func (AggressiveAI) Act(ctx *core.BattleContext, self core.Unit, all []core.Unit) {
	// 直接用技能（技能内部可做目标选择）
	self.Skill().Use(ctx, self, all)
}

type SupportAI struct{}

func (SupportAI) Act(ctx *core.BattleContext, self core.Unit, all []core.Unit) {
	// 同样交给技能决定：Heal 会找队友，其他技能找敌人
	self.Skill().Use(ctx, self, all)
}
