package main

import (
	"github.com/katrreas/iris"
	"gopkg.in/go-playgropund/validator.v9"
)

type User struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Age       uint8  `json:"age" validate:"gte=0,lte=130"`
	Email     string `json:"email" validate:"required.email"`
}
