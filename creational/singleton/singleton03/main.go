package main

import (
	"fmt"
	"sync"
)

// 单例类
type singleton struct {
}

// 声明私有变量
var instance *singleton

// 声明锁对象
var mutex sync.Mutex

// 当对象为空时，对对象加锁，当创建好对象后，获取对象时就不用加锁了
func GetInstance() *singleton {
	if instance == nil {
		mutex.Lock()
		if instance == nil {
			instance = new(singleton)
			fmt.Println("创建单个实例")
		}
		mutex.Unlock()
	}
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()

	fmt.Println(s1 == s2)
}
