package main

import (
	"backup/backup"
	"fmt"
)

type HallOrder struct {
	HallUP   []bool
	HallDOWN []bool
}

type Order struct {
	CabOrder  []bool
	HallOrder HallOrder
}

func main() {
	// Order
	//orders := LoadOrderJSON()
	//fmt.Println(orders)
	boolOrder := []bool{true, true, true}
	backup.SaveCab(boolOrder)
	data := backup.LoadCab("orders.txt")
	fmt.Println(data)

	// JSON
	//json_data := backup.LoadOrderJSON("test.json")
	backup.LoadOrderJSON("test.json")
}
