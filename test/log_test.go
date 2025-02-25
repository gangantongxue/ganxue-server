package test

import (
	"ganxue-server/utils/log"
	"testing"
)

func TestLog(t *testing.T) {
	log.Debug("test")
	log.Error("errTest")
}
