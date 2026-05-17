package main

import (
	"fmt"

	"example.com/structs/user"
)

type str string

func (text str) log() {
	fmt.Println(text)
}

func main() {
	var name str
	name = "Harshal"

	name.log()

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birth date (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("harshal@example.com", "12345")

	admin.DisplayUserDetails()
	admin.DisplayUserDetails()

	appUser.DisplayUserDetails()
	appUser.ClearUserName()
	appUser.DisplayUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
