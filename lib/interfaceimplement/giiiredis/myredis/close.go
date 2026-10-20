package myredis

import (
	"fmt"
)

func (r *MyRedis) Close() error {
	if thisErr := r.conn.Client().Close(); thisErr != nil {
		return fmt.Errorf("failed to close Redis conn connection: %w", thisErr)
	}

	return nil
}
