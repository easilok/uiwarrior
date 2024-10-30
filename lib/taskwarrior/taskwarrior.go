package taskwarrior

import (
	"encoding/json"
	"os/exec"
	"sort"

	"github.com/easilok/uiwarrior/lib/taskwarrior/models"
)

type Taskwarrior struct {
	Tasks []models.Task
}

func (tw *Taskwarrior) Load() error {
	cmd := exec.Command("task", "export")
	stdout, err := cmd.Output()

	if err != nil {
		return err
	}

	err = json.Unmarshal(stdout, &tw.Tasks)

	return err
}

func (tw *Taskwarrior) Pending() []models.Task {
	var pendingTasks []models.Task
	for _, task := range tw.Tasks {
		if task.Status != "completed" {
			pendingTasks = append(pendingTasks, task)
		}
	}

	sort.Slice(pendingTasks, func(i, j int) bool {
		return pendingTasks[i].Urgency > pendingTasks[j].Urgency
	})

	return pendingTasks
}

func (tw *Taskwarrior) Completed() []models.Task {
	var completedTasks []models.Task
	for _, task := range tw.Tasks {
		if task.Status == "completed" {
			completedTasks = append(completedTasks, task)
		}
	}

	return completedTasks
}

func (tw *Taskwarrior) Projects() []string {
	m := make(map[string]int)
	for _, task := range tw.Tasks {
		if v, ok := task.Project.Value(); ok {
			m[v] = m[v] + 1
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	return keys
}

func (tw *Taskwarrior) Tags() []string {
	m := make(map[string]int)
	for _, task := range tw.Tasks {
		for _, tag := range task.Tags {
			m[tag] = m[tag] + 1
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	return keys
}
