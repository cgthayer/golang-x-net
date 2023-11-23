//go:build smart_rwlock

package http2

import (
	"log"
	"sync"
)

type SmartLock struct {
	mu sync.RWMutex
}

func (sl *SmartLock) Lock(name string) {
	log.Printf("xxx %s write lock", name)
	sl.mu.Lock()
}

func (sl *SmartLock) Unlock() {
	sl.mu.Unlock()
}

func (sl *SmartLock) RLock(name string) {
	log.Printf("xxx %s read lock", name)
	sl.mu.RLock()
}

func (sl *SmartLock) RUnlock() {
	sl.mu.RUnlock()
}
