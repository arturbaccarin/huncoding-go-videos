// https://youtu.be/7yPhLLP_Jj0
package main

import (
	"errors"
	"math/rand"
	"time"

	"github.com/sony/gobreaker"
)

func mockService() (string, error) {
	if rand.Intn(100) > 50 {
		return "success", nil
	}

	return "", errors.New("error trying to proccess request")
}

func logState(name string, from gobreaker.State, to gobreaker.State) {
	println(name, "state changed from", stateToString(from), "to", stateToString(to))
}

func stateToString(state gobreaker.State) string {
	switch state {
	case gobreaker.StateClosed:
		return "closed"
	case gobreaker.StateOpen:
		return "open"
	case gobreaker.StateHalfOpen:
		return "half-open"
	default:
		return "unknown"
	}
}

func main() {
	settings := gobreaker.Settings{
		Name:        "my-circuitbreaker",
		MaxRequests: 3,
		Interval:    5 * time.Second,
		Timeout:     5 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 2
		},
		OnStateChange: logState,
	}

	cb := gobreaker.NewCircuitBreaker(settings)

	for i := 0; i < 10; i++ {
		result, err := cb.Execute(func() (interface{}, error) {
			return mockService()
		})

		if err == nil {
			println(result.(string))
		} else {
			println(err.Error())
		}
	}
}
