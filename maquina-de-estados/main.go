// https://youtu.be/wjf3Zmw4oUU
package main

import "fmt"

const (
	StateCreated    = "created"
	StateProcessing = "processing"
	StateCompleted  = "completed"
	StateCancelled  = "cancelled"

	EventProcess   = "process"
	EventeComplete = "complete"
	EventCancel    = "cancel"
)

type Callback func(from, to, event string)

type StateMachine struct {
	state       string
	transitions map[string]map[string]string
	callback    map[string]Callback
}

func NewStateMachine(initialState string, transitions map[string]map[string]string, callback map[string]Callback) *StateMachine {
	return &StateMachine{
		state:       initialState,
		transitions: transitions,
		callback:    callback,
	}
}

func (sm *StateMachine) Trigger(event string) error {
	if nextState, ok := sm.transitions[sm.state][event]; ok {
		if callback, exists := sm.callback[event]; exists {
			callback(sm.state, nextState, event)
		}

		sm.state = nextState
		return nil
	}

	return fmt.Errorf("invalid transition from %s to %s", sm.state, event)
}

func (sm *StateMachine) State() string {
	return sm.state
}

func main() {
	transitions := map[string]map[string]string{
		StateCreated: {
			EventProcess: StateProcessing,
			EventCancel:  StateCancelled,
		},
		StateProcessing: {
			EventeComplete: StateCompleted,
			EventCancel:    StateCancelled,
		},
	}

	callbacks := map[string]Callback{
		EventProcess: func(from, to, event string) {
			fmt.Println("Processing", event, "from", from, "to", to)
		},
		EventeComplete: func(from, to, event string) {
			fmt.Println("Completing", event, "from", from, "to", to)
		},
		EventCancel: func(from, to, event string) {
			fmt.Println("Cancelling", event, "from", from, "to", to)
		},
	}

	sm := NewStateMachine(StateCreated, transitions, callbacks)

	events := []string{EventProcess, EventeComplete, EventCancel}
	for _, event := range events {
		if err := sm.Trigger(event); err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Current state:", sm.State())
}

/*
const (
	StateOpen   = "open"
	StateClosed = "closed"
	StateLocked = "locked"
)

const (
	EventOpen     = "open"
	EventClosed   = "closed"
	EventLocked   = "locked"
	EventUnlocked = "unlocked"
)

type StateMachine struct {
	state       string
	transitions map[string]map[string]string
}

func NewStateMachine(initialState string, transitions map[string]map[string]string) *StateMachine {
	return &StateMachine{
		state:       initialState,
		transitions: transitions,
	}
}

func (sm *StateMachine) State() string {
	return sm.state
}

func (sm *StateMachine) Trigger(event string) error {
	nextStep, ok := sm.transitions[sm.state][event]
	if ok {
		fmt.Println("Triggering", event, "from", sm.state, "to", nextStep)
		sm.state = nextStep
		return nil
	}

	return fmt.Errorf("invalid transition from %s to %s", sm.state, event)
}

func main() {
	transition := map[string]map[string]string{
		StateOpen: {
			EventClosed: StateClosed,
			EventLocked: StateLocked,
		},
		StateClosed: {
			EventOpen: StateOpen,
		},
		StateLocked: {
			EventUnlocked: StateOpen,
		},
	}

	sm := NewStateMachine(StateOpen, transition)

	event := []string{
		EventClosed,
		EventLocked,
		EventUnlocked,
		EventOpen,
	}

	for _, e := range event {
		err := sm.Trigger(e)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Final state is", sm.State())
}
*/
