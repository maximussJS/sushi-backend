package rate_limit

import (
	"golang.org/x/time/rate"
	"time"
)

type RateLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

func NewRateLimiter(rateLimit rate.Limit, burst int) *RateLimiter {
	return &RateLimiter{
		limiter:  rate.NewLimiter(rateLimit, burst),
		lastSeen: time.Now(),
	}
}

func (r *RateLimiter) Allow() bool {
	r.lastSeen = time.Now()
	return r.limiter.Allow()
}

func (r *RateLimiter) IsExpired(expiration time.Duration) bool {
	return time.Since(r.lastSeen) > expiration
}
