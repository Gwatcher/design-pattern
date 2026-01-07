package core

import "time"

type EventType string

const (
	EventTurnStart EventType = "turn_start"
	EventCast      EventType = "cast"
	EventDamage    EventType = "damage"
	EventHeal      EventType = "heal"
	EventDeath     EventType = "death"
	EventBattleEnd EventType = "battle_end"
)

type Event struct {
	Type EventType
	Msg  string
	Time time.Time
}
