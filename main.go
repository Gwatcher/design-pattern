package main

import (
	"fmt"

	"designer/battle/interval/core"
	"designer/battle/interval/factory"
	"designer/battle/interval/strategy"
)

func main() {
	bus := core.NewEventBus()

	// 订阅事件（观察者）
	bus.Subscribe(core.EventTurnStart, func(e core.Event) { fmt.Println(e.Msg) })
	bus.Subscribe(core.EventCast, func(e core.Event) { fmt.Println("  ", e.Msg) })
	bus.Subscribe(core.EventDamage, func(e core.Event) { fmt.Println("    ", e.Msg) })
	bus.Subscribe(core.EventHeal, func(e core.Event) { fmt.Println("    ", e.Msg) })
	bus.Subscribe(core.EventDeath, func(e core.Event) { fmt.Println("  ☠", e.Msg) })
	bus.Subscribe(core.EventBattleEnd, func(e core.Event) { fmt.Println(e.Msg) })

	uf := factory.NewUnitFactory()

	// 两队阵容（写死）
	var units []core.Unit

	// Team A
	u1, _ := uf.Create(factory.WarriorType, "A1", "A", strategy.AggressiveAI{})
	u2, _ := uf.Create(factory.MageType, "A2", "A", strategy.AggressiveAI{})
	u3, _ := uf.Create(factory.PriestType, "A3", "A", strategy.SupportAI{})
	units = append(units, u1, u2, u3)

	// Team B
	v1, _ := uf.Create(factory.WarriorType, "B1", "B", strategy.AggressiveAI{})
	v2, _ := uf.Create(factory.MageType, "B2", "B", strategy.AggressiveAI{})
	v3, _ := uf.Create(factory.PriestType, "B3", "B", strategy.SupportAI{})
	units = append(units, v1, v2, v3)

	battle := core.NewBattle(bus, units)
	battle.Run(50)
}
