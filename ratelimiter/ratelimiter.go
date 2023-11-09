package ratelimiter

import "time"

type RateLimiter struct {
	C chan bool
}

func New(maxRequestsPerMinute int) *RateLimiter {
	c := make(chan bool, maxRequestsPerMinute)

	for i := 0; i < maxRequestsPerMinute; i++ {
		go func() {
			for {
				<-time.After(time.Minute)
				<-c
			}
		}()
	}

	return &RateLimiter{
		C: c,
	}
}
