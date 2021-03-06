package backup

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

const BACKUP_FILENAME = "orders"
const BACKUP_PATH = "orders/"

func SaveOrder(queue []int, cab []bool, hall [][]bool) {

	var orders Order

	orders.CabOrder = cab
	orders.HallOrder.Up = hall[0]
	orders.HallOrder.Down = hall[1]
	orders.PriorityQueue = queue

	filedata, errMarshal := json.MarshalIndent(&orders, "", " ")
	check(errMarshal)

	filepath := BACKUP_PATH + BACKUP_FILENAME + ".json"
	errWrite := ioutil.WriteFile(filepath, filedata, 0644)
	check(errWrite)
}

func LoadOrder() ([]int, []bool, [][]bool) {

	var orders Order

	filepath := BACKUP_PATH + BACKUP_FILENAME + ".json"
	filedata, errRead := ioutil.ReadFile(filepath)
	check(errRead)

	errUnmarshal := json.Unmarshal(filedata, &orders)
	check(errUnmarshal)

	cab := orders.CabOrder
	hallup := orders.HallOrder.Up
	halldown := orders.HallOrder.Down
	hall := [][]bool{hallup, halldown}
	queue := orders.PriorityQueue

	return queue, cab, hall
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
