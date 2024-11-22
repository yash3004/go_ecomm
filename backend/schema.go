package main

type User struct {
	UserID         string
	Username       string
	Password       string
	Email          string
	CartID         string
	TransanctionId string
	PhoneNo        string
}

type ReadUser struct {
	UserID   string
	Username string
	Email    string
	PhoneNo  string
}

type AuthUser struct {
	UserID   *string
	Username *string
	Password string
}
