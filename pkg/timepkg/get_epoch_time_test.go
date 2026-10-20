package timepkg

import (
	"log"
	"testing"
	"time"
)

func TestGetEpochTime(t *testing.T) {
	log.Println(GetEpochTime())

	time.Sleep(1 * time.Second)

	log.Println(GetEpochTime())

	log.Println(time.Now())
	log.Println(time.Now().UnixNano())
}
