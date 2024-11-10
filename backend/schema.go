package main

import "hash"

type User struct {
	UserID   string
	Username string
	Password hash.Hash64
	Email string 
	CartID string 
	TransanctionId string 
}