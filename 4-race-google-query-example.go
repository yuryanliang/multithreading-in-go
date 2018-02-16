package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	query := "Our Query"
	respond := make(chan string)

	go googleIt(respond, query)

	queryResp := <-respond

	fmt.Printf("Sent query:\t\t %s\n", query)
	fmt.Printf("Got Response:\t\t %s\n", queryResp)
}

func googleIt(respond chan<- string, query string) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	respond <- "A Google Response"
}
