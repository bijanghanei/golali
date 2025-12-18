package golali_test

import (
	"testing"
	"time"

	"github.com/bijanghanei/golali"
)

func TestGregorianToJalaliKnownDates(t *testing.T) {
	g := time.Date(2024, time.March, 20, 0, 0, 0, 0, time.UTC)
	back := golali.ToJalaliDateTime(g)
	j := golali.Date(1403, golali.Farvardin, 1, 0, 0, 0, 0, time.UTC)

	if back.Year() != j.Year() || back.Month() != j.Month() || back.Day() != j.Day() {
		t.Fatalf("gregorianToJalali(%d,%d,%d) = (%d,%d,%d), want (%d,%d,%d)",
			g.Year(), g.Month(), g.Day(),
			back.Year(), back.Month(), back.Day(),
			j.Year(), j.Month(), j.Day())
	}
}

func TestJalaliToGregorianKnownDates(t *testing.T) {
	g := time.Date(2024, time.March, 20, 0, 0, 0, 0, time.UTC)
	j := golali.Date(1403, golali.Farvardin, 1, 0, 0, 0, 0, time.UTC)
	back := j.ToTime()

	if back.Year() != g.Year() || back.Month() != g.Month() || back.Day() != g.Day() {
		t.Fatalf("jalaliToGregorian(%d,%d,%d) = (%d,%d,%d), want (%d,%d,%d)",
			j.Year(), j.Month(), j.Day(),
			back.Year(), back.Month(), back.Day(),
			g.Year(), g.Month(), g.Day())
	}
}

func TestRoundTripConversion(t *testing.T) {
	j := golali.Date(1403, golali.Khordad, 15, 12, 30, 45, 123456789, time.UTC)
	g := j.ToTime()
	back := golali.ToJalaliDateTime(g)

	if back.Year() != j.Year() || back.Month() != j.Month() || back.Day() != j.Day() ||
		back.Hour() != j.Hour() || back.Minute() != j.Minute() || back.Second() != j.Second() {
		t.Fatalf("round-trip failed: original %v, back %v", j, back)
	}
}