package main

import (
	"fmt"
	"time"
)

func main() {
	request := make(chan int, 5) //canal que simula las solicitudes

	for i := 1; i <= 5; i++ {
		request <- i
	}

	close(request)

	limiter := time.Tick(1000 * time.Millisecond) //cada 1 segundo informa, constantemente

	for req := range request {
		<-limiter // bloquea el for, y cada vez que haga un tick lo desbloquea, es como hacer solo una lectura
		fmt.Println("Request: ", req)
	}

	fmt.Println("*")

	//************************************

	rafaga := make(chan time.Time, 3)

	go func() {
		for t := range time.Tick(1000 * time.Millisecond) {
			for i := 0; i < 3; i++ {
				rafaga <- t
				fmt.Println("*")
			}
		}
	}()

	rafagaRequest := make(chan int, 15)
	for i := 1; i <= 15; i++ {
		rafagaRequest <- i
	}

	close(rafagaRequest)

	for req := range rafagaRequest {
		<-rafaga
		fmt.Println("Request: ", req, time.Now())
	}

}
