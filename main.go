package main

import (
	"fmt"
	"github.com/leodev0/cb/circuitbreaker"
	"github.com/leodev0/cb/externalservice"
	"github.com/leodev0/cb/restclient"
	"time"
)

func main() {
	go externalservice.InitServer()

	cb := circuitbreaker.Instantiate()

	fmt.Printf("Call with circuit breaker")

	// execute 100 requests in different circumstances to test the circuit breaker
	for i := 0; i < 100; i++ {
		_, err := cb.Execute(func() (interface{}, error) {
			err := restclient.GetEndpoint()
			return nil, err
		})
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
