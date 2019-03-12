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

func configReader() {
	jsonRaw, err := os.Open("user.json")
	if err != nil {
		log.Println(err)
		return
	}

	defer jsonRaw.Close()

	userByte, err := ioutil.ReadAll(jsonRaw)

	if err != nil {
		log.Println(err)
		return
	}

	var user User
	json.Unmarshal(userByte, &user)
	log.Println(user.Token)
}