package golali_test

import (
	"fmt"

	"github.com/bijanghanei/golali"
)

func ExampleNow() {
	now := golali.Now()
	fmt.Println(now.Format("%Y/%m/%d %R"))
	// Output:
	// 1404/09/27 12:34 (example; actual depends on current time)
}

func ExampleDate() {
	d := golali.Date(1403, golali.Khordad, 5, 14, 30, 0, 0, golali.IRST())
	fmt.Println(d.Format("%B %d, %Y - %w"))
	// Output:
	// خرداد 05, 1403 - شنبه
}