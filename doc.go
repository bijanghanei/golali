// Package jalali provides a comprehensive implementation of the Persian
// (Jalali/Shamsi) calendar with full support for date/time operations,
// conversion to/from Gregorian, formatting, parsing, and time zones.
//
// It is designed to be accurate, fast, and safe for use in production systems.
//
// Features:
//   - Bidirectional conversion between Jalali and Gregorian calendars
//   - Full time.Time compatibility (ToTime(), Now(), Add(), etc.)
//   - Custom formatting with Persian month/weekday names
//   - Strict parsing with layout validation
//   - Leap year detection
//   - Time zone support including IRST (Asia/Tehran)
//
// Example:
//
//	now := jalali.Now()
//	fmt.Println(now.Format("%Y/%m/%d")) // e.g., 1404/09/27
//
//	birth := jalali.Date(1375, jalali.Farvardin, 1, 0, 0, 0, 0, jalali.IRST())
//	fmt.Println(birth.Format("%B %d, %Y")) // فروردین 01, 1375
//
// The implementation follows algorithms from "Calendrical Calculations"
// by Nachum Dershowitz and Edward M. Reingold.
package golali