package ghafpkg

import (
	"time"
)

func TimeGetDetailsWithPointer(paramDate *time.Time) (*int, *int, *int) {
	var yearExport *int
	var monthExport *int
	var dayExport *int

	var year int
	var day int
	var monthNumber int
	var monthName time.Month

	year, monthName, day = paramDate.Date()
	monthNumber = int(monthName)

	yearExport = &year
	monthExport = &monthNumber
	dayExport = &day

	return yearExport, monthExport, dayExport
}

func TimeGetFullDetails(paramDate time.Time) (*int, *int, *int, *int, *int, *int) {
	var yearExport *int
	var monthExport *int
	var dayExport *int
	var hourExport *int
	var minuteExport *int
	var secondExport *int

	var year int
	var day int
	var monthNumber int
	var monthName time.Month

	year, monthName, day = paramDate.Date()
	monthNumber = int(monthName)

	yearExport = &year
	monthExport = &monthNumber
	dayExport = &day

	hour := paramDate.Hour()
	minute := paramDate.Minute()
	second := paramDate.Second()

	hourExport = &hour
	minuteExport = &minute
	secondExport = &second

	return yearExport, monthExport, dayExport, hourExport, minuteExport, secondExport
}
