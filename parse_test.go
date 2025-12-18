package golali_test

import (
	"testing"

	"github.com/bijanghanei/golali"
)

func TestParse(t *testing.T) {
	loc := golali.IRST()

	tests := []struct {
		layout string
		value  string
		year   int
		month  golali.Month
		day    int
		hour   int
		min    int
		sec    int
	}{
		{"YYYY/MM/DD HH:MM:SS", "1403/06/15 14:30:00", 1403, golali.Shahrivar, 15, 14, 30, 0},
		{"YYYY-MM-DD HH:MM", "1402-08-20 09:45", 1402, golali.Aban, 20, 9, 45, 0},
		{"YYYY/MM/DD", "1401/12/29", 1401, golali.Esfand, 29, 0, 0, 0},
	}

	for _, tt := range tests {
		parsed, err := golali.ParseInLocation(tt.layout, tt.value, loc)
		if err != nil {
			t.Fatal(err)
		}
		if parsed.Year() != tt.year || parsed.Month() != tt.month || parsed.Day() != tt.day ||
			parsed.Hour() != tt.hour || parsed.Minute() != tt.min || parsed.Second() != tt.sec {
			t.Errorf("Parse(%q, %q) = %v, want year=%d month=%d day=%d time=%02d:%02d:%02d",
				tt.layout, tt.value, parsed, tt.year, tt.month, tt.day, tt.hour, tt.min, tt.sec)
		}
		if parsed.Location() != loc {
			t.Errorf("Location not preserved")
		}
	}
}

func TestParseErrors(t *testing.T) {
	badCases := []struct {
		layout, value string
	}{
		{"YYYY/MM/DD", "1403/13/01"},       // invalid month
		{"YYYY/MM/DD", "1403/06/32"},       // invalid day
		{"YYYY/MM/DD HH:MM:SS", "1403/06/15 24:00:00"}, // invalid hour
		{"YYYY/MM/DD", "1403-06-15"},       // separator mismatch
		{"YYYY/MM/DD", "1403/06/"},         // incomplete
	}

	for _, tt := range badCases {
		_, err := golali.Parse(tt.layout, tt.value)
		if err == nil {
			t.Errorf("Parse(%q, %q) should fail but succeeded", tt.layout, tt.value)
		}
	}
}