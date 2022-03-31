package main

import (
	"backup/backup"
	"fmt"
)

func main() {
	// Order
	test_order := []bool{true, false, false, false, false, true}
	fmt.Printf("Type: %T, Test order:   %v\n", test_order, test_order)

	backup.SaveCab(test_order) // Save

	backup_order := backup.LoadOrder("orders.txt") // Load

	fmt.Printf("Type: %T, Loaded order: %v \n", backup_order, backup_order)

}
