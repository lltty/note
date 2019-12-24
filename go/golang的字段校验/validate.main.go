package main

import (
	"fmt"

	gvalidator "github.com/go-playground/validator"
)

type PRegisterReq struct {
	Username   string `validate:"gt=0,lt=8"`
	Password   string `validate:"gt=0"`
	RePassword string `validate:"eqfield=Password"`
	Email      string `validate:"email"`
}

var validator *gvalidator.Validate

func init() {
	validator = gvalidator.New()
}

func rpegister(user PRegisterReq) {

}

func validate(req PRegisterReq) error {
	err := validator.Struct(req)
	if err != nil {
		fmt.Printf("验证失败:%v", err)
	}

	rpegister(req)
	return nil
}

func main() {
	t := PRegisterReq{
		Username:   "toby12345",
		Password:   "123",
		RePassword: "123",
		Email:      "1@123.com",
	}

	validate(t)
}
