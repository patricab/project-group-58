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

// CONSTANTS TO ACCOUNT FOR
const fname = "orders.txt"
const BACKUP_FILENAME = "orders"
const N_ELEVATORS = 3
const M_FLOORS = 4
const N_FILES = 3

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

	for i := 0; i < N_FILES; i++ {
		filename := BACKUP_FILENAME + strconv.Itoa(i) + ".json"
		errWrite := ioutil.WriteFile(filename, filedata, 0644)
		check(errWrite)
	}

}

func LoadOrderJSON() ([]bool, [][]bool) {

	m := map[string][][]bool{
		"cab":      {{}},
		"hallUp":   {{}},
		"hallDown": {{}},
	}

	m["cab"] = make([][]bool, N_FILES)
	m["hallUp"] = make([][]bool, N_FILES)
	m["hallDown"] = make([][]bool, N_FILES)

	var orders Order

	for i := 0; i < N_FILES; i++ {
		filename := BACKUP_FILENAME + strconv.Itoa(i) + ".json"
		fmt.Println(filename)
		filedata, err := ioutil.ReadFile(filename)
		check(err)

		//var orders Order
		errUnmarshal := json.Unmarshal(filedata, &orders)
		check(errUnmarshal)

		c := orders.CabOrder
		fmt.Println(c)

		// Append data to list for comparison
		// LAST CALL TO STRUCT OVERWRITES DATA IN ALL SLICES
		m["cab"][i] = c
		m["hallUp"][i] = orders.HallOrder.Up
		m["hallDown"][i] = orders.HallOrder.Down
		fmt.Println(m["cab"])
	}

	// Compare data
	areEqual := true
	for i := 0; i < N_FILES; i++ {
		areEqual = areElementsEqual(m["cab"][0], m["cab"][i])
		areEqual = areElementsEqual(m["hallUp"][0], m["hallUp"][i])
		areEqual = areElementsEqual(m["hallDown"][0], m["hallDown"][i])
	}

	if areEqual != true {
		// Find most common element
		// ADD THEM INTO FIRST SLICE
		fmt.Println("ORDERS ARE NOT EQUAL!")
	}

	cab := m["cab"][0]
	hallup := m["hallUp"][0]
	halldown := m["hallDown"][0]
	hall := [][]bool{hallup, halldown}
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

func areElementsEqual(x, y []bool) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true

}
