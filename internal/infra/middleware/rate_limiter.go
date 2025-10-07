package middleware

import (
	"net"
	"net/http"
	"time"

	"github.com/ulule/limiter/v3"
	stdlibmiddleware "github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
)

type RateLimiterConfig struct {
	Limit  int
	Period time.Duration
}

func NewRateLimiterMiddleware(cfg RateLimiterConfig) func(http.Handler) http.Handler {
	// In-memory store
	store := memory.NewStore()
	rate := limiter.Rate{
		Period: cfg.Period,
		Limit:  int64(cfg.Limit),
	}
	instance := limiter.New(store, rate)

	return stdlibmiddleware.NewMiddleware(instance,
		stdlibmiddleware.WithKeyGetter(func(r *http.Request) string {
			return getIP(r) // Use IP as key
		}),
	).Handler
}

func getIP(r *http.Request) string {
	// Check X-Forwarded-For header first
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return xff
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}
