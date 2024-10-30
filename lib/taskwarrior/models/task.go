package models

import (
	"fmt"

	"github.com/easilok/uiwarrior/lib/taskwarrior/types"
)

type Task struct {
	Uuid        string       `json:"uuid"`
	Id          uint32       `json:"id"`
	Description string       `json:"description"`
	Start       types.TWTime `json:"start"`
	End         types.TWTime `json:"end"`
	Due         types.TWTime `json:"due"`
	Wait        types.TWTime `json:"wait"`
	Schedule    types.TWTime `json:"scheduled"`
	Entry       types.TWTime `json:"entry"`
	Modified    types.TWTime `json:"modified"`
	Project     Project      `json:"project"`
	Status      Status       `json:"status"`
	Urgency     float32      `json:"urgency"`
	Tags        []string     `json:"tags"`
}

func (t *Task) Text() string {
	project := "-"
	if v, ok := t.Project.Value(); ok {
		project = v
	}
	due := t.Due.Date()
	if len(due) == 0 {
		due = "-"
	}
	active := ""
	if _, ok := t.Start.Value(); ok {
		if _, ok = t.End.Value(); !ok {
			active = " * "
		}
	}

	return fmt.Sprintf(
		"%d | %.1f | %s%s | %s | %s | D: %s |",
		t.Id, t.Urgency, t.Description, active, project, t.Status, due,
	)
}
