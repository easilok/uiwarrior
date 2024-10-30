package main

import (
	"fmt"

	"github.com/easilok/uiwarrior/lib/taskwarrior"
)

func main() {

	var taskWarrior taskwarrior.Taskwarrior
	if err := taskWarrior.Load(); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Completed Tasks")
	fmt.Println("----------------")
	for _, task := range taskWarrior.Completed() {
		fmt.Println(task.Text())
	}

	fmt.Println()
	fmt.Println("Pending Tasks")
	fmt.Println("----------------")
	for _, task := range taskWarrior.Pending() {
		fmt.Println(task.Text())
	}

	fmt.Println()
	fmt.Println("Projects")
	fmt.Println("----------------")
	fmt.Printf("%v\n", taskWarrior.Projects())

	fmt.Println()
	fmt.Println("Tags")
	fmt.Println("----------------")
	fmt.Printf("%v\n", taskWarrior.Tags())
}
