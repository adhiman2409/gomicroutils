package timex

import (
	"fmt"
	"time"
)

func GetLocalTime(timeZone string) (time.Time, error) {
	utcTime := time.Now().UTC()
	backupTimeZone := "Asia/Kolkata" // Default time zone
	if timeZone != "" {
		backupTimeZone = timeZone
	}
	loc, err := time.LoadLocation(backupTimeZone)
	if err != nil {
		fmt.Println("Error loading location:", err)
		return time.Time{}, fmt.Errorf("invalid time zone: %s", backupTimeZone)
	}

	return utcTime.In(loc), nil
}

func GetLocalDayMonthYear(timeZone string) (string, string, string, error) {
	utcTime := time.Now().UTC()
	backupTimeZone := "Asia/Kolkata" // Default time zone
	if timeZone != "" {
		backupTimeZone = timeZone
	}
	loc, err := time.LoadLocation(backupTimeZone)
	if err != nil {
		fmt.Println("Error loading location:", err)
		return "", "", "", fmt.Errorf("invalid time zone: %s", backupTimeZone)
	}

	t := utcTime.In(loc)
	day := fmt.Sprintf("%02d", t.Day())
	month := fmt.Sprintf("%02d", t.Month())
	year := fmt.Sprintf("%04d", t.Year())

	return day, month, year, nil
}
