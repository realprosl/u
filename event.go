package u

import "fmt"

type Event struct {
	state    bool
	eval func()bool
	action   func()
}

func (e *Event) On(action func()){
	e.action = action
	go func() {
		for {
			e.state = e.eval()
			fmt.Println("state", e.state)
			if e.state {
				e.action()
				return
			}
		}
	}()
}
func NewEvent(cond func() bool) *Event {

	var newEvent Event = Event{
		state: false,
		eval : cond ,
	}
	return &newEvent
}
