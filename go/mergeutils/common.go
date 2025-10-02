package mergeutils

import (
	"bytes"
	"html/template"
	"time"
)

func MergeTemplate(temp string, data any) (*bytes.Buffer, error) {
	t, err := template.New("delivrio").Parse(temp)
	if err != nil {
		return nil, err
	}

	var msg bytes.Buffer
	err = t.Execute(&msg, data)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}

// Offset = 1 = "today at midnight"
func FutureMidnight(dayOffset time.Duration) time.Time {
	currentTime := time.Now()
	future := currentTime.Add(24 * time.Hour * dayOffset)
	midnight := time.Date(
		future.Year(),
		future.Month(),
		future.Day(),
		0,
		0,
		0,
		0,
		future.Location(),
	)
	return midnight
}
