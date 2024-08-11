package booking

import (
	"errors"
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(errors.New(err.Error()))
	}

	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	now := time.Now()
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(errors.New(err.Error()))
	}

	return now.After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(errors.New(err.Error()))
	}

	h := t.Hour()
	return h >= 12 && h < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	formatLayout := "Monday, January 2, 2006, at 15:04"
	parseLayout := "1/2/2006 15:04:05"
	t, err := time.Parse(parseLayout, date)
	if err != nil {
		panic(errors.New(err.Error()))
	}

	return fmt.Sprintf("You have an appointment on %s.", t.Format(formatLayout))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	t := time.Now()
	return time.Date(t.Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
