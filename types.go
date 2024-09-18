package main

import (
	"math/rand/v2"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.IntN(10000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.IntN(10000000)),
		CreatedAt: time.Now().UTC(),
	}
}
