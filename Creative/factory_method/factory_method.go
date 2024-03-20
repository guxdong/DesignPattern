package creative_fm

import "fmt"

// 抽象接口
type Fruit interface {
	Show()
}

type AbstractFactory interface {
	CrearteFruit() Fruit
}

// 基础类模块
type Apple struct{}

func (a *Apple) Show() {
	fmt.Println("this is apple")
}

type Banana struct{}

func (b *Banana) Show() {
	fmt.Println("this is Banana")
}

type Pear struct{}

func (p *Pear) Show() {
	fmt.Println("this is Pear")
}

// 工厂模块
type AppleFactory struct{}

func (a *AppleFactory) CrearteFruit() Fruit {
	fruit := new(Apple)
	return fruit
}

type BananaFactory struct{}

func (b *BananaFactory) CrearteFruit() Fruit {
	fruit := new(Banana)
	return fruit
}

type PearFactory struct{}

func (p *PearFactory) CrearteFruit() Fruit {
	fruit := new(Pear)
	return fruit
}

// 业务逻辑
func LogicsFM() {
	var absFactory AbstractFactory
	absFactory = new(AppleFactory)
	var fruit Fruit
	fruit = absFactory.CrearteFruit()
	fruit.Show()
	absFactory = new(PearFactory)
	fruit = absFactory.CrearteFruit()
	fruit.Show()
}
