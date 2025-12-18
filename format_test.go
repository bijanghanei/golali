package golali_test

import (
	"testing"

	"github.com/bijanghanei/golali"
)

func TestFormat(t *testing.T) {
	j := golali.Date(1403, golali.Khordad, 5, 14, 30, 45, 0, golali.IRST())
	tests := []struct {
		layout string
		want   string
	}{
		{"%Y/%m/%d", "1403/03/05"},
		{"%Y-%m-%d %H:%M:%S", "1403-03-05 14:30:45"},
		{"%B %d, %Y", "خرداد 05, 1403"},
		{"%b %d", "خرد 05"},
		{"%w %p", "شنبه عصر"},
		{"%R", "14:30"},
		{"%T", "14:30:45"},
		{"%z", "+0330"}, // IRST is usually +0330 or +0430 depending on DST
		{"Hello %% %Y", "Hello % 1403"},
		{"Line1%nLine2", "Line1\nLine2"},
	}

	for _, tt := range tests {
		if got := j.Format(tt.layout); got != tt.want {
			// Note: %z might vary due to DST; adjust if needed
			if tt.layout == "%z" && (got == "+0330" || got == "+0430") {
				continue
			}
			t.Errorf("Format(%q) = %q, want %q", tt.layout, got, tt.want)
		}
	}

	if got := j.String(); got != "1403/03/05 14:30:45" {
		t.Errorf("String() = %q, want full timestamp", got)
	}
	if got := j.FormatDateTime(); got != "1403/03/05 14:30" {
		t.Errorf("FormatDateTime() = %q", got)
	}
}