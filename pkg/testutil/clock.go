package testutil

import (
	"time"

	"github.com/5822791760/hr/pkg/coreutil"
)

type fakeClock struct {
	MockTime time.Time
}

func (c fakeClock) Now() time.Time {
	return c.MockTime
}

func (c fakeClock) After(d time.Duration) <-chan time.Time {
	ch := make(chan time.Time, 1)
	ch <- c.MockTime.Add(d)
	return ch
}

func GetFakeClock() (coreutil.Clock, time.Time) {
	now := time.Now()
	clock := fakeClock{MockTime: now}
	return clock, now
}
