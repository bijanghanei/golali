package golali

import (
	"fmt"
	"strings"
)

// Format returns a formatted string according to the layout.
func (j JalaliDateTime) Format(layout string) string {
	var builder strings.Builder
	length := len(layout)
	i := 0

	for i < length {
		if layout[i] == '%' && i+1 < length {
			specifier := layout[i : i+2]
			switch specifier {
			case "%n":
				builder.WriteByte('\n')
			case "%%":
				builder.WriteByte('%')
			case "%Y":
				builder.WriteString(fmt.Sprintf("%04d", j.year))
			case "%y":
				builder.WriteString(fmt.Sprintf("%02d", j.year%100))
			case "%m":
				builder.WriteString(fmt.Sprintf("%02d", j.month))
			case "%B":
				builder.WriteString(FaJalaliMonthName[j.month])
			case "%b":
				name := FaJalaliMonthName[j.month]
				if len(name) >= 3 {
					s := string([]rune(name)[:3])
					builder.WriteString(s)
				} else {
					builder.WriteString(name)
				}
			case "%d":
				builder.WriteString(fmt.Sprintf("%02d", j.day))
			case "%H":
				builder.WriteString(fmt.Sprintf("%02d", j.hour))
			case "%M":
				builder.WriteString(fmt.Sprintf("%02d", j.min))
			case "%S":
				builder.WriteString(fmt.Sprintf("%02d", j.sec))
			case "%p":
				if j.hour < 12 {
					builder.WriteString("صبح")
				} else {
					builder.WriteString("عصر")
				}
			case "%w":
				builder.WriteString(FaWeekDays[j.Weekday()])
			case "%z":
				_, offset := j.Zone()
				sign := "+"
				if offset < 0 {
					sign = "-"
					offset = -offset
				}
				hours := offset / 3600
				minutes := (offset % 3600) / 60
				builder.WriteString(fmt.Sprintf("%s%02d%02d", sign, hours, minutes))
			case "%Z":
				if j.location != nil {
					builder.WriteString(j.location.String())
				}
			case "%R":
				builder.WriteString(fmt.Sprintf("%02d:%02d", j.hour, j.min))
			case "%T":
				builder.WriteString(fmt.Sprintf("%02d:%02d:%02d", j.hour, j.min, j.sec))
			default:
				builder.WriteString(specifier)
			}
			i += 2
		} else {
			builder.WriteByte(layout[i])
			i++
		}
	}

	return builder.String()
}

// FormatDateTime returns formatted date time as %Y/%m/%d %R
func (jdt JalaliDateTime) FormatDateTime() string {
	return jdt.Format("%Y/%m/%d %R")
}

// String returns formatted as %Y/%m/%d %T
func (jdt JalaliDateTime) String() string {
	return jdt.Format("%Y/%m/%d %T")
}