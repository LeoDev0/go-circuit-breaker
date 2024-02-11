package circuitbreaker

import (
	"fmt"
	"github.com/sony/gobreaker"
	"time"
)

func Instantiate() *gobreaker.CircuitBreaker {
	return gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "my-circuit-breaker",
		MaxRequests: 3,
		Timeout:     3 * time.Second,
		Interval:    1 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 3
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			fmt.Printf("CircuitBreaker '%s' changed from '%s' to '%s'\n", name, from, to)
		},
	})
}
