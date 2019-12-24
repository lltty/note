package main

import (
	"errors"
)

type RegisterReq struct {
	Username       string `json:"username"`
	PasswordNew    string `json:"password_new"`
	PasswordRepeat string `json:"password_repeat"`
	Email          string `json:email`
}

func emailValid(email string) bool {

	return true
}

func register(user RegisterReq) {

}

func Register(req RegisterReq) error {
	if len(req.Username) > 0 {
		if len(req.PasswordNew) > 0 && len(req.PasswordRepeat) > 0 {
			if req.PasswordNew == req.PasswordRepeat {
				if emailValid(req.Email) {
					register(req)
					return nil
				} else {
					return errors.New("invalid email")
				}
			} else {
				return errors.New("password and reinput must be the same")
			}
		} else {
			return errors.New("password and password reinput require")
		}
	} else {
		return errors.New("name cannot be empty")
	}
}

func RegisterOpt(req RegisterReq) error {

	if len(req.Username) > 0 {
		return errors.New("name cannot be empty")
	}

	if len(req.PasswordNew) == 0 || len(req.PasswordRepeat) == 0 {
		return errors.New("password and password reinput require")
	}

	//非友好写法
	/*if emailValid(req.Email) {
		register(req)
		return nil
	} else {
		return errors.New("invalid email")
	}*/

	if !emailValid(req.Email) {
		return errors.New("invalid email")
	}

	register(req)
	return nil
}
