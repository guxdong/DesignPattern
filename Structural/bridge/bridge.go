package bridge

import "fmt"

// Traffic 交通方式接口
type Traffic interface {
	Transport() string
}

/* 交通方式的实现类 */
type Train struct{}

func (t *Train) Transport() string {
	return "train"
}

type Plane struct{}

func (p *Plane) Transport() string {
	return "plane"
}

// Location 活动接口
type Location interface {
	Name() string
	Activity() string
}

type namedLocation struct {
	name string
}

// Name 获取地点名称
func (n *namedLocation) Name() string {
	return n.name
}

/* 活动的实现类 */
type Mountain struct {
	namedLocation
}

func NewMountain(name string) *Mountain {
	return &Mountain{
		namedLocation{
			name: name,
		},
	}
}

func (m *Mountain) Activity() string {
	return "climbing"
}

type Park struct {
	namedLocation
}

func NewPark(name string) *Park {
	return &Park{
		namedLocation{
			name: name,
		},
	}
}

func (p *Park) Activity() string {
	return "walking"
}

type SeaSide struct {
	namedLocation
}

func NewSeaSide(name string) *SeaSide {
	return &SeaSide{
		namedLocation{
			name: name,
		},
	}
}

func (s *SeaSide) Activity() string {
	return "surfing"
}

// Experience 经历描述接口
type Experience interface {
	Describe()
}

/* 经历描述的实现类 */
type Travel struct {
	traffic  Traffic
	location Location
}

func NewTravel(traffic Traffic, location Location) *Travel {
	return &Travel{
		traffic:  traffic,
		location: location,
	}
}

func (t *Travel) Describe() {
	fmt.Printf("乘坐 %s 去 %s %s.\n", t.traffic.Transport(), t.location.Name(), t.location.Activity())
}

type Adventure struct {
	traffic  Traffic
	location Location
}

func NewAdventure(traffic Traffic, location Location) *Adventure {
	return &Adventure{
		traffic:  traffic,
		location: location,
	}
}

func (a *Adventure) Describe() {
	fmt.Printf("乘坐 %s 去 %s %s.\n", a.traffic.Transport(), a.location.Name(), a.location.Activity())
}

// 业务逻辑
func Logic() {
	var experience Experience
	experience = NewTravel(new(Plane), NewSeaSide("三亚"))
	experience.Describe()
	experience = NewAdventure(new(Train), NewMountain("黄山"))
	experience.Describe()
}
