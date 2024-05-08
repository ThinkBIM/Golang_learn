package main

import "sync"

var rwMutex sync.RWMutex

var mutex sync.Mutex

func RwMutex() {
	rwMutex.Lock()

	// 加读锁
	rwMutex.RLock()

	defer rwMutex.RUnlock()

	defer rwMutex.Unlock()
}
