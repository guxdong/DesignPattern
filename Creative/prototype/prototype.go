package prototype

import (
	"fmt"
)

// 原型接口
type Prototype interface {
	GetName() string
	GetAge() int
	Clone() Prototype
}

// 具体原型结构体
type ConcretePrototype struct {
	Name string
	Age  int
}

func (p *ConcretePrototype) GetName() string {
	return p.Name
}

func (p *ConcretePrototype) GetAge() int {
	return p.Age
}

// 实现Clone方法
func (p *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{
		Name: p.Name,
		Age:  p.Age,
	}
}

func Logics() {
	// 创建原型对象
	prototype := &ConcretePrototype{
		Name: "Prototype",
		Age:  20,
	}

	// 克隆原型对象
	clone := prototype.Clone()

	// 打印原型对象和克隆对象的名称
	fmt.Println("原型对象的名称:", prototype.Name)
	fmt.Println("克隆对象的名称:", clone.GetName())
	// 打印原型对象和克隆对象的Age
	fmt.Println("原型对象的Age:", prototype.Age)
	fmt.Println("克隆对象的Age:", clone.GetAge())
}
