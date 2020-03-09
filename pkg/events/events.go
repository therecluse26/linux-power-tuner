package events

import (
	"time"
)

type ConditionType int

const (
	MonitoredFile ConditionType = iota
	ChargingStatus
	BatteryLevel
	CpuLoad
)

type Condition struct {
	Type ConditionType
	ExpectedValue interface{}
}

// Listener contains a slice of events to listen for
type Listener struct {
	Events []Event
}

// Event represents a particular
type Event struct {
	Condition Condition
	Timestamp time.Time
	Reactions []Reaction
	PollingTimeMs int32
	Count int32
	Active bool
}

type Reaction struct {
	Name string
	Function func(...interface{})
	Options []interface{}
	Active bool
}
/*
 * Fire an event and all reactions registered to event
 */
func (e *Event) Fire(){
	for _, r := range e.Reactions {
		r.Function(r.Options)
	}
}

func (e *Event) CheckValue(){
	foundValue := ""

	if foundValue == e.Condition.ExpectedValue {
		e.Fire()
	}
}

/*
 * Registers new event to Listener
 */
func (l *Listener) RegisterEvent (event Event) {
	l.Events = append(l.Events, event)
}

/*
 * Creates new event
 */
func EventFactory(condition Condition, reactions []Reaction, pollingtime int32) Event {
	return Event{ Condition: condition,
					Timestamp: time.Now(),
					Reactions: reactions,
					PollingTimeMs: pollingtime,
					Count: 0,
					Active: true }
}
