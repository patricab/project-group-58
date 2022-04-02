package main

import (
	"fmt"
	"network/network"
	"strconv"
	"sync"
	"time"
)

// var timeout = 2

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go node("1", 2)
	go node("2", 3)

	wg.Wait()
}

func node(id string, dest int) {
	tx := make(chan network.Msg)
	rx := make(chan network.Msg)

	go network.Handler(id, tx, rx)

	_id, _ := strconv.Atoi(id)
	go recieved_cmdACK(_id, tx, rx)

	alive := make(chan bool)
	exit := make(chan bool)
	go watchdog(_id, dest, alive, exit, tx, rx)

	for {
		select {
		case a := <-alive:
			if a {
				fmt.Println("Node alive; Recieved ACK")
			} else {
				fmt.Println("Node dead")
			}
		}
	}
}

/**
 * Automatic reply to cmdACK commands
 * (Run this function as a goroutine!)
 */
func recieved_cmdACK(id int, tx chan network.Msg, rx chan network.Msg) {
	for {
		select {
		case m := <-rx:
			/* Check that we have recieved heartbeat command */
			if (m.Command == network.CmdACK) && (m.Data == 0) {
				/* Send ACK */
				time.Sleep(1 * time.Second)
				fmt.Println("Slave: Recieved message, sending ACK")
				msg := network.Msg{id, m.Id, network.CmdACK, 1}
				tx <- msg
			}
		}
	}
}

/**
 * Watchdog: Check if node is still alive
 * id - ID of current node
 * dest - (int) ID of target node
 * alive - (channel) Status of target node
 * exit - (channel) Turn off watchdog
 */
func watchdog(id int, dest int, alive chan bool, exit chan bool, tx chan network.Msg, rx chan network.Msg) {
	// var rx_msg = network.Msg{dest, id, network.CmdACK, 0}
	// timer := time.NewTimer(3 * time.Second)
	// var timer time.Timer

	for {
		select {
		/* Stop watchdog */
		case <-exit:
			return

		default:
			/* Send heartbeat command */
			msg := network.Msg{id, dest, network.CmdACK, 0}
			fmt.Println("Master: Sending message")
			tx <- msg

			/* Start timer */
			// timer = *time.NewTimer(time.Duration(timeout) * time.Second)
			// timer = *time.NewTimer(3 * time.Second)
			// timer.Reset(3 * time.Second)
		}

		select {
		case m := <-rx:
			/* Recieved response */
			if (m.Command == network.CmdACK) && (m.Data == 1) {
				// timer.Stop()  // Stop timer
				alive <- true // Signal user that target is alive
				fmt.Printf("Master: Received: %#v\n", m)
			}

		/* Timer ended, no response */
		case <-time.After(3 * time.Second):
			// case <-timer.C:
			alive <- false

		}
	}
}
