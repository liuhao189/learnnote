package main

import (
	"fmt"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
)

type Address struct {
	Street string `json:"street" validate:"required"`
	City   string `json:"city" validate:"required"`
	Plant  string `json:"plant" validate:"required"`
	Phone  string `json:"phone" validate:"required"`
}

type User struct {
	FirstName      string     `json:"fname"`
	LastName       string     `json:"lname"`
	Age            uint8      `json:"age" validate:"gte=0,lte=130"`
	Email          string     `json:"email" validate:"required.email"`
	FavouriteColor string     `json:"favColor" validate:"hexcolor|rgb|rgba"`
	Addresses      []*Address `json:"address" validate: "required,dive,required"`
}

func UserStructLevelValidation(s1 validator.StructLevel) {
	user := s1.Current().Interface().(User)
	fmt.Println(user.FirstName)
}

var validate *validator.Validate

func main() {
	validate = validator.New()

	validate.RegisterStructValidation(UserStructLevelValidation, User{})

	app := iris.New()

	app.Post("/user", func(ctx iris.Context) {
		var user User
		if err := ctx.ReadJSON(&user); err != nil {
			ctx.WriteString(err.Error())
			return
		}
		err := validate.Struct(user)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.WriteString(err.Error())
				return
			}

			ctx.StatusCode(iris.StatusBadRequest)
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println()
				fmt.Println(err.Namespace())
				fmt.Println(err.Field())
				fmt.Println(err.StructNamespace())
				fmt.Println(err.StructField())
				fmt.Println(err.Tag())
				fmt.Println(err.ActualTag())
				fmt.Println(err.Kind())
				fmt.Println(err.Type())
				fmt.Println(err.Value())
				fmt.Println(err.Param())
				fmt.Println()
			}
		}
	})

	app.Run(iris.Addr(":8080"))
}
