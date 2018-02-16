package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	respond := make(chan string, 3)

	go checkDNS(respond, "pragmacoders.com", "ns1.nameserver.com")
	go checkDNS(respond, "pragmacoders.com", "ns2.nameserver.com")
	go checkDNS(respond, "pragmacoders.com", "ns3.nameserver.com")
	go checkDNS(respond, "pragmacoders.com", "ns4.nameserver.com")
	go checkDNS(respond, "pragmacoders.com", "ns5.nameserver.com")

	for i := 1; i <= 3; i++ {
		queryResp := <-respond
		fmt.Printf("Got Response:\t %s\n", queryResp)
	}
}

func checkDNS(respond chan<- string, query string, ns string) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	respond <- fmt.Sprintf("%s responded to query: %s", ns, query)
}
