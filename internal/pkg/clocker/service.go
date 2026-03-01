package clocker

import (
	"time"
)

// Now returns current time in Asia/Makassar timezone (GMT+8)
func Now() time.Time {
	loc, _ := time.LoadLocation("Asia/Makassar")
	return time.Now().In(loc)
}

// Parse converts time to Asia/Makassar timezone while preserving the instant
func Parse(oldTime time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Makassar")
	return oldTime.In(loc)
}

// ParseWithoutCalculation changes timezone without converting the instant
// Example: 10:00 UTC becomes 10:00 Asia/Makassar
func ParseWithoutCalculation(oldTime time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Makassar")
	return time.Date(oldTime.Year(), oldTime.Month(), oldTime.Day(), oldTime.Hour(), oldTime.Minute(), oldTime.Second(), oldTime.Nanosecond(), loc)
}

// ParseToLocal parses time string in given format to Asia/Makassar timezone
func ParseToLocal(format string, oldTime string) (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Makassar")
	if err != nil {
		return time.Time{}, err
	}

	newTime, err := time.ParseInLocation(format, oldTime, loc)
	if err != nil {
		return time.Time{}, err
	}

	return newTime, nil
}

// StringToTime parses time string with given format
func StringToTime(f string, tf string) (time.Time, error) {
	t, err := time.Parse(f, tf)
	if err != nil {
		return Now(), err
	}

	return t, nil
}
