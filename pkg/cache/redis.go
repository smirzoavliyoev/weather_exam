package cache

import (
	"context"
	"time"
	"weather/pkg/config"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type Cacher interface {
	Set(ctx context.Context, key string, value interface{}, d time.Duration) error
	Get(ctx context.Context, key string, wanted interface{}) error
}
type cacher struct {
	redisCache *cache.Cache
}

func GetCacher(cfg config.Configer) Cacher {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server": cfg.GetString("REDIS_SERVER_ADDRESS"),
		},
	})

	cache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return &cacher{
		redisCache: cache,
	}

}

func (r *cacher) Set(ctx context.Context, key string, value interface{}, d time.Duration) error {
	return r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   d,
	})
}

func (r *cacher) Get(ctx context.Context, key string, wanted interface{}) error {
	return r.redisCache.Get(ctx, key, wanted)
}
