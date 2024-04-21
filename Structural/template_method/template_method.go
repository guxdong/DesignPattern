package template_method

import "fmt"

type Actor interface {
	// DressUp 装扮，模板方法
	DressUp()
}

// dressBehavior 定义装扮的一些方法，都是私有方法，通过DressUp模板方法调用
type dressBehavior interface {
	// makeUp 化妆
	makeUp() string
	// wearClothe 穿衣
	wearClothe() string
}

// BaseActor 演员基类
type BaseActor struct {
	roleName      string // 角色名
	dressBehavior        // 装扮行为
}

// DressUp 实现模板方法，定义装扮行为的骨架
func (b *BaseActor) DressUp() {
	str := fmt.Sprintf("扮演 %s 的", b.roleName)
	str += b.makeUp()
	str += b.wearClothe()
	fmt.Println(str)
}

type FemaleActor struct {
	BaseActor // 隐式继承
}

func NewFemaleActor(roleName string) *FemaleActor {
	actor := new(FemaleActor)
	actor.roleName = roleName
	actor.dressBehavior = actor
	return actor
}

func (f *FemaleActor) makeUp() string {
	return "女演员涂口红、画眉毛; "
}

func (f *FemaleActor) wearClothe() string {
	return "穿上连衣裙; "
}

type MaleActor struct {
	BaseActor // 隐式继承
}

func NewMaleActor(roleName string) *MaleActor {
	actor := new(MaleActor)
	actor.roleName = roleName
	actor.dressBehavior = actor
	return actor
}

func (m *MaleActor) makeUp() string {
	return "刮胡子，整理发型; "
}

func (m *MaleActor) wearClothe() string {
	return "穿上西装; "
}

// Logic 业务逻辑
func Logic() {
	femaleActor := NewFemaleActor("妈妈")
	maleActor := NewMaleActor("爸爸")
	femaleActor.DressUp()
	maleActor.DressUp()
}
