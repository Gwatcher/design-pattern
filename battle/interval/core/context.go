package core

type BattleContext struct {
	Bus *EventBus
	Rng *Rand // 简单随机工具（见 battle.go 里）
}
