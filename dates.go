package main

import (
	"time"
)

func getMondayDate(today time.Time) time.Time {
	switch today.Weekday() {
	case time.Monday:
		return today
	case time.Tuesday:
		return today.AddDate(0, 0, -1)
	case time.Wednesday:
		return today.AddDate(0, 0, -2)
	case time.Thursday:
		return today.AddDate(0, 0, -3)
	case time.Friday:
		return today.AddDate(0, 0, -4)
	case time.Saturday:
		return today.AddDate(0, 0, -5)
	case time.Sunday:
		return today.AddDate(0, 0, -6)
	}

	return today
}
