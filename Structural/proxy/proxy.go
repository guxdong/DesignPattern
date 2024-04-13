package proxy

import "fmt"

type HouseSeller interface {
	SellHouse(address, buyer string)
}

// HouseSeller 房屋出售者
type HouseOwner struct{}

func (h *HouseOwner) SellHouse(address, buyer string) {
	fmt.Printf("最终商讨价格后，与%s签署购买地址为%s的购房协议。\n", buyer, address)
}

// HouseProxy 房产中介
type HouseProxy struct {
	houseSeller HouseSeller
}

func NewHouseProxy() *HouseProxy {
	// 在代理对象中隐藏实际的目标对象，实现访问控制
	houseSeller := &HouseOwner{}
	return &HouseProxy{
		houseSeller: houseSeller,
	}
}

// viewHouse 看房介绍基本情况
func (h *HouseProxy) viewHouse(address string, buyer string) {
	fmt.Printf("带买家%s看位于%s的房屋，并介绍基本情况\n", buyer, address)
}

// preBargain 初步沟通价格
func (h *HouseProxy) preBargain(address string, buyer string) {
	fmt.Println("讨价还价后，初步达成购买意向")
}

// SellHouse 中介卖房，看房->初步谈价->最终和房东签协议
func (h *HouseProxy) SellHouse(address string, buyer string) {
	h.viewHouse(address, buyer)
	h.preBargain(address, buyer)
	h.houseSeller.SellHouse(address, buyer)
}

func Logics() {
	proxy := NewHouseProxy()
	proxy.SellHouse("北京市海淀区中关村大街，2号院1号楼4单元502室", "李四")
}
