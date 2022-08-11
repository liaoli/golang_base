package main

import (
	"image"
	"sync"
)

/**
*@author: 廖理
*@date:2022/8/11
**/

//sync.Once
//在某些场景下我们需要确保某些操作即使在高并发的场景下也只会被执行一次，例如只加载一次配置文件等。
//
//Go语言中的sync包中提供了一个针对只执行一次场景的解决方案——sync.Once，sync.Once只有一个Do方法，其签名如下：
//
//func (o *Once) Do(f func())
//注意：如果要执行的函数f需要传递参数就需要搭配闭包来使用。

var icons map[string]image.Image
var loadIconsOnce sync.Once

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

func loadIcon(s string) image.Image {
	return &image.RGBA{}
}

// Icon1 被多个goroutine调用时不是并发安全的
func Icon1(name string) image.Image {
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

//Icon sync.Once可以保证只执行一次
func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons) //只会执行一次
	return icons[name]
}

//单例sync.Once
type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
