// https://youtu.be/wjf3Zmw4oUU
package main

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
