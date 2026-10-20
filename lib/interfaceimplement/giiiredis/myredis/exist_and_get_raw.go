package myredis

import (
	"context"
	"fmt"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
)

func (r *MyRedis) ExistAndGetRaw(
	paramCtx context.Context,
	paramDefaultKey string,
	paramKey string,
) (string, bool, error) {
	var thisKeyForRedis = ghafpkg.BuildKeyPattern(paramDefaultKey, paramKey)

	exist, err := r.conn.Client().Exists(paramCtx, thisKeyForRedis).Result()
	if err != nil {
		return "", false, fmt.Errorf("redis exist error: %w", err) // 500 Internal Server Error
	}

	if exist == 0 {
		return "", false, nil // 404 Not Found
	}

	thisResult, thisErr := r.conn.Client().Get(paramCtx, thisKeyForRedis).Result()
	if thisErr != nil {
		return "", false, fmt.Errorf("redis get error: %w", err) // 500 Internal Server Error
	}

	return thisResult, true, nil // found item
}
