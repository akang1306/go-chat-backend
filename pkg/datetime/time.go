package datetime

import (
	"fmt"
	"time"
)

type Time time.Time

func Now() Time {
	return Time(time.Now().UTC())
}

func (t Time) String() string {
	return time.Time(t).Format(time.RFC3339)
}

func Parse(timeStr string) (Time, error) {
	time, err := time.Parse(time.RFC3339, timeStr)
	return Time(time), err
}

func (t Time) MarshalJSON() ([]byte, error) {
	timeStr := fmt.Sprintf("\"%s\"", t.String())
	return []byte(timeStr), nil
}
