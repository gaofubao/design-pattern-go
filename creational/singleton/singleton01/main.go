package main

// 单例类
type singleton struct {
}

// 声明私有变量
var instance *singleton

// 获取单例对象
func GetInstance() *singleton {
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
