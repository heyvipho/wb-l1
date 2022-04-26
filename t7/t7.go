package main

import (
	"fmt"
	"sync"
	"time"
)

type SyncMap struct {
	rwm sync.RWMutex
	m   map[interface{}]interface{}
}

func CreateSyncMap() SyncMap {
	return SyncMap{
		m: make(map[interface{}]interface{}),
	}
}

func (v *SyncMap) Set(key interface{}, value interface{}) {
	v.rwm.Lock()
	defer v.rwm.Unlock()

	v.m[key] = value
}

func (v *SyncMap) Get(key interface{}) interface{} {
	v.rwm.RLock()
	defer v.rwm.RUnlock()

	return v.m[key]
}

func main() {
	m := CreateSyncMap()

	go func() {
		m.Set(1, 11)
	}()
	go func() {
		m.Set(2, 22)
	}()
	go func() {
		m.Set(3, 33)
	}()
	go func() {
		m.Set(4, 44)
	}()

	time.Sleep(time.Second)

	fmt.Println(m.Get(1))
	fmt.Println(m.Get(2))
	fmt.Println(m.Get(3))
	fmt.Println(m.Get(4))
}
