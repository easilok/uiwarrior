package models

import (
	"fmt"
	"time"

	"github.com/easilok/uiwarrior/utils"
)

type Task struct {
	Uuid        string               `json:"uuid"`
	Id          uint32               `json:"id"`
	Description string               `json:"description"`
	Start       utils.Option[string] `json:"start"`
	End         utils.Option[string] `json:"end"`
	Due         utils.Option[string] `json:"due"`
	Wait        utils.Option[string] `json:"wait"`
	Schedule    utils.Option[string] `json:"scheduled"`
	Entry       string               `json:"entry"`
	Modified    string               `json:"modified"`
	Project     utils.Option[string] `json:"project"`
	Status      string               `json:"status"`
	Urgency     float32              `json:"urgency"`
	Tags        []string             `json:"tags"`
}

func (t *Task) Text() string {
	project := "-"
	if v, ok := t.Project.Value(); ok {
		project = v
	}
	due := AsDateString(t.Due)
	if len(due) == 0 {
		due = "-"
	}
	urgency := int32(t.Urgency)
	active := ""
	if _, ok := t.Start.Value(); ok {
		if _, ok = t.End.Value(); !ok {
			active = " * "
		}
	}

	return fmt.Sprintf(
		"%d | %d | %s%s | %s | %s | D: %s |",
		t.Id, urgency, t.Description, active, project, t.Status, due,
	)
}

func AsDateString(d utils.Option[string]) string {
	if v, ok := d.Value(); ok {
		t, err := time.Parse("20060102T150405Z", v)
		if err == nil {
			return t.Format("2006-01-02")
		}
	}

	return ""
}

func AsTime(d utils.Option[string]) (time.Time, bool) {
	if v, ok := d.Value(); ok {
		t, err := time.Parse("20060102T150405Z", v)
		if err == nil {
			return t, true
		}
	}

	return time.Now(), false
}
