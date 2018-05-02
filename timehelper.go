package timehelper

import "time"

//TimeNowUTC get time now UTC
func TimeNowUTC() time.Time {
	return time.Now().UTC()
}

//TimeNowUTCString get time now UTC string
func TimeNowUTCString() string {
	return ConvertTimeToString(TimeNowUTC())
}

// ConvertUnixToTimeUTC unix -> time
func ConvertUnixToTimeUTC(unix int64) time.Time {
	return time.Unix(unix, 0).UTC()
}

// ConvertTimeToString time.Time -> string
func ConvertTimeToString(time time.Time) string {
	return time.UTC().Format("2006-01-02T15:04:05.999999Z")
}

// ConvertStringToTime string -> time.Time
func ConvertStringToTime(str string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05.999999Z", str)
}
