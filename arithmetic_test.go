package golali_test

import (
	"testing"
	"time"

	"github.com/bijanghanei/golali"
)

func TestAddDays(t *testing.T) {
	j := golali.Date(1403, golali.Esfand, 30, 0, 0, 0, 0, time.UTC)
	j2 := j.AddDays(1)
	if j2.Month() != golali.Farvardin || j2.Day() != 1 || j2.Year() != 1404 {
		t.Errorf("AddDays(1) from 1403/12/29 = %v, want 1404/01/01", j2)
	}
}

func TestAddYearsLeap(t *testing.T) {
	j := golali.Date(1403, golali.Esfand, 30, 0, 0, 0, 0, time.UTC) // Leap year
	j2 := j.AddYears(1)
	if j2.Year() != 1404 || j2.Month() != golali.Esfand || j2.Day() != 29 { // Next year not leap
		t.Errorf("AddYears(1) from leap Esfand 30 = %v, want 1404/12/29", j2)
	}
}

func TestAddMonthsOverflow(t *testing.T) {
	j := golali.Date(1403, golali.Bahman, 30, 0, 0, 0, 0, time.UTC)
	j2 := j.AddMonths(2)
	if j2.Month() != golali.Farvardin || j2.Day() != 30 { // clamped
		t.Errorf("AddMonths(2) from 30 Bahman = %v, want Farvardin 30", j2)
	}
}

func TestDaysInBetween(t *testing.T) {
	a := golali.Date(1403, golali.Farvardin, 1, 0, 0, 0, 0, time.UTC)
	b := golali.Date(1403, golali.Farvardin, 11, 0, 0, 0, 0, time.UTC)
	if days := a.DaysInBetween(b); days != 10 {
		t.Errorf("DaysInBetween = %d, want 10", days)
	}
	if days := b.DaysInBetween(a); days != 10 { // order independent
		t.Errorf("DaysInBetween reverse = %d, want 10", days)
	}
}