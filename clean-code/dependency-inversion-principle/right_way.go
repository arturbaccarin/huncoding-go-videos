package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type UserRepository interface {
	InsertUser(userID string, name string) *Usr
	GetUserByID(userID string) *Usr
}

type mysqlRepository struct{}

type userServices struct {
	repo UserRepository
}

var (
	inMemoUsrs map[string]Usr
)

func init() {
	inMemoUsrs = make(map[string]Usr)
}

type Usr struct {
	id   string
	name string
}

func main() {
	repo := &mysqlRepository{}
	serv := &userServices{repo: repo}

	name := "Name"
	userInserted := serv.InsertUserServices(name)
	fmt.Printf("User inserted: %#v \n", userInserted)

	userReturned := serv.GetUserByIDServices(userInserted.id)
	fmt.Printf("User returned from get: %#v \n", userReturned)
}

// Baixo nível
func (mysqlRepository) InsertUser(userID string, name string) *Usr {
	user := &Usr{
		id:   userID,
		name: name,
	}

	inMemoUsrs[userID] = *user
	return user
}

// Baixo nível
func (mysqlRepository) GetUserByID(userID string) *Usr {
	user, ok := inMemoUsrs[userID]
	if !ok {
		return nil
	}

	return &user
}

// Alto nível
func (u *userServices) GetUserByIDServices(userID string) *Usr {
	return u.repo.GetUserByID(userID)
}

// Alto nível
func (u *userServices) InsertUserServices(name string) *Usr {
	userID := strconv.Itoa(rand.Intn(100000))

	return u.repo.InsertUser(userID, name)
}
