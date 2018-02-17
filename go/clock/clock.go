package clock

import (
	"fmt"
)

// Clock represents hours and minutes in 24 hour time format
type Clock struct {
	hours   int
	minutes int
}

// New creates a clock with minutes and hours formatted
func New(h, m int) Clock {
	hours, minutes := formatTime(h, m)

	return Clock{hours, minutes}
}

// Add takes in the minutes to be added to a clock
func (clock Clock) Add(time int) Clock {
	clock.hours, clock.minutes = formatTime(clock.hours, clock.minutes+time)

	return clock
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hours, clock.minutes)
}

func formatTime(h, m int) (int, int) {
	extraHours := m / 60
	minutes := m % 60
	hours := (h + extraHours) % 24

	if minutes < 0 {
		hours--
		minutes = 60 + minutes
	}

	if hours < 0 {
		hours = 24 + hours
	}

	return hours, minutes
}
