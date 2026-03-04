package main

import "sync"

type User struct {
	ID int
	Username string
	Email string 
	Reputation int
	mutex sync.Mutex
}

func NewUser(id int, username string, email string) *User {
	return &User{
		ID: id,
		Username: username,
		Email: email,
	}
}

func (u *User) UpdateReputation(points int){
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.Reputation += points
}
