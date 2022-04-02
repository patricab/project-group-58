package backup

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

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type HallOrder struct {
	HallUP   []bool
	HallDOWN []bool
}

type Order struct {
	CabOrder  []bool
	HallOrder HallOrder
}

const fname = "orders.txt"

func SaveCab(order []bool) {
	// Saves cab order as a string in a text file.

	order_string := ""

	for i := 0; i < len(order); i++ {
		order_string = order_string + " " + strconv.FormatBool(order[i])
	}

	order_string = order_string[1:]

	err := ioutil.WriteFile(fname, []byte(order_string), 0644)
	check(err)
}

func SaveOrderJSON(order Order) {
	fmt.Println("Under production!")
}

func LoadOrderJSON(filename string) ([]bool, [][]bool) {

	data, err := ioutil.ReadFile(filename)
	check(err)

	var orders Order
	json.Unmarshal(data, &orders)

	cab := orders.CabOrder
	hallup := orders.HallOrder.HallUP
	halldown := orders.HallOrder.HallDOWN
	hall := [][]bool{hallup, halldown}
	fmt.Println(cab)
	fmt.Println(hall)
	return cab, hall
}

func LoadCab(filename string) []bool {
	// Reads saved calls and parses it into a bool array.

	data, err := ioutil.ReadFile(filename)
	check(err)

	stringData := strings.Split(string(data), " ")
	savedOrder := make([]bool, len(stringData))

	for i := 0; i < len(stringData); i++ {
		boolVal, err := strconv.ParseBool(stringData[i])
		check(err)

		savedOrder[i] = boolVal
	}
	return savedOrder
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
