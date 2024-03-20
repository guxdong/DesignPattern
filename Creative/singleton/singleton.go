package creative_s

import "sync"

type singleton struct{}

var instance *singleton
var iOnce sync.Once
var mutex sync.Mutex

// init init函数在main函数之前执行，并且是自动执行，每个包中可以有多个init函数，每个源文件中也可以有多个init函数。
func init() {
	if instance == nil {
		instance = &singleton{}
	}
}

// GetHungryInstance 饿汉模式
func GetHungryInstance() *singleton {
	return instance
}

// GetLazyInstance 懒汉模式
func GetLazyInstance() *singleton {
	mutex.Lock()
	if instance == nil {
		instance = &singleton{}
	}
	mutex.Unlock()
	return instance
}

// GetLazyInstanceDoubleCheck 双重检查单例模式
func GetLazyInstanceDoubleCheck() *singleton {
	if instance == nil {
		mutex.Lock()
		if instance == nil {
			instance = &singleton{}
		}
		mutex.Unlock()
	}
	return instance
}

func GetInstance() *singleton {
	iOnce.Do(func() {
		instance = &singleton{}
	})
	return instance
}
