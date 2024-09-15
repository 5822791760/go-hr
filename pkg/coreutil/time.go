package coreutil

import "time"

type Clock interface {
	Now() time.Time
	After(d time.Duration) <-chan time.Time
}

type realClock struct{}

func NewClock() realClock {
	return realClock{}
}

func (realClock) Now() time.Time                         { return time.Time{} }
func (realClock) After(d time.Duration) <-chan time.Time { return time.After(d) }
