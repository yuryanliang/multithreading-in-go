package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	respond := make(chan string, 5)
	var wg sync.WaitGroup

	wg.Add(5)
	go checkDNS(respond, &wg, "pragmacoders.com", "ns1.nameserver.com")
	go checkDNS(respond, &wg, "pragmacoders.com", "ns2.nameserver.com")
	go checkDNS(respond, &wg, "pragmacoders.com", "ns3.nameserver.com")
	go checkDNS(respond, &wg, "pragmacoders.com", "ns4.nameserver.com")
	go checkDNS(respond, &wg, "pragmacoders.com", "ns5.nameserver.com")

	wg.Wait()
	close(respond)

	for queryResp := range respond {
		fmt.Printf("Got Response:\t %s\n", queryResp)
	}
}

func checkDNS(respond chan<- string, wg *sync.WaitGroup, query string, ns string) {
	defer wg.Done()

	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	respond <- fmt.Sprintf("%s responded to query: %s", ns, query)
}
