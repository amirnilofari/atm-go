package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to the ATM bank")
	fmt.Println("Insert the card into the ATM and enter your four-digit password:")

	var password int
	fmt.Scan(&password)
	if len(strconv.Itoa(password)) != 4 {
		fmt.Println("Something wrong.Please try again!")
		os.Exit(0)
	}
	fmt.Println("Your password is: ", password)

	fmt.Println("Now you can select the desired number and perform the desired operation")
	fmt.Println("[1] - Your balance account")
	fmt.Println("[2] - Withdraw from your account")
	fmt.Println("[3] - Deposit to your account")
	fmt.Println("[4] - Transfer from your account")

	var operation int
	fmt.Scan(&operation)
	if operation == 1 {
		fmt.Println("Your balance account is: ", "1500$")
	}
	if operation == 2 {
		fmt.Println("Enter the withdrawal amount:")
		var withdrawal int
		fmt.Scan(&withdrawal)
		fmt.Println("The amount of " + strconv.Itoa(withdrawal) + "$ was deducted from your account")
	}
	if operation == 3 {
		fmt.Println("Deposit")
	}
	if operation == 4 {
		fmt.Println("Transfer")
	}
	// operations := []string{"Balance", "Withdraw", "Deposit", "Transfer"}
	// b := operations[1]
	// fmt.Printf("%T", b)

	// if operations[1] == "Withdraw" {
	// 	fmt.Println("Withdraw")
	// }
}
