package rate_limit

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"sushi-backend/pkg/config"
	"sushi-backend/pkg/logger"
	"sync"
	"time"
)

type IPRateLimiter struct {
	logger logger.ILogger
	config config.IConfig
	ips    map[string]*RateLimiter
	mu     *sync.RWMutex
	rate   int
	burst  int
}

func NewIPRateLimiter(deps IpRateLimiterDependencies) *IPRateLimiter {
	r := &IPRateLimiter{
		logger: deps.Logger,
		config: deps.Config,
		ips:    make(map[string]*RateLimiter),
		mu:     &sync.RWMutex{},
		rate:   deps.Config.GetIpRateLimitRate(),
		burst:  deps.Config.GetIpRateLimitBurst(),
	}

	r.logger.Log(
		fmt.Sprintf(
			"IPRateLimiter created with rate %d and burst %d, cleanup interval %s, IP expiration %s",
			r.rate,
			r.burst,
			deps.Config.GetIpRateLimitCleanupInterval(),
			deps.Config.GetIpRateLimitExpiration(),
		),
	)

	go r.cleanup(deps.ShutdownContext)

	return r
}

func (i *IPRateLimiter) AddIP(ip string) *RateLimiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := NewRateLimiter(rate.Every(time.Second*time.Duration(i.rate)), i.burst)

	i.ips[ip] = limiter

	i.logger.Debug(fmt.Sprintf("added IP %s to the map, total size %d", ip, len(i.ips)))

	return limiter
}

func (i *IPRateLimiter) GetLimiter(ip string) *RateLimiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]

	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}

	i.mu.Unlock()

	return limiter
}

func (i *IPRateLimiter) GetActiveIpsCount() int {
	i.mu.RLock()
	defer i.mu.RUnlock()

	return len(i.ips)
}

func (i *IPRateLimiter) cleanup(ctx context.Context) {
	ticker := time.NewTicker(i.config.GetIpRateLimitCleanupInterval())
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			i.mu.Lock()
			for ip := range i.ips {
				delete(i.ips, ip)
			}
			i.logger.Debug(fmt.Sprintf("IPRateLimiter shutdown, total IPs size %d", len(i.ips)))
			i.mu.Unlock()
			return
		case <-ticker.C:
			i.mu.Lock()
			i.logger.Debug(fmt.Sprintf("Running Ip rate limit cleanup, total IPs size %d", len(i.ips)))
			for ip, limiter := range i.ips {
				if limiter.IsExpired(i.config.GetIpRateLimitExpiration()) {
					i.logger.Debug(fmt.Sprintf("Deleting expired IP %s from the map", ip))
					delete(i.ips, ip)
				}
			}
			i.mu.Unlock()
		}
	}
}
