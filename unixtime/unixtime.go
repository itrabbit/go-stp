package unixtime

import (
	"strconv"
	"time"
)

// Time - Time in seconds (UNIX Timestamp)
type Time int64

// Time (UnixTime) - UnixTime to time.Time
func (t Time) Time() time.Time {
	return time.Unix(int64(t), 0)
}

// String (UnixTime) - UnixTime to String
func (t Time) String() string {
	return strconv.FormatInt(int64(t), 10)
}

// Format (UnixTime) - UnixTime to String By Format
func (t Time) Format(layout string) string {
	return t.Time().UTC().Format(layout)
}

// AddDate (UnixTime) - Add Date
func (t Time) AddDate(years, months, days int) Time {
	return Time(t.Time().AddDate(years, months, days).Unix())
}

// Add (UnixTime) - Add Date
func (t Time) Add(d time.Duration) Time {
	return Time(t.Time().Add(d).Unix())
}

// Now - Create Unix Time from Now time
func Now() Time {
	return Time(time.Now().Unix())
}
