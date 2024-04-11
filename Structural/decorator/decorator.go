package decorator

import "fmt"

type Phone interface {
	Show()
}

// 装饰器基础类（该类本应该为interface，但是Golang interface语法不可以有成员属性）
type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {}

type HuaWei struct{}

func (h *HuaWei) Show() {
	fmt.Println("秀出了华为手机")
}

type Oppo struct{}

func (o *Oppo) Show() {
	fmt.Println("秀出了OPPO手机")
}

type PhoneCaseDecorator struct {
	Decorator
}

func NewPhoneCaseDecorator(phone Phone) *PhoneCaseDecorator {
	return &PhoneCaseDecorator{Decorator{phone: phone}}
}

func (p *PhoneCaseDecorator) Show() {
	p.phone.Show()
	// 装饰额外的方法
	fmt.Println("装了手机壳，增强了手机的抗摔能力")
}

type ScreenFilmDecorator struct {
	Decorator
}

func NewSreenFilmDecorator(phone Phone) *ScreenFilmDecorator {
	return &ScreenFilmDecorator{Decorator{phone: phone}}
}

func (s *ScreenFilmDecorator) Show() {
	s.phone.Show()
	// 装饰额外的方法
	fmt.Println("贴了屏幕膜，增强了手机屏幕的防刮能力")
}

// 业务逻辑
func Logic() {
	var phone Phone
	phone = new(HuaWei)
	phone.Show()
	fmt.Println("--------")

	// 给手机贴膜
	var screenFilmPhone Phone
	screenFilmPhone = NewSreenFilmDecorator(phone)
	screenFilmPhone.Show()
	fmt.Println("--------")

	// 给手机装手机壳
	var phoneCasePhone Phone
	phoneCasePhone = NewPhoneCaseDecorator(phone)
	phoneCasePhone.Show()
	fmt.Println("--------")

	// 给贴了膜的手机装手机壳
	phoneCasePhone = NewPhoneCaseDecorator(screenFilmPhone)
	phoneCasePhone.Show()
}
