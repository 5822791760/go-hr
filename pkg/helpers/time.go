package helpers

import "time"

func CurrentTime() string {
	currentTime := time.Now()
	return currentTime.Format(time.RFC3339)
}
