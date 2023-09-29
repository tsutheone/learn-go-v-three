// Package main: General overview of interfaces mocking DB and their unit test.
package main

import (
	"fmt"
	"log"
)

type User struct {
	ID   int
	Name string
}

// Temporary Mock service to act as a DB.
// We could also use a real DB service alongside.
type MockDatastore struct {
	Users map[int]User
}

func (md MockDatastore) GetUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("User ID: %v not found", id)
	} else {
		return user, nil
	}
}

func (md MockDatastore) SaveUser(u User) error {
	_, ok := md.Users[u.ID]
	if ok {
		return fmt.Errorf("User ID: %v already exists", u.ID)
	} else {
		md.Users[u.ID] = u
		return nil
	}
}

type Datastore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

// True service which through the interface is able to handle Mock or DB types.
type Service struct {
	ds Datastore
}

func (s *Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

func (s *Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

func main() {
	mdb := MockDatastore{
		Users: make(map[int]User),
	}
	srvc := Service{
		ds: mdb,
	}
	u1 := User{
		ID:   1,
		Name: "James",
	}

	err := srvc.SaveUser(u1)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	ru, err := srvc.GetUser(u1.ID)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	fmt.Println(ru.ID, ru.Name)
}
