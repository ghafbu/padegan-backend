package ghafpkg

import (
	"time"
)

type DateDetails struct {
	Year  *int `json:"year"`
	Month *int `json:"month"`
	Day   *int `json:"day"`
}

type DateDetailsWithTime struct {
	Year   *int `json:"year"`
	Month  *int `json:"month"`
	Day    *int `json:"day"`
	Hour   *int `json:"hour"`
	Minute *int `json:"minute"`
	Second *int `json:"Second"`
}

func TimeGenerateStructDetailsWithPointer(paramDate *time.Time, paramDetails *DateDetails) *DateDetails {
	paramDetails.Year,
		paramDetails.Month,
		paramDetails.Day = TimeGetDetailsWithPointer(paramDate)

	return paramDetails
}

func TimeGenerateStructFullDetails(paramDate time.Time, paramDetails *DateDetailsWithTime) *DateDetailsWithTime {
	paramDetails.Year,
		paramDetails.Month,
		paramDetails.Day,
		paramDetails.Hour,
		paramDetails.Minute,
		paramDetails.Second = TimeGetFullDetails(paramDate)

	return paramDetails
}

func TimeGenerateStructFullDetailsWithoutPointer(paramDate time.Time, paramDetails DateDetailsWithTime) DateDetailsWithTime {
	paramDetails.Year,
		paramDetails.Month,
		paramDetails.Day,
		paramDetails.Hour,
		paramDetails.Minute,
		paramDetails.Second = TimeGetFullDetails(paramDate)

	return paramDetails
}
