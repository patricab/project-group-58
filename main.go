package main

import (
	"fmt"

	"./backup"
)

type HallOrder struct {
	Up   []bool
	Down []bool
}

type Order struct {
	CabOrder      []bool
	HallOrder     HallOrder
	PriorityQueue []int
}

func main() {

	var orders Order // Initializing the custom struct

	// // Creating the custom orders
	orders.PriorityQueue = []int{4, 3, 1}
	orders.CabOrder = []bool{false, true, false, true}
	orders.HallOrder.Up = []bool{false, false, false}
	orders.HallOrder.Down = []bool{false, true, true}
	queue := orders.PriorityQueue
	cab := orders.CabOrder
	hall := [][]bool{orders.HallOrder.Up, orders.HallOrder.Down}

	backup.SaveOrder(queue, cab, hall)                       // Save order to file
	queueBackup, cabBackup, hallBackup := backup.LoadOrder() // Read from file

	// Print
	fmt.Printf("Type: %T, Queue: %v\n", queueBackup, queueBackup)
	fmt.Printf("Type: %T, Cab: %v\n", cabBackup, cabBackup)
	fmt.Printf("Type: %T, Hall: %v\n", hallBackup, hallBackup)

}
