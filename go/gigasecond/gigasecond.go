package gigasecond

import (
	"time"
)

const gigasecond = 1000000000

// AddGigasecond gets a time and calcultes the date when a gigasecond will
// elapse.
func AddGigasecond(t time.Time) time.Time {
	newDate := t.Add(time.Second * time.Duration(gigasecond))
	return newDate
}
