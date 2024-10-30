package types

import (
	"encoding/json"
	"time"
)

type TWTime struct {
	time *time.Time
}

func (t *TWTime) UnmarshalJSON(bytes []byte) error {

	if len(bytes) < 12 {
		return nil
	}

	var value string
	err := json.Unmarshal(bytes, &value)
	if err != nil {
		return err
	}

	t.time = new(time.Time)
	*t.time, err = time.Parse("20060102T150405Z", value)

	return err
}

func (t *TWTime) Value() (out time.Time, ok bool) {
	if t.time == nil {
		return
	}
	return *t.time, true
}

func (t *TWTime) Date() string {
	if v, ok := t.Value(); ok {
		return v.Format("2006-01-02")
	}

	return ""
}
