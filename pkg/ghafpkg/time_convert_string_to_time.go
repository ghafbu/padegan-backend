package ghafpkg

import (
	"errors"
	"time"
)

func TimeConvertStringToTimeByFormat(paramBirthDate string, paramFormat string) (time.Time, error) {
	var thisBirthDate, err = time.Parse(paramFormat, paramBirthDate)
	if err != nil {
		return thisBirthDate, errors.New("birthdate_parse_error")
	}

	return thisBirthDate, nil
}

func TimeConvertStringToTime(paramBirthDate string) (time.Time, error) {
	// تجزیه تاریخ با فرمت RFC3339
	thisBirthDate, err := time.Parse(time.RFC3339, paramBirthDate)
	if err != nil {
		return thisBirthDate, errors.New("birthdate_parse_error")
	}

	return thisBirthDate, nil
}
