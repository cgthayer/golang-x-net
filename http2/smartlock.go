//go:build !smart_rwlock

package http2

import (
	"log"
	"sync"
)

type SmartLock struct {
	mu sync.Mutex
}

func (sl *SmartLock) Lock(name string) {
	log.Printf("xxx %s fake write lock", name)
	sl.mu.Lock()
}

func (sl *SmartLock) Unlock() {
	sl.mu.Unlock()
}

func (sl *SmartLock) RLock(name string) {
	log.Printf("xxx %s fake read lock", name)
	sl.mu.Lock()
}

func (sl *SmartLock) RUnlock() {
	sl.mu.Unlock()
}
