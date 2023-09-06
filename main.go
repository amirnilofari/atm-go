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
		var deposit_amount int
		Deposit(deposit_amount)
	}
	if operation == 4 {
		var cardNumber string
		var transferAmount int
		Transfer(cardNumber, transferAmount)
	}

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

func Deposit(amount int) {
	fmt.Println("Enter the amount you want to deposit into your account:")
	fmt.Scan(&amount)

	fmt.Println("The amount of $" + strconv.Itoa(amount) + " added from your account")
	fmt.Println("Your account balance is $" + strconv.Itoa(balance+amount))
}

func Transfer(card_number string, amount int) {
	fmt.Println("Enter the 16-digit card number of the destination for money transfer:")
	fmt.Scan(&card_number)
	if len(card_number) != 16 {
		fmt.Println("The entered card number is not valid.Please try again!")
		os.Exit(0)
	}
	fmt.Println("Now enter the amount of money to transfer:")
	fmt.Scan(&amount)
	if amount > balance {
		fmt.Println("The requested amount exceeds your account balance!")
	} else {
		fmt.Println("The amount of $" + strconv.Itoa(amount) + " was deducted from your account")
		fmt.Println("Your account balance is $" + strconv.Itoa(balance-amount))
	}
}
