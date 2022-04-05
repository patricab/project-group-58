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
	var num string
	flag.StringVar(&id, "id", "", "id of this peer")
	flag.StringVar(&port, "port", "", "port of this peer")
	flag.StringVar(&num, "num", "", "number of elevator nodes")
	flag.Parse()

	_id, _ := strconv.Atoi(id)
	_port, _ := strconv.Atoi(port)
	_num, _ := strconv.Atoi(num)

	var wg sync.WaitGroup
	wg.Add(1)
	go Distributor(_id, _port, _num)
	wg.Wait()
}
