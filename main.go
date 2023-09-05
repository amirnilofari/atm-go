package main

import (
	"fmt"
	"os"
	"strconv"
)

var balance int

func main() {
	balance = 2500
	fmt.Println("Welcome to the ATM bank")
	fmt.Println("Insert the card into the ATM and enter your four-digit password:")

	var password int
	CheckPassword(password)

	var operation int
	fmt.Scan(&operation)

	if operation == 1 {
		Balance()
	}
	if operation == 2 {
		var amount int
		Withdrawal(amount)
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

func CheckPassword(pass int) {
	fmt.Scan(&pass)
	if len(strconv.Itoa(pass)) != 4 {
		fmt.Println("Something wrong.Please try again!")
		os.Exit(0)
	}

	fmt.Println("Now you can select the desired number and perform the desired operation")
	fmt.Println("[1] - Your balance account")
	fmt.Println("[2] - Withdraw from your account")
	fmt.Println("[3] - Deposit to your account")
	fmt.Println("[4] - Transfer from your account")
}

func Balance() {
	fmt.Println("Your account balance is $" + strconv.Itoa(balance))
}

func Withdrawal(amount int) {
	fmt.Println("Enter the withdrawal amount:")
	fmt.Scan(&amount)

	if amount > balance {
		fmt.Println("The requested amount exceeds your account balance!")
	} else {
		fmt.Println("The amount of $" + strconv.Itoa(amount) + " was deducted from your account")
		fmt.Println("Your account balance is $" + strconv.Itoa(balance-amount))
	}

}
