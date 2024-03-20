package creative

import "fmt"

type Fruit interface {
	Show()
}

type Apple struct{}

func (a *Apple) Show() {
	fmt.Println("this is apple")
}

type Banana struct{}

func (b *Banana) Show() {
	fmt.Println("this is banana")
}

type Pear struct{}

func (p *Pear) Show() {
	fmt.Println("this is pear")
}

type SimpleFactory struct{}

func (s *SimpleFactory) CreateFruit(name string) Fruit {
	var fruit Fruit
	if name == "apple" {
		fruit = new(Apple)
	} else if name == "banana" {
		fruit = new(Banana)
	} else if name == "pear" {
		fruit = new(Pear)
	}
	return fruit
}

// 业务逻辑
func logicsSF() {
	factory := new(SimpleFactory)

	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()
}
