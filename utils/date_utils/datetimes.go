package date_utils

import "time"

const datetimeLayout = "2006-01-02T15:04:05Z"

func GetNowTime() string {
	now := time.Now().UTC()
	return now.Format(datetimeLayout)
}
