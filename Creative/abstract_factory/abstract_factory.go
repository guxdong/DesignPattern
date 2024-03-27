package creative_af

import "fmt"

// 产品接口
type AbstractApple interface {
	ShowApple()
}

type AbstractBanana interface {
	ShowBanana()
}

type AbstractPear interface {
	ShowPear()
}

// 工厂接口
type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

/* 具体产品类 */
// 中国产品族
type ChinaApple struct{}

func (a *ChinaApple) ShowApple() {
	fmt.Println("this is China apple")
}

type ChinaBanana struct{}

func (b *ChinaBanana) ShowBanana() {
	fmt.Println("this is China banana")
}

type ChinaPear struct{}

func (p *ChinaPear) ShowPear() {
	fmt.Println("this is China pear")
}

// 美国产品族
type USApple struct{}

func (u *USApple) ShowApple() {
	fmt.Println("this is US apple")
}

type USBanana struct{}

func (u *USBanana) ShowBanana() {
	fmt.Println("this is US banana")
}

type USPear struct{}

func (u *USPear) ShowPear() {
	fmt.Println("this is US pear")
}

/* 具体工厂类 */
type ChinaFactory struct{}

func (c *ChinaFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(ChinaApple)
	return apple
}

func (c *ChinaFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(ChinaBanana)
	return banana
}

func (c *ChinaFactory) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(ChinaPear)
	return pear
}

type USFactory struct{}

func (u *USFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(USApple)
	return apple
}

func (u *USFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(USBanana)
	return banana
}

func (u *USFactory) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(USPear)
	return pear
}

// 业务逻辑
func Logics() {
	var absFactory AbstractFactory
	var apple AbstractApple
	var pear AbstractPear

	// 中国工厂
	absFactory = new(ChinaFactory)

	// 中国工厂生产产品
	apple = absFactory.CreateApple()
	apple.ShowApple()

	pear = absFactory.CreatePear()
	pear.ShowPear()

	// 美国工厂
	absFactory = new(USFactory)

	// 美国工厂生产产品
	apple = absFactory.CreateApple()
	apple.ShowApple()

	pear = absFactory.CreatePear()
	pear.ShowPear()
}
