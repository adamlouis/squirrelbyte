package present

import (
	"time"
)

// ToAPITime returns an RFC3339 time from a golang time
func ToAPITime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// ToInternalTime returns a golang time from a RFC3339
func ToInternalTime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}
