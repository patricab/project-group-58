package main

import (
	"fmt"

	"./backup"
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
	// boolOrder := []bool{true, true, true}
	// backup.SaveCab(boolOrder)
	// data := backup.LoadCab("orders.txt")
	// fmt.Println(data)
	// fmt.Printf("%T\n", data)

	// JSON
	//json_data := backup.LoadOrderJSON("test.json")
	cab, hall := backup.LoadOrderJSON("test.json")
	fmt.Printf("Type: %T, Cab: %v\n", cab, cab)
	fmt.Printf("Type: %T, Cab: %v\n", hall, hall)
}
