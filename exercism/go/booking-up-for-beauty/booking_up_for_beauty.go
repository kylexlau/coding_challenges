package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/02/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// Friday, March 8, 1974 12:02:02
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	h := int(t.Hour())
	if h >= 12 && h < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	ft := t.Format("Monday, January 2, 2006, at 15:04")

	return fmt.Sprintf("You have an appointment on %s.", ft)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	anniversaryMonth := time.September
	anniversaryDay := 15
	currentYear := time.Now().Year()
	anniversary := time.Date(currentYear, anniversaryMonth, anniversaryDay, 0, 0, 0, 0, time.UTC)
	return anniversary
}
