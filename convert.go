package golali

import (
	"fmt"
	"time"
)

// gregorianToJalali converts a Gregorian date to a Jalali date
func gregorianToJalali(gYear int, gMonth time.Month, gDay int) (jYear int, jMonth Month, jDay int) {
	gy := gYear - 1600
	gm := int(gMonth) - 1
	gd := gDay - 1

	gDayNo := 365*gy + (gy+3)/4 - (gy+99)/100 + (gy+399)/400
	for i := 0; i < gm; i++ {
		gDayNo += gregorianDaysInMonth[i]
	}
	if gm > 1 && ((gy%4 == 0 && gy%100 != 0) || (gy%400 == 0)) {
		gDayNo++
	}
	gDayNo += gd

	jDayNo := gDayNo - 79
	jNp := jDayNo / 12053
	jDayNo %= 12053

	jYear = 979 + 33*jNp + 4*(jDayNo/1461)
	jDayNo %= 1461

	if jDayNo >= 366 {
		jYear += (jDayNo - 1) / 365
		jDayNo = (jDayNo - 1) % 365
	}

	var i int
	for i = 0; i < 11 && jDayNo >= jalaliDaysInMonth[i]; i++ {
		jDayNo -= jalaliDaysInMonth[i]
	}
	jMonth = Month(i + 1)
	jDay = jDayNo + 1

	return
}

// jalaliToGregorian converts a Jalali date to a Gregorian date
func jalaliToGregorian(jYear int, jMonth int, jDay int) (gYear int, gMonth int, gDay int) {
	jy := jYear - 979
	jm := jMonth - 1
	jd := jDay - 1

	jDayNo := 365*jy + (jy/33)*8 + (jy%33+3)/4
	for i := 0; i < jm; i++ {
		jDayNo += jalaliDaysInMonth[i]
	}
	jDayNo += jd

	gDayNo := jDayNo + 79
	gy := 1600 + 400*(gDayNo/146097)
	gDayNo %= 146097

	leap := true
	if gDayNo >= 36525 {
		gDayNo--
		gy += 100 * (gDayNo / 36524)
		gDayNo %= 36524

		if gDayNo >= 365 {
			gDayNo++
		} else {
			leap = false
		}
	}

	gy += 4 * (gDayNo / 1461)
	gDayNo %= 1461

	if gDayNo >= 366 {
		leap = false
		gDayNo--
		gy += gDayNo / 365
		gDayNo %= 365
	}

	var i int
	for i = 0; gDayNo >= gregorianDaysInMonth[i]+boolToInt(i == 1 && leap); i++ {
		gDayNo -= gregorianDaysInMonth[i] + boolToInt(i == 1 && leap)
	}
	gMonth = i + 1
	gDay = gDayNo + 1

	return gy, gMonth, gDay
}

// ToJalaliDateTime converts time.Time to JalaliDateTime
func ToJalaliDateTime(t time.Time) JalaliDateTime {
	jYear, jMonth, jDay := gregorianToJalali(t.Year(), t.Month(), t.Day())
	return JalaliDateTime{
		year:     jYear,
		month:    jMonth,
		day:      jDay,
		hour:     t.Hour(),
		min:      t.Minute(),
		sec:      t.Second(),
		nanosec:  t.Nanosecond(),
		location: t.Location(),
	}
}

// ToTime converts JalaliDateTime to time.Time
func (j JalaliDateTime) ToTime() time.Time {
	if j.location == nil {
		j.location = time.Local
	}
	gYear, gMonth, gDay := jalaliToGregorian(j.year, int(j.month), j.day)
	return time.Date(gYear, time.Month(gMonth), gDay, j.hour, j.min, j.sec, j.nanosec, j.location)
}

// Now returns the current JalaliDateTime.
func Now() JalaliDateTime {
	return ToJalaliDateTime(time.Now())
}

// Date returns a new JalaliDateTime value representing the given date and time.
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) JalaliDateTime {
	if year < 1 || year > 9999 {
		panic(fmt.Sprintf("year out of range: %d", year))
	}
	if month < Farvardin || month > Esfand {
		panic(fmt.Sprintf("invalid month: %v", month))
	}
	if day < 1 || day > daysInMonth(year, month) {
		panic(fmt.Sprintf("day out of range: %d", day))
	}
	if hour < 0 || hour > 23 {
		panic(fmt.Sprintf("hour out of range: %d", hour))
	}
	if min < 0 || min > 59 {
		panic(fmt.Sprintf("minute out of range: %d", min))
	}
	if sec < 0 || sec > 59 {
		panic(fmt.Sprintf("second out of range: %d", sec))
	}
	if nsec < 0 || nsec > 999999999 {
		panic(fmt.Sprintf("nanosecond out of range: %d", nsec))
	}

	return JalaliDateTime{
		year:     year,
		month:    month,
		day:      day,
		hour:     hour,
		min:      min,
		sec:      sec,
		nanosec:  nsec,
		location: loc,
	}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func daysInMonth(year int, month Month) int {
	if month < Farvardin || month > Esfand {
		panic(fmt.Sprintf("invalid month: %v", month))
	}
	if month <= Shahrivar {
		return 31
	}
	if month <= Bahman {
		return 30
	}
	if isLeapJalaliYear(year) {
		return 30
	}
	return 29
}

func isLeapJalaliYear(year int) bool {
	rm := year % 33
	leapYearsRemainders := map[int]bool{
		1:  true,
		5:  true,
		9:  true,
		13: true,
		17: true,
		22: true,
		26: true,
		30: true,
	}
	return leapYearsRemainders[rm]
}

// IsLeapJalaliYear returns true if the year is a leap year in the Jalali calendar.
func (j JalaliDateTime) IsLeapJalaliYear() bool {
	return isLeapJalaliYear(j.year)
}

// DaysInMonth returns the number of days in the month of the JalaliDateTime.
func (j JalaliDateTime) DaysInMonth() int {
	return daysInMonth(j.year, j.month)
}