package iterator

import "testing"

// 迭代器接口
type Iterator interface {
	HasNext() bool
	Next() *User
}

// 集合接口
type Collection interface {
	CreateIterator() Iterator
}

type User struct {
	Name string
	Age  int
}

// 具体迭代器
type UserIterator struct {
	Index int
	Users []*User
}

func (u *UserIterator) HasNext() bool {
	return u.Index < len(u.Users)
}

func (u *UserIterator) Next() *User {
	if u.HasNext() {
		user := u.Users[u.Index]
		u.Index++
		return user
	}
	return nil
}

// 具体集合
type UserCollection struct {
	Users []*User
}

// 创建迭代器
func (u *UserCollection) CreateIterator() Iterator {
	return &UserIterator{
		Index: 0,
		Users: u.Users,
	}
}

func TestIterator(t *testing.T) {
	user1 := &User{
		Name: "user1",
		Age:  10,
	}
	user2 := &User{
		Name: "user2",
		Age:  20,
	}

	userCollection := &UserCollection{
		Users: []*User{user1, user2},
	}

	iterator := userCollection.CreateIterator()

	for iterator.HasNext() {
		user := iterator.Next()
		t.Logf("user: %v", user)
	}
}
