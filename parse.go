package golali

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Parse parses the value according to layout in local time.
func Parse(layout, value string) (JalaliDateTime, error) {
	return ParseInLocation(layout, value, time.Local)
}

// ParseInLocation parses the value according to layout in the given location.
func ParseInLocation(layout, value string, location *time.Location) (JalaliDateTime, error) {
	if len(layout) != len(value) {
		return JalaliDateTime{}, errors.New("value does not match the layout")
	}

	layoutParts, layoutSeps, err := tokenize(layout)
	if err != nil {
		return JalaliDateTime{}, fmt.Errorf("invalid layout: %v", err)
	}
	valueParts, valueSeps, err := tokenize(value)
	if err != nil {
		return JalaliDateTime{}, fmt.Errorf("invalid value: %v", err)
	}

	if len(layoutParts) != len(valueParts) {
		return JalaliDateTime{}, errors.New("layout and value have mismatch structure")
	}

	for i, ls := range layoutSeps {
		if ls != valueSeps[i] {
			return JalaliDateTime{}, fmt.Errorf("separator mismatch at position %d: expected %c, got %c", i, ls, valueSeps[i])
		}
	}

	var year, month, day, hour, min, sec int
	for i, part := range layoutParts {
		num, err := strconv.Atoi(valueParts[i])
		if err != nil {
			return JalaliDateTime{}, fmt.Errorf("invalid value for part << %v >> : %v", part, err)
		}
		switch part {
		case "YYYY":
			if num < 1 || num > 9999 {
				return JalaliDateTime{}, errors.New("year out of range (1-9999)")
			}
			year = num
		case "MM":
			if i > 0 && layoutParts[i-1] == "HH" {
				if num < 0 || num > 59 {
					return JalaliDateTime{}, errors.New("minute out of range (0-59)")
				}
				min = num
			} else {
				if num < 1 || num > 12 {
					return JalaliDateTime{}, errors.New("month out of range (1-12)")
				}
				month = num
			}
		case "DD":
			if num < 1 || num > 31 {
				return JalaliDateTime{}, errors.New("day out of range (1-31)")
			}
			day = num
		case "HH":
			if num < 0 || num > 23 {
				return JalaliDateTime{}, errors.New("hour out of range (0-23)")
			}
			hour = num
		case "SS":
			if num < 0 || num > 59 {
				return JalaliDateTime{}, errors.New("seconds out of range (0-59)")
			}
			sec = num
		default:
			return JalaliDateTime{}, fmt.Errorf("invalid layout token at position %d: %s", i, part)
		}
	}
	return JalaliDateTime{
		year:     year,
		month:    Month(month),
		day:      day,
		hour:     hour,
		min:      min,
		sec:      sec,
		nanosec:  0,
		location: location,
	}, nil
}

func tokenize(s string) (parts []string, seps []rune, err error) {
	var sb strings.Builder
	for i, r := range s {
		if r == '/' || r == ':' || r == '-' || r == ' ' {
			if sb.Len() > 0 {
				parts = append(parts, sb.String())
				sb.Reset()
			} else if i > 0 {
				return nil, nil, errors.New("consecutive separators")
			}
			seps = append(seps, r)
		} else {
			sb.WriteRune(r)
		}
	}
	if sb.Len() > 0 {
		parts = append(parts, sb.String())
	}
	if len(parts) == 0 {
		return nil, nil, errors.New("empty input")
	}
	return parts, seps, nil
}