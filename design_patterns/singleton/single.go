package singlepattern

//饿汉式实现单例模式
var single *singleton

type singleton struct {
}

func init() {
	single = new(singleton)
}

func GetInstance() *singleton {
	return single
}
