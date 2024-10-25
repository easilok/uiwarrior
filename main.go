package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"sort"

	"github.com/easilok/uiwarrior/models"
)

func main() {

	// exec.Command("task", "context", "none")
	cmd := exec.Command("task", "export")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var tasks []models.Task
	err = json.Unmarshal(stdout, &tasks)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	pendingTasks := []models.Task{}

	for _, task := range tasks {
		if task.Status != "completed" {
			pendingTasks = append(pendingTasks, task)
		}
	}

	sort.Slice(pendingTasks, func(i, j int) bool {
		return pendingTasks[i].Urgency > pendingTasks[j].Urgency
	})

	for _, task := range pendingTasks {
		fmt.Println(task.Text())
	}
}
