package locker

import (
	"sync"
)

type Locker struct {
	mu sync.Map // map[int64]*sync.Mutex
}

func NewLocker() *Locker {
	return &Locker{}
}

func (l *Locker) getMutex(userID int64) *sync.Mutex {
	mutexInterface, ok := l.mu.Load(userID)
	if ok {
		return mutexInterface.(*sync.Mutex)
	}

	mutex := &sync.Mutex{}
	l.mu.Store(userID, mutex)
	return mutex
}

func (l *Locker) Lock(userID int64) {
	l.getMutex(userID).Lock()
}

func (l *Locker) Unlock(userID int64) {
	l.getMutex(userID).Unlock()
}
