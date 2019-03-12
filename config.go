package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type User struct {
	Token string `json:"token"`
}

func configReader() (User, error) {
	jsonRaw, err := os.Open("user.json")
	if err != nil {
		log.Println(err)
		return User{}, err
	}

	defer jsonRaw.Close()

	userByte, err := ioutil.ReadAll(jsonRaw)

	if err != nil {
		log.Println(err)
		return User{}, err
	}

	var user User
	json.Unmarshal(userByte, &user)

	return user, nil
}
