package tmpl

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	t1 := time.Date(2020, 10, 20, 19, 55, 9, 0, time.UTC)
	now = func() time.Time {
		return t1
	}

	got := Now()

	if got != t1 {
		t.Errorf("Now() unexpected value:\n got:%+v\n exp: %+v", got, t1)

	}
}
