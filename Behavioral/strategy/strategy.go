package strategy

import "fmt"

// WeaponStrategy 武器策略(抽象的策略)
type WeaponStrategy interface {
	UseWeapone()
}

/* 具体的策略 */
type AK47 struct{}

func (a *AK47) UseWeapone() {
	fmt.Println("使用 AK47 战斗")
}

type Knife struct{}

func (k *Knife) UseWeapone() {
	fmt.Println("使用 匕首 战斗")
}

// Hero 环境类
type Hero struct {
	strategy WeaponStrategy
}

func (h *Hero) SetWeaponStrategy(strategy WeaponStrategy) {
	h.strategy = strategy
}

func (h *Hero) Fight() {
	h.strategy.UseWeapone()
}

// Logic 业务逻辑
func Logic() {
	hero := Hero{}
	hero.SetWeaponStrategy(new(AK47))
	hero.Fight()

	hero.SetWeaponStrategy(new(Knife))
	hero.Fight()
}
