package cache

import (
	"context"
	"fmt"
	"time"
)

func KeyUpcomingIPOs() string {
	return "ipos:upcoming"
}

func KeyActiveIPOs() string {
	return "ipos:active"
}

func KeyIPO(id string) string {
	return fmt.Sprintf("ipo:%s", id)
}

func KeyIPOSubscription(id string) string {
	return fmt.Sprintf("ipo:%s:subscription", id)
}

func KeyIPOGMP(id string) string {
	return fmt.Sprintf("ipo:%s:gmp", id)
}

// CacheClient defines the cache-aside interface for Redis
type CacheClient interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Delete(ctx context.Context, keys ...string) error
}
