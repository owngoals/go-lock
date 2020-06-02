package golock

import (
	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
)

// LockedFunc 获得到锁后执行的
type LockedFunc func() error

// NewRedsyncLock 创建 redsync 锁
func NewRedsyncLock(p *redis.Pool) *redsync.Redsync {
	return redsync.New([]redsync.Pool{p})
}

func RedsyncLockMutex(r *redsync.Redsync, name string, f LockedFunc, options ...redsync.Option) error {
	l := r.NewMutex(name, options...)
	if err := l.Lock(); err != nil {
		return err
	}
	defer l.Unlock()
	return f()
}
