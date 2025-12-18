package golali

import "time"

// IRST returns the Asia/Tehran location.
func IRST() *time.Location {
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		panic(err)
	}
	return loc
}

// UTC returns the JalaliDateTime in UTC.
func (j JalaliDateTime) UTC() JalaliDateTime {
	return ToJalaliDateTime(j.ToTime().UTC())
}

// Local returns the JalaliDateTime in local time zone.
func (j JalaliDateTime) Local() JalaliDateTime {
	return JalaliDateTime{
		year:     j.year,
		month:    j.month,
		day:      j.day,
		hour:     j.hour,
		min:      j.min,
		sec:      j.sec,
		nanosec:  j.nanosec,
		location: time.Local,
	}
}

// In returns the JalaliDateTime in the specified time zone.
func (j JalaliDateTime) In(loc *time.Location) JalaliDateTime {
	return ToJalaliDateTime(j.ToTime().In(loc))
}

// Zone returns the time zone name and offset.
func (j JalaliDateTime) Zone() (string, int) {
	return j.ToTime().Zone()
}