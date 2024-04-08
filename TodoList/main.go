package main

import (
	"fmt"
)

var (
	weekDays  = [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	taskCount int
	taskState bool
	response  int
	task string
)

func main() {
	// Code to generate plan for the week, given already existing days
	fmt.Println("How many tasks are you willing to pursue")
	fmt.Scanln(&taskCount)
	var tasks = []string{}
	// Declaration of days and number of tasks per day
	for i := 0; i < taskCount; i++ {

		fmt.Scanln(&task)
		tasks = append(tasks, task)
		
	}
	fmt.Println(tasks[:])
	// Loop filling the necessary info

	// Check condition of task - Completed, Pending, Unable to perfom, Postponed

	// Postponing of tasks to different days

	// Record tasks completed, and pending giving end of week feedback
}