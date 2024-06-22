package interfaces

import "fmt"

type User struct {
	ID   string
	Name string
}

type UserMethods interface {
	GetUserByID(id string)
	InsertUser(name string)
}

func (u *User) GetUserByID(id string)  {}
func (u *User) InsertUser(name string) {}

func test() {
	u := &User{}
	// u := &UserMocks{}

	// user satisfaz userMethods
	var userMethods UserMethods

	userMethods = u
	fmt.Println(userMethods)
}

type UserMock struct{}

func (u *UserMock) GetUserByID(id string)  {}
func (u *UserMock) InsertUser(name string) {}
