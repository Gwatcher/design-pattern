package factory

import (
	"designer/battle/interval/core"
	"designer/battle/interval/units"
	"fmt"
)

type UnitType string

const (
	WarriorType UnitType = "warrior"
	MageType    UnitType = "mage"
	PriestType  UnitType = "priest"
)

type UnitFactory struct {
	ctors map[UnitType]func(id, team string, ai core.AIStrategy) core.Unit
}

func NewUnitFactory() *UnitFactory {
	f := &UnitFactory{ctors: map[UnitType]func(string, string, core.AIStrategy) core.Unit{}}
	f.Register(WarriorType, func(id, team string, ai core.AIStrategy) core.Unit { return units.NewWarrior(id, team, ai) })
	f.Register(MageType, func(id, team string, ai core.AIStrategy) core.Unit { return units.NewMage(id, team, ai) })
	f.Register(PriestType, func(id, team string, ai core.AIStrategy) core.Unit { return units.NewPriest(id, team, ai) })
	return f
}

func (f *UnitFactory) Register(t UnitType, ctor func(id, team string, ai core.AIStrategy) core.Unit) {
	f.ctors[t] = ctor
}

func (f *UnitFactory) Create(t UnitType, id, team string, ai core.AIStrategy) (core.Unit, error) {
	ctor, ok := f.ctors[t]
	if !ok {
		return nil, fmt.Errorf("unknown unit type: %s", t)
	}
	return ctor(id, team, ai), nil
}
