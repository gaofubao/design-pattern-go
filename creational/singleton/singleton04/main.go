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

var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
		fmt.Println("创建单个实例")
	})
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()

	fmt.Println(s1 == s2)
}
