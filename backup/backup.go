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
	Up   []bool
	Down []bool
}

type Order struct {
	CabOrder  []bool
	HallOrder HallOrder
}

const fname = "orders.txt"
const jsonfile = "orders.json"

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

func SaveOrderJSON(cab []bool, hall [][]bool) {

	var orders Order

	orders.CabOrder = cab
	orders.HallOrder.Up = hall[0]
	orders.HallOrder.Down = hall[1]

	filedata, errMarshal := json.MarshalIndent(&orders, "", " ")
	check(errMarshal)
	fmt.Printf("%T is the type for the data\n", filedata)

	errWrite := ioutil.WriteFile(jsonfile, filedata, 0644)
	check(errWrite)

}

func LoadOrderJSON() ([]bool, [][]bool) {
	// Returns Cab and Hall Orders

	data, errRead := ioutil.ReadFile(jsonfile)
	check(errRead)

	var orders Order
	errUnmarshal := json.Unmarshal(data, &orders)
	check(errUnmarshal)

	cab := orders.CabOrder
	hallup := orders.HallOrder.Up
	halldown := orders.HallOrder.Down
	hall := [][]bool{hallup, halldown}
	// fmt.Println(cab)
	// fmt.Println(hall)
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
