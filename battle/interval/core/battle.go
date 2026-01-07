package core

import (
	"math/rand"
	"time"
)

type Rand struct{ r *rand.Rand }

func NewRand() *Rand {
	return &Rand{r: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func (rr *Rand) Intn(n int) int { return rr.r.Intn(n) }

type Unit interface {
	ID() string
	Name() string
	Team() string
	IsAlive() bool

	HP() int
	MaxHP() int
	TakeDamage(n int) (died bool)
	Heal(n int)

	Skill() Skill
	AI() AIStrategy
}

type Skill interface {
	Name() string
	Use(ctx *BattleContext, caster Unit, all []Unit)
}

type AIStrategy interface {
	Act(ctx *BattleContext, self Unit, all []Unit)
}

type Battle struct {
	Ctx   *BattleContext
	Units []Unit
}

func NewBattle(bus *EventBus, units []Unit) *Battle {
	return &Battle{
		Ctx:   &BattleContext{Bus: bus, Rng: NewRand()},
		Units: units,
	}
}

func (b *Battle) Run(maxRounds int) {
	for round := 1; round <= maxRounds; round++ {
		if b.isOver() {
			return
		}
		b.Ctx.Bus.Publish(EventTurnStart, "===== Round "+itoa(round)+" =====")

		for _, u := range b.Units {
			if !u.IsAlive() || b.isOver() {
				continue
			}
			u.AI().Act(b.Ctx, u, b.Units)
		}
	}

	b.Ctx.Bus.Publish(EventBattleEnd, "Battle ended: reached max rounds")
}

func (b *Battle) isOver() bool {
	aliveTeams := map[string]bool{}
	for _, u := range b.Units {
		if u.IsAlive() {
			aliveTeams[u.Team()] = true
		}
	}
	return len(aliveTeams) <= 1
}

// tiny int->string without fmt（保持依赖最小）
func itoa(x int) string {
	if x == 0 {
		return "0"
	}
	neg := x < 0
	if neg {
		x = -x
	}
	var buf [20]byte
	i := len(buf)
	for x > 0 {
		i--
		buf[i] = byte('0' + x%10)
		x /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}
