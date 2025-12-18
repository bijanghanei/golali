package golali

import "time"

// Add adds a duration to the JalaliDateTime.
func (j JalaliDateTime) Add(d time.Duration) JalaliDateTime {
	return ToJalaliDateTime(j.ToTime().Add(d))
}

// Sub returns the duration between two JalaliDateTimes.
func (j JalaliDateTime) Sub(u JalaliDateTime) time.Duration {
	return j.ToTime().Sub(u.ToTime())
}

// After returns true if j is after u.
func (x JalaliDateTime) After(y JalaliDateTime) bool {
	return x.UnixNano() > y.UnixNano()
}

// Before returns true if j is before u.
func (x JalaliDateTime) Before(y JalaliDateTime) bool {
	return x.UnixNano() < y.UnixNano()
}

// Unix returns the Unix timestamp.
func (jdt JalaliDateTime) Unix() int64 {
	return jdt.ToTime().Unix()
}

// UnixNano returns the Unix nano timestamp (UTC).
func (jdt JalaliDateTime) UnixNano() int64 {
	return jdt.ToTime().UTC().UnixNano()
}

// DaysInBetween returns the number of days between two dates.
func (s JalaliDateTime) DaysInBetween(e JalaliDateTime) int {
	if e.Before(s) {
		s, e = e, s
	}
	unix1 := s.Unix()
	unix2 := e.Unix()
	return int((unix2 - unix1) / 86400)
}

// AddYears adds years, handling leap Esfand special case.
func (j JalaliDateTime) AddYears(n int) JalaliDateTime {
	updatedYear := j.year + n
	if updatedYear < 1 {
		return JalaliDateTime{}
	}
	if j.IsLeapJalaliYear() && j.month == Esfand && j.day == 30 {
		j.day = 29
	}
	return JalaliDateTime{
		year:     updatedYear,
		month:    j.month,
		day:      j.day,
		hour:     j.hour,
		min:      j.min,
		sec:      j.sec,
		nanosec:  j.nanosec,
		location: j.location,
	}
}

// AddMonths adds months, adjusting day if needed.
func (j JalaliDateTime) AddMonths(n int) JalaliDateTime {
	years := n / 12
	months := n % 12
	updatedYear := j.year + years
	updatedMonth := int(j.month) + months
	if updatedMonth > 12 {
		updatedYear++
		updatedMonth -= 12
	}
	days := daysInMonth(updatedYear, Month(updatedMonth))
	if j.day > days {
		j.day = days
	}
	return JalaliDateTime{
		year:     updatedYear,
		month:    Month(updatedMonth),
		day:      j.day,
		hour:     j.hour,
		min:      j.min,
		sec:      j.sec,
		nanosec:  j.nanosec,
		location: j.location,
	}
}

// AddDays adds days using Gregorian equivalent.
func (j JalaliDateTime) AddDays(n int) JalaliDateTime {
	t := j.ToTime().AddDate(0, 0, n)
	return ToJalaliDateTime(t)
}