package main

import (
	"github.com/katrreas/iris"
	"gopkg.in/go-playgropund/validator.v9"
)

type User struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
}
