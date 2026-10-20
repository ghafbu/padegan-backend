package methods

import (
	"io"
	"log"
)

func LogActiveStatus(paramStatus bool) {
	if !paramStatus {
		log.SetOutput(io.Discard)
	}
}
