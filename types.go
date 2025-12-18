package golali

import (
	"fmt"
	"time"
)

// gregorianDaysInMonth contains the number of days in each month in the Gregorian calendar.
var gregorianDaysInMonth = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// jalaliDaysInMonth contains the number of days in each month in the Jalali calendar.
var jalaliDaysInMonth = []int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}

// EnJalaliMonthName contains the names of the months in the Jalali calendar in English.
var EnJalaliMonthName = []string{
	"",
	"Farvardin", "Ordibehesht", "Khordad",
	"Tir", "Mordad", "Shahrivar",
	"Mehr", "Aban", "Azar",
	"Dey", "Bahman", "Esfand",
}

// FaJalaliMonthName contains the names of the months in the Jalali calendar in Persian.
var FaJalaliMonthName = []string{
	"",
	"فروردین", "اردیبهشت", "خرداد",
	"تیر", "مرداد", "شهریور",
	"مهر", "آبان", "آذر",
	"دی", "بهمن", "اسفند",
}

// JalaliMonthNumber contains the two-digit numbers of the months in the Jalali calendar.
var JalaliMonthNumber = []string{
	"",
	"01", "02", "03",
	"04", "05", "06",
	"07", "08", "09",
	"10", "11", "12",
}

// Month represents a month in the Jalali calendar.
type Month int

const (
	Farvardin Month = 1 + iota
	Ordibehesht
	Khordad
	Tir
	Mordad
	Shahrivar
	Mehr
	Aban
	Azar
	Dey
	Bahman
	Esfand
)

// String returns the English name of the month.
func (m Month) String() string {
	if m < 1 || m > 12 {
		panic(fmt.Sprintf("invalid month value: %v", int(m)))
	}
	return EnJalaliMonthName[m]
}

// FaString returns the Persian name of the month.
func (m Month) FaString() string {
	if m < 1 || m > 12 {
		panic(fmt.Sprintf("invalid month value: %v", int(m)))
	}
	return FaJalaliMonthName[m]
}

// FaWeekDays contains the names of the weekdays in Persian.
var FaWeekDays = []string{"یکشنبه", "دوشنبه", "سه‌شنبه", "چهارشنبه", "پنج‌شنبه", "جمعه", "شنبه"}

// EnWeekDays contains the names of the weekdays in English.
var EnWeekDays = []string{"1Shanbeh", "2Shanbeh", "3Shanbeh", "4Shanbeh", "5Shanbeh", "Joomeh", "Shanbeh"}

// Weekday represents a day of the week in the Jalali calendar.
type Weekday int

const (
	Yekshanbe Weekday = iota
	Doshanbe
	Seshanbe
	Chaharshanbe
	Panjshanbe
	Joomeh
	Shanbe
)

// String returns the English name of the weekday.
func (w Weekday) String() string {
	if w < 0 || w > 6 {
		panic(fmt.Sprintf("invalid weekday value: %v", int(w)))
	}
	return EnWeekDays[w]
}

// FaString returns the Persian name of the weekday.
func (w Weekday) FaString() string {
	if w < 0 || w > 6 {
		panic(fmt.Sprintf("invalid weekday value: %v", int(w)))
	}
	return FaWeekDays[w]
}

// JalaliDateTime represents a date and time in the Jalali calendar
type JalaliDateTime struct {
	year     int
	month    Month
	day      int
	hour     int
	min      int
	sec      int
	nanosec  int
	location *time.Location
}

// Location returns the location of the JalaliDateTime.
func (j JalaliDateTime) Location() *time.Location {
	return j.location
}

// Year returns the year of the Jalali date.
func (j JalaliDateTime) Year() int {
	return j.year
}

// Month returns the month of the Jalali date.
func (j JalaliDateTime) Month() Month {
	return j.month
}

// Day returns the day of the month of the Jalali date.
func (j JalaliDateTime) Day() int {
	return j.day
}

// Hour returns the hour of the Jalali time.
func (j JalaliDateTime) Hour() int {
	return j.hour
}

// Minute returns the minute of the Jalali time.
func (j JalaliDateTime) Minute() int {
	return j.min
}

// Second returns the second of the Jalali time.
func (j JalaliDateTime) Second() int {
	return j.sec
}

// Weekday returns the day of the week of the Jalali date.
func (j JalaliDateTime) Weekday() Weekday {
	gYear, gMonth, gDay := jalaliToGregorian(j.year, int(j.month), j.day)
	gDate := time.Date(gYear, time.Month(gMonth), gDay, j.hour, j.min, j.sec, j.nanosec, j.location)
	weekday := gDate.Weekday()
	switch weekday {
	case time.Sunday:
		return Yekshanbe
	case time.Monday:
		return Doshanbe
	case time.Tuesday:
		return Seshanbe
	case time.Wednesday:
		return Chaharshanbe
	case time.Thursday:
		return Panjshanbe
	case time.Friday:
		return Joomeh
	case time.Saturday:
		return Shanbe
	default:
		panic(fmt.Sprintf("invalid weekday value: %v", weekday))
	}
}