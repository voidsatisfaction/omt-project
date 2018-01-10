package utils

import (
	"testing"
	"time"
)

func TestGetTimerId(t *testing.T) {
	tests := []struct {
		expect string
		time   time.Time
	}{
		{"23:00", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)},
		{"2:34", time.Date(2009, time.November, 10, 2, 34, 0, 0, time.UTC)},
	}

	for _, test := range tests {
		got := GetTimerId(test.time)
		if test.expect != got {
			t.Errorf("GetTimer Id is invalid")
			t.Errorf("Expect %+v, got %+v\n", test.expect, got)
		}
	}
}
