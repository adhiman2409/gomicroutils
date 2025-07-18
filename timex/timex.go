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
