package adapter

import "fmt"

// TypeCPlug Type-C接口
type TypeCPlug interface {
	ConnectTypeC() string
}

type Oppo struct {
	name string
}

func NewOppo(name string) *Oppo {
	return &Oppo{
		name: name,
	}
}

// ConnectTypeC OPPO使用的是Type-C接口
func (o *Oppo) ConnectTypeC() string {
	return fmt.Sprintf("%v 使用Type-C接口", o.name)
}

// LightingPlug lighting接口
type LightingPlug interface {
	ConnectLighting() string
}

type IPhone struct {
	name string
}

func NewIPhone(name string) *IPhone {
	return &IPhone{
		name: name,
	}
}

// ConnectLighting iPhone使用的是Lighting接口
func (i *IPhone) ConnectLighting() string {
	return fmt.Sprintf("%v 使用Lighting接口", i.name)
}

// CommonPlug 通用的USB接口
type CommonPlug interface {
	ConnectUSB() string
}

// TypeCPlugAdapter Type-C接口适配USB接口
type TypeCPlugAdapter struct {
	typec TypeCPlug
}

func NewTypeCPlugAdapter(typec TypeCPlug) *TypeCPlugAdapter {
	return &TypeCPlugAdapter{
		typec: typec,
	}
}

func (t *TypeCPlugAdapter) ConnectUSB() string {
	return fmt.Sprintf("%v 适配USB ", t.typec.ConnectTypeC())
}

// LightingPlugAdapter lighting接口适配USB
type LightingPlugAdapter struct {
	lighting LightingPlug
}

func NewLightingPlugAdapter(lighting LightingPlug) *LightingPlugAdapter {
	return &LightingPlugAdapter{
		lighting: lighting,
	}
}

func (l *LightingPlugAdapter) ConnectUSB() string {
	return fmt.Sprintf("%v 适配USB", l.lighting.ConnectLighting())
}

type PowerBank struct {
	brand string
}

func NewPowerBank(brand string) *PowerBank {
	return &PowerBank{
		brand: brand,
	}
}

func (p *PowerBank) Charge(plug CommonPlug) {
	fmt.Printf("使用 %v 充电宝给 %v 进行充电.\n", p.brand, plug.ConnectUSB())
}

// Logic 业务逻辑
func Logic() {
	powerBank := NewPowerBank("小米")

	var commonPlug CommonPlug

	// 将iPhone的Lighting接口转换成通用的USB接口
	commonPlug = NewLightingPlugAdapter(NewIPhone(("iPhone 15 Pro Max")))
	// 通过USB接口进行充电
	powerBank.Charge(commonPlug)

	// 将OPPO的Type-C接口转换成通用的USB接口
	commonPlug = NewTypeCPlugAdapter(NewOppo("OPPO 10000"))
	// 通过USB接口进行充电
	powerBank.Charge(commonPlug)
}
