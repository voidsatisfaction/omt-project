package utils

import "time"

func JapanTimeNow() time.Time {
	location := time.FixedZone("Asia/Tokyo", 9*60*60)
	return time.Now().In(location)
}
