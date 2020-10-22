package tmpl

import (
	"time"
)

// for testing
var now = time.Now

// Now returns the current time.
func Now() time.Time {
	return now()
}
