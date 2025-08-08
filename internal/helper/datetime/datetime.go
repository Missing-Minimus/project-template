package datetime

import (
	"strconv"
	"time"
)

func GetCurrentDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func ParseTimestampToTime(timestamp string) (time.Time, error) {
	parsedTimestamp, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(parsedTimestamp, 0), nil
}
