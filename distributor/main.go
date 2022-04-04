package main

import (
	. "distributor/distributor"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go Distributor(1)
	wg.Wait()
}
