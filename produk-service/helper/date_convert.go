package helper

import (
	"time"
)

func ConvertDateTime(now string) string {
	convertTime, errTime := time.Parse(time.RFC3339, now)
	dataTime := ""
	if errTime != nil {

	} else {
		if convertTime.Format("2006-01-02 15:04:05") == "0001-01-01 00:00:00" || convertTime.Format("2006-01-02 15:04:05") == "1900-01-01 00:00:00" {
			dataTime = ""
		} else {
			dataTime = convertTime.Format("2006-01-02 15:04:05")
		}
	}
	return dataTime
}

func ConvertDate(now string) string {
	date, errDate := time.Parse(time.RFC3339, now)
	timeReturn := ""
	// fmt.Println(date.Format("2006-01-02"), errDate, "coba Convert")
	if errDate != nil {

	} else {
		if date.Format("2006-01-02") == "0001-01-01" || date.Format("2006-01-02") == "1900-01-01" {

		} else {
			timeReturn = date.Format("2006-01-02")

		}
	}
	return timeReturn
}
