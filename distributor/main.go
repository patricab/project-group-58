package main

import (
	. "distributor/distributor"
	"flag"
	"strconv"
	"sync"
)

func main() {
	var id string
	var port string
	flag.StringVar(&id, "id", "", "id of this peer")
	flag.StringVar(&port, "port", "", "port of this peer")
	flag.Parse()

	_id, _ := strconv.Atoi(id)
	_port, _ := strconv.Atoi(port)

	var wg sync.WaitGroup
	wg.Add(1)
	go Distributor(_id, _port)
	wg.Wait()
}
