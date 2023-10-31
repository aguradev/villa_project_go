package utils

import "time"

const (
	// YYYY-MM-DD: 2022-03-23
	YYYYMMDD = "2006-01-02"
	// 24h hh:mm:ss: 14:23:20
	HHMMSS24h = "15:04:05"
	// 12h hh:mm:ss: 2:23:20 PM
	HHMMSS12h = "3:04:05 PM"
	// text date: March 23, 2022
	TextDate = "January 2, 2006"
	// text date with weekday: Wednesday, March 23, 2022
	TextDateWithWeekday = "Monday, January 2, 2006"
	// abbreviated text date: Mar 23 Wed
	AbbrTextDate = "Jan 2 Mon"
)

func ConvertClockTime(setTime string) *time.Time {

	loc := time.Local

	value, err := time.ParseInLocation(HHMMSS24h, setTime, loc)

	if err != nil {
		return nil
	}

	return &value

}

func ConvertDate(date string) (*time.Time, error) {
	TimeParse, TimeErr := time.ParseInLocation("2006-01-02", date, time.Local)

	if TimeErr != nil {
		return nil, TimeErr
	}

	return &TimeParse, nil
}

func GetSubDate(start_date time.Time, end_date time.Time) int {

	days := int(end_date.Sub(start_date).Hours() / 24)

	return days

}
