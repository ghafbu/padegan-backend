package myredis

import (
	"context"
	"errors"
	"time"

	redispkg "github.com/ghafbu/padegan-backend/pkg/cachepkg/redis"
	"github.com/redis/go-redis/v9"
)

type MyRedis struct {
	ctx     context.Context
	timeout time.Duration
	conn    *redispkg.Redis
}

func Init(client *redispkg.Redis, timeout time.Duration) (*MyRedis, error) {
	return &MyRedis{
		ctx:     context.Background(),
		timeout: timeout,
		conn:    client,
	}, nil
}

func (m *MyRedis) Ping() error {
	return m.conn.Client().Ping(m.ctx).Err()
}

// GetWithContext gets the value for the given key with a context.
func (m *MyRedis) GetWithContext(ctx context.Context, key string) ([]byte, error) {
	//return m.conn.Client().Get(ctx, key).Bytes()

	ctx, cancel := m.withTimeout()
	defer cancel()

	value, err := m.conn.Client().Get(ctx, key).Bytes()

	if errors.Is(err, redis.Nil) {
		return nil, nil
	}

	return value, err
}

// Get gets the value for the given key.
func (m *MyRedis) Get(key string) ([]byte, error) {
	ctx, cancel := m.withTimeout()
	defer cancel()

	return m.GetWithContext(ctx, key)
}

// SetWithContext stores the given value for the given key
// with an expiration value.
func (m *MyRedis) SetWithContext(
	ctx context.Context,
	key string,
	val []byte,
	exp time.Duration,
) error {
	return m.conn.Client().Set(ctx, key, val, exp).Err()
}

// Set stores the given value for the given key
// along with an expiration value.
func (m *MyRedis) Set(
	key string,
	val []byte,
	exp time.Duration,
) error {
	ctx, cancel := m.withTimeout()
	defer cancel()

	return m.SetWithContext(ctx, key, val, exp)
}

// DeleteWithContext deletes the value for the given key with a context.
func (m *MyRedis) DeleteWithContext(
	ctx context.Context,
	key string,
) error {
	return m.conn.Client().Del(ctx, key).Err()
}

// Delete deletes the value for the given key.
func (m *MyRedis) Delete(key string) error {
	ctx, cancel := m.withTimeout()
	defer cancel()

	return m.DeleteWithContext(ctx, key)
}

// ResetWithContext resets the storage and deletes all keys with a context.
func (m *MyRedis) ResetWithContext(ctx context.Context) error {
	return m.conn.Client().FlushDB(ctx).Err()
}

// Reset resets the storage and deletes all keys.
func (m *MyRedis) Reset() error {
	ctx, cancel := m.withTimeout()
	defer cancel()

	return m.ResetWithContext(ctx)
}

// Close closes the storage.
func (m *MyRedis) Close() error {
	return m.conn.Client().Close()
}

func (m *MyRedis) withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(m.ctx, m.timeout)
}
