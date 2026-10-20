package ghafpkg

import (
	"math/rand"
	"time"
)

func GHRandomBool() bool {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	return r.Intn(2) == 1
}
