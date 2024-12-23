package rate_limit

type IIpRateLimiter interface {
	GetLimiter(ip string) *RateLimiter
	AddIP(ip string) *RateLimiter
	GetActiveIpsCount() int
}
