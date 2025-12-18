package golali_test

import (
	"testing"
	"time"

	"github.com/bijanghanei/golali"
)

func TestMonthString(t *testing.T) {
	tests := []struct {
		m    golali.Month
		en   string
		fa   string
	}{
		{golali.Farvardin, "Farvardin", "فروردین"},
		{golali.Ordibehesht, "Ordibehesht", "اردیبهشت"},
		{golali.Esfand, "Esfand", "اسفند"},
	}

	for _, tt := range tests {
		if got := tt.m.String(); got != tt.en {
			t.Errorf("Month(%d).String() = %q, want %q", tt.m, got, tt.en)
		}
		if got := tt.m.FaString(); got != tt.fa {
			t.Errorf("Month(%d).FaString() = %q, want %q", tt.m, got, tt.fa)
		}
	}
}

func TestWeekdayString(t *testing.T) {
	// 1403 Farvardin 1 = 2024 March 20 = Wednesday → Chaharshanbe in original mapping
	j := golali.Date(1403, golali.Farvardin, 1, 0, 0, 0, 0, time.UTC)
	w := j.Weekday()

	if got := w.String(); got != "4Shanbeh" {
		t.Errorf("Weekday.String() = %q, want \"4Shanbeh\"", got)
	}
	if got := w.FaString(); got != "چهارشنبه" {
		t.Errorf("Weekday.FaString() = %q, want \"چهارشنبه\"", got)
	}
}

func TestLeapYearViaMethod(t *testing.T) {
	leapYears := []int{1379, 1383, 1387, 1391, 1395, 1399,1403, 1408}
	nonLeapYears := []int{1380, 1381, 1382, 1384, 1402,1407, 1405,1404}

	for _, y := range leapYears {
		j := golali.Date(y, golali.Esfand, 1, 0, 0, 0, 0, time.UTC)
		if !j.IsLeapJalaliYear() {
			t.Errorf("Year %d should be leap", y)
		}
		if j.DaysInMonth() != 30 { // Esfand in leap year
			t.Errorf("Esfand in leap year %d should have 30 days, got %d", y, j.DaysInMonth())
		}
	}

	for _, y := range nonLeapYears {
		j := golali.Date(y, golali.Esfand, 1, 0, 0, 0, 0, time.UTC)
		if j.IsLeapJalaliYear() {
			t.Errorf("Year %d should NOT be leap", y)
		}
		if j.DaysInMonth() != 29 {
			t.Errorf("Esfand in non-leap year %d should have 29 days, got %d", y, j.DaysInMonth())
		}
	}
}