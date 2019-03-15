package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type User struct {
	Token string `json:"token"`
	Id    string `json:"id"`
}

func ConfigReader() (User, error) {
	userByte, err := jsonAsBytes()
	if err != nil {
		// TODO: error handler
	}

	var user User
	json.Unmarshal(userByte, &user)

	return user, nil
}

func ConfigWriter(User) error {
	userByte, err := jsonAsBytes()
	if err != nil {
		// TODO: error handler
	}

	var user User
	json.Unmarshal(userByte, &user)

	return nil
}

func jsonAsBytes() ([]byte, error) {
	jsonRaw, err := os.Open("user.json")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer jsonRaw.Close()

	userByte, err := ioutil.ReadAll(jsonRaw)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return userByte, nil
}
