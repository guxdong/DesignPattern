package iterator

import "fmt"

type User struct {
	Name string
	Age  int
}

// Iterator 迭代器接口
type Iterator interface {
	// HasNest 是否有下一个
	HasNest() bool
	// GetNext 获取元素
	GetNext() *User
}

// UserIterator 具体的迭代器类
type UserIterator struct {
	index int
	Users []*User
}

func (u *UserIterator) HasNest() bool {
	return u.index < len(u.Users)
}

func (u *UserIterator) GetNext() *User {
	if u.HasNest() {
		user := u.Users[u.index]
		u.index++
		return user
	}
	return nil
}

// Collection 集合接口
type Collection interface {
	// CreateIterator 创建迭代器
	CreateIterator() Iterator
}

// UserCollection 具体的集合类
type UserCollection struct {
	Users []*User
}

func (u *UserCollection) CreateIterator() Iterator {
	return &UserIterator{
		Users: u.Users,
	}
}

// Logic 业务逻辑
func Logic() {
	user1 := User{
		Name: "张三",
		Age:  30,
	}
	user2 := User{
		Name: "李四",
		Age:  20,
	}
	collection := UserCollection{
		Users: []*User{&user1, &user2},
	}
	iterator := collection.CreateIterator()
	for iterator.HasNest() {
		user := iterator.GetNext()
		fmt.Printf("%+v\n", user)
	}
}
