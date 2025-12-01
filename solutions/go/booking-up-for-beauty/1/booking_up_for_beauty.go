package booking

import "time"

const usLayout = "1/2/2006 15:04:05"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    t, _ := time.Parse(usLayout, date)
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    t, err := time.Parse("January 2, 2006 15:04:05", date)
    if err != nil {
        return false
    }
    return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	h := t.Hour()
	return h >= 12 && h < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, _ := time.Parse(usLayout, date)
	return "You have an appointment on " + t.Format("Monday, January 2, 2006, at 15:04") + "."

}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    year := time.Now().Year()
	return time.Date(year, time.September, 15, 0, 0, 0, 0, time.UTC)
}
