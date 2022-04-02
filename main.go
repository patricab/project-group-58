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
	CabOrder  []bool
	HallOrder HallOrder
}

func main() {

	// var orders Order // Initializing the custom struct

	// // Creating the custom orders
	// orders.CabOrder = []bool{false, false, false, true}
	// orders.HallOrder.Up = []bool{false, false, false}
	// orders.HallOrder.Down = []bool{false, true, false}
	// cab := orders.CabOrder
	// hall := [][]bool{orders.HallOrder.Up, orders.HallOrder.Down}

	// backup.SaveOrderJSON(cab, hall)       // Save order to file
	cab2, hall2 := backup.LoadOrderJSON() // Read from file

	// Print
	fmt.Printf("Type: %T, Cab: %v\n", cab2, cab2)
	fmt.Printf("Type: %T, Hall: %v\n", hall2, hall2)

}
