package main

import (
	"fmt"
	"log"
)

// Global variables
const NUME_PROC int = 5

// Structure
type Process struct {
	id int
	bt int
	at int
	ct int
	wt int
}

func show_processes(proc []Process) {
	var i int
	fmt.Print("ID\tBT\tAT\n")
	for i = 0; i < NUME_PROC; i++ {
		fmt.Printf("%d\t%d\t%d\n",
			proc[i].id, proc[i].bt, proc[i].at)
	}
}

func insert_processes(proc []Process) {
	// Variable i declaration
	var i int

	// Inserting processes
	for i = 0; i < NUME_PROC; i++ {
		fmt.Println("Insert ID, BT, AT: ")
		_, err := fmt.Scan(&proc[i].id, &proc[i].bt, &proc[i].at)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func order_processes_at(proc []Process) {
	// Variable declaration
	var i, j int
	var aux Process

	for i = 0; i < NUME_PROC-1; i++ {
		for j = i + 1; j < NUME_PROC; j++ {
			if proc[j].at < proc[i].at {
				aux = proc[i]
				proc[i] = proc[j]
				proc[j] = aux
			}
		}
	}
}

func show_gantt(proc []Process) {
	// Variable declaration
	var i, ms_init, ms_end int

	for i = 0; i < NUME_PROC; i++ {
		if i == 0 {
			ms_init = proc[i].at
		} else {
			if proc[i].at <= ms_end {
				ms_init = ms_end
			} else {
				ms_init = proc[i].at
			}
		}

		ms_end = ms_init + proc[i].bt

		// Update CT and WT
		proc[i].ct = ms_end - proc[i].at
		proc[i].wt = proc[i].ct - proc[i].bt
		fmt.Printf("[%d-P%d-%d]", ms_init, proc[i].id, ms_end)
	}
	fmt.Println()
}

func order_processes_id(proc []Process) {
	// Variable declaration
	var i, j int
	var aux Process

	// Ordering processes
	for i = 0; i < NUME_PROC-1; i++ {
		for j = i + 1; j < NUME_PROC; j++ {
			if proc[j].id < proc[i].id {
				aux = proc[i]
				proc[i] = proc[j]
				proc[j] = aux
			}
		}
	}
}

func show_processes_wt_ct(proc []Process) {
	// Variable declaration
	var i int
	var sum_wt, sum_ct int = 0, 0
	var avg_wt, avg_ct float64

	// Showing processes
	fmt.Println("ID\tWT\tCT")
	for i = 0; i < NUME_PROC; i++ {
		fmt.Printf("%d\t%d\t%d\n", proc[i].id, proc[i].wt, proc[i].ct)
		sum_wt += proc[i].wt
		sum_ct += proc[i].ct
	}
	avg_wt = float64(sum_wt) / float64(NUME_PROC)
	avg_ct = float64(sum_ct) / float64(NUME_PROC)

	fmt.Println()
	fmt.Printf("Average WT: %f\n", avg_wt)
	fmt.Printf("Average CT: %f\n", avg_ct)
}

func main() {
	// Procedure declaration
	proc := make([]Process, NUME_PROC)

	// Inserting processes
	insert_processes(proc)
	fmt.Println()
	fmt.Println("The inserted processes are: ")
	show_processes(proc)
	fmt.Println()

	// Ordering processes by AT
	order_processes_at(proc)
	fmt.Println("The ordered processes by AT are: ")
	show_processes(proc)
	fmt.Println()

	// Show Gantt diagram
	fmt.Println("The Gantt diagram is: ")
	show_gantt(proc)
	fmt.Println()

	// Order processes by ID
	order_processes_id(proc)
	fmt.Println("The ordered processes by ID, with its WT and CT are: ")
	show_processes_wt_ct(proc)
}
