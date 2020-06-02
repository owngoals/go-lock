package golock

import (
	"github.com/go-redsync/redsync"
	goredis "github.com/owngoals/go-redis"
	"testing"
	"time"
)

func TestRedsyncLockMutex(t *testing.T) {
	p := goredis.CreatePool("127.0.0.1", 6379, 1, "")
	r := NewRedsyncLock(p)
	start := time.Now()
	userId := 1
	if err := RedsyncLockMutex(r, "golock_test", func() error {
		t.Logf("start: %v -- end: %v", start, time.Now())
		t.Log("userId =", userId)
		return nil
	}, redsync.SetExpiry(10*time.Second), redsync.SetTries(3), redsync.SetRetryDelay(3*time.Second)); err != nil {
		t.FailNow()
	}

}
