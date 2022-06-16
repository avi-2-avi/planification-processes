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
	fmt.Print("ID\tAT\tBT\n")
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

func order_processes_sjf(proc []Process) []Process {
	var NUM_PROC_LEFT int = NUME_PROC
	var time int = 0
	var is_busy bool = false

	var current_proc Process
	var time_left int = 0

	var waiting_list []Process
	var new_proc []Process

	fmt.Println("We begin here!")
	for {
		print("Time: ", time, "ms\n")

		if is_busy {
			time_left--

			// End busy time
			if time_left == 0 {
				is_busy = false

				if len(waiting_list) > 0 {
					// Remove current process from waiting list
					for i := 0; i < len(waiting_list); i++ {
						if waiting_list[i].id == current_proc.id {
							waiting_list = append(waiting_list[:i], waiting_list[i+1:]...)
							break
						}
					}

					fmt.Println("Process P", current_proc.id, "ended.")

					// Substract processes left
					NUM_PROC_LEFT--
				}
			}
		} else if is_busy == false && len(waiting_list) > 0 {
			// Start process with time AT for next milisecond
			fmt.Println("Waiting list: ", waiting_list)
			fmt.Println(len(waiting_list))

			var shortest_bt int = 10000

			for i := 0; i < len(waiting_list); i++ {
				for j := i; j < len(waiting_list); j++ {
					if waiting_list[j].bt < shortest_bt {
						shortest_bt = waiting_list[j].bt
						current_proc = waiting_list[j]
					}
				}
			}
			// New time left
			time_left = shortest_bt

			// Now it's busy
			is_busy = true

			fmt.Println("Process P", current_proc.id, "started.")

			// Add process in the last position of new_proc
			new_proc = append(new_proc, current_proc)

			time_left--
		}

		fmt.Println("NUM_PROC_LEFT: ", NUM_PROC_LEFT)

		// Check AT of processes and add them to waiting list
		for i := 0; i < len(proc); i++ {
			if proc[i].at == time {
				waiting_list = append(waiting_list, proc[i])
				fmt.Println("Process", proc[i].id, "added to waiting list")
			}
		}

		// Terminate loop
		if NUM_PROC_LEFT == 0 {
			break
		}

		// Count time
		time++
		fmt.Println()
	}

	// Update processes
	fmt.Println("New processes: ", new_proc)
	return new_proc
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

	// Ordering processes by SJF
	proc = order_processes_sjf(proc)
	fmt.Println("The ordered processes by SJF are: ")
	// Update processes
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
