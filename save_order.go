package main

/* Utility for saving and loading backups of orders to .txt files */

/*
EXAMPLE FILE STRUCTURE

CAB ORDERS
Floor	1		2		3		4
		bool	bool	bool	bool

HALL ORDERS
		1		2		3		4
up		bool	bool	bool
down			bool	bool	bool

*/

import(
	"os"
	"fmt"
)

func SaveOrder(order orderType, filename string) {
	// Save orders to file

	f, err := os.Open(filename)
	check(err)

	defer f.Close() // Idiomatic way

	// 1. Initialization or check if files are okay.

	// 2. Saving Orders: Modifying the file(s)

}

func LoadOrders() order orderType {
	// Reads from the saved orders and returns content

	f, err := os.Open(filename)
	check(err)

	defer f.Close() // Idiomatic way

	// 1. Read from file

	// 2. Return chosen structure
}

func ReadOrders(filename string) data string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return data // Might not be string
}

func check(e error) {
	if e != nil {
		Log.Fatal(e)
	}
}