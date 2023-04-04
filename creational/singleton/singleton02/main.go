package main

import "sync"

// 单例类
type singleton struct {
}

// 声明私有变量
var instance *singleton

// 声明锁对象
var mutex sync.Mutex

// 加锁保证协程安全
func GetInstance() *singleton {
	mutex.Lock()
	defer mutex.Unlock()
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()

	println(s1 == s2)
}
