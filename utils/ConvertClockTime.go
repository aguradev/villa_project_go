package utils

import "time"

func ConvertClockTime(setTime string) time.Time {

	value, _ := time.Parse(time.Kitchen, setTime)

	return value.Local()

}
