package datetime

import (
	"fmt"
	"time"
)

type Time time.Time

func Now() Time {
	return Time(time.Now().UTC())
}

func (t Time) MarshalJSON() ([]byte, error) {
	timeStr := fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RFC3339))
	return []byte(timeStr), nil
}
