package main

/* Utility for saving and loading backups of orders */

import(
	"os"
	"fmt"
)

func SaveOrder(order orderType, filename string) {
	// Saving given orders to a file
	// Hall and cab?

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// Save orders

}

func LoadOrders() order orderType {
	// Reads from the saved orders and returns content
	// ReadOrders
	// ///

	// Return file content or do decisions here?
}

func ReadOrders(filename string) data string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return data // Might not be string
}