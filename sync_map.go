package main

import (
	"fmt"
	"strconv"
	"sync"
)

/**
*@author: 廖理
*@date:2022/8/11
**/

var sysMap = make(map[string]int)

func get(key string) int {
	rwMutex.RLock()
	v := sysMap[key]
	rwMutex.RUnlock()
	return v
}

func set(key string, value int) {
	rwMutex.Lock()
	sysMap[key] = value
	rwMutex.Unlock()
}

func syncSysMap() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//将上面的代码编译后执行，会报出fatal error: concurrent map writes错误。加锁才能解决问题
//我们不能在多个 goroutine 中并发对内置的 map 进行读写操作，否则会存在数据竞争问题。
//像这种场景下就需要为 map 加锁来保证并发的安全性了，Go语言的sync包中提供了一个开箱即用的并发安全版 map——sync.Map。
//开箱即用表示其不用像内置的 map 一样使用 make 函数初始化就能直接使用。
//同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。

//func (m *Map) Store(key, value interface{})	存储key-value数据
//func (m *Map) Load(key interface{}) (value interface{}, ok bool)	查询key对应的value
//func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)	查询或存储key对应的value
//func (m *Map) LoadAndDelete(key interface{}) (value interface{}, loaded bool)	查询并删除key
//func (m *Map) Delete(key interface{})	删除key
//func (m *Map) Range(f func(key, value interface{}) bool)	对map中的每个key-value依次调用f
//下面的代码示例演示了并发读写sync.Map。

// 并发安全的map
var syncMap = sync.Map{}

func syncMapDemo() {
	wg := sync.WaitGroup{}
	// 对m执行20个并发的读写操作
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			syncMap.Store(key, n)         // 存储key-value
			value, _ := syncMap.Load(key) // 根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
