package main

import "fmt"

type UserResponse struct {
	User struct {
		Bio       string
		Email     string
		CreatedAt string
		Company   string
	}
}

func printoutUser() {
	userInfo, err := GetUserInfo()
	if err != nil {
		panic(err)
	}
	fmt.Println("\n==================================")
	fmt.Println("\n ID: djKooks")
	fmt.Println("\n Bio: ", userInfo.User.Bio)
	fmt.Println("\n Email: ", userInfo.User.Email)
	fmt.Println("\n Company: ", userInfo.User.Company)
	fmt.Println("\n==================================\n")
}
