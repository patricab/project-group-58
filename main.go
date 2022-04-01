package main

import (
	"fmt"

	"./backup"
)

func main() {
	// Order
	//orders := LoadOrderJSON()
	//fmt.Println(orders)
	fmt.Println("hey")
	data := backup.LoadCab("orders.txt")
	fmt.Println(data)
}
