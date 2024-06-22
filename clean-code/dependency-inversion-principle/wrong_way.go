package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

var (
	inMemoUsers map[string]User
)

func init() {
	inMemoUsers = make(map[string]User)
}

type User struct {
	id   string
	name string
}

func mainWrong() {
	name := "Name"
	userInserted := InsertUserServices(name)
	fmt.Printf("User inserted: %#v \n", userInserted)

	userReturned := GetUserByIDServices(userInserted.id)
	fmt.Printf("User returned from get: %#v \n", userReturned)
}

// Baixo nível
func InsertUserIntoMySQLRepo(userID string, name string) *User {
	user := &User{
		id:   userID,
		name: name,
	}

	inMemoUsers[userID] = *user
	return user
}

// Baixo nível
func GetUserByIDFromMySQLRepo(userID string) *User {
	user, ok := inMemoUsers[userID]
	if !ok {
		return nil
	}

	return &user
}

// Alto nível
func GetUserByIDServices(userID string) *User {
	return GetUserByIDFromMySQLRepo(userID)
}

// Alto nível
func InsertUserServices(name string) *User {
	userID := strconv.Itoa(rand.Intn(100000))

	// Se mudar para o mongo, vai ter que mudar aqui dentro do alto nível
	// InsertUserIntoMongoRepo
	return InsertUserIntoMySQLRepo(userID, name)
}
