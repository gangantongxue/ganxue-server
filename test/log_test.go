package test

import (
	"ganxue-server/utils/log"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	log.Debug("test")
	log.Error("errTest")
	time.Sleep(time.Second * 5)
}
