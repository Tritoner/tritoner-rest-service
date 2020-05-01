package main

import "time"

type User struct {
	Id        uint64    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Birthday  time.Time `json:"birthday"`
	Bio       string    `json:"bio"`
}

type Users []User
