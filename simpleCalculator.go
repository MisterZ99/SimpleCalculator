package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
)

var button string

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func addition() {
	var (
		number  string
		number1 float64
		number2 float64
	)
	clear()
	for {
		fmt.Print("\t\t\t--\tSimple Calculator\t--\n\n\n")
		fmt.Print("\t\t\t\t Addition!\n\n")
		fmt.Printf("\t\t\t\t %.2f + %.2f = %.2f\n\n", number1, number2, number1+number2)
		fmt.Println("\n\n\tDigit the digit 1 then confirm with 'enter' to set the number left.\n\tDigit the digit 2 then confirm with 'enter' to set the number right.\n\n")
		fmt.Println("\tOther digit or integer number then confirm with 'enter' to back to menu.\n")
		fmt.Println("\t\t\t\t\tMr. Z!©\n")
		fmt.Scanf("%s", &number)
		_, err := strconv.Atoi(number)
		if err == nil {
			if number[0] == '1' {
				fmt.Print("Number 1:")
				fmt.Scanf("%f", &number1)
			} else if number[0] == '2' {
				fmt.Print("Number 2:")
				fmt.Scanf("%f", &number2)
			} else {
				menu()
			}
		}
		clear()
	}
}

func subtraction() {
	var (
		number  string
		number1 float64
		number2 float64
	)
	clear()
	for {
		fmt.Print("\t\t\t--\tSimple Calculator\t--\n\n\n")
		fmt.Print("\t\t\t\t Subtraction!\n\n")
		fmt.Printf("\t\t\t\t %.2f - %.2f = %.2f\n\n", number1, number2, number1-number2)
		fmt.Println("\n\n\tDigit the digit 1 then confirm with 'enter' to set the number left.\n\tDigit the digit 2 then confirm with 'enter' to set the number right.\n\n")
		fmt.Println("\tOther digit or integer number then confirm with 'enter' to back to menu.\n")
		fmt.Println("\t\t\t\t\tMr. Z!©\n")
		fmt.Scanf("%s", &number)
		_, err := strconv.Atoi(number)
		if err == nil {
			if number[0] == '1' {
				fmt.Print("Number 1:")
				fmt.Scanf("%f", &number1)
			} else if number[0] == '2' {
				fmt.Print("Number 2:")
				fmt.Scanf("%f", &number2)
			} else {
				menu()
			}
		}
		clear()
	}
}

func multiplication() {
	var (
		number  string
		number1 float64
		number2 float64
	)
	clear()
	for {
		fmt.Print("\t\t\t--\tSimple Calculator\t--\n\n\n")
		fmt.Print("\t\t\t\t Multiplication!\n\n")
		fmt.Printf("\t\t\t\t %.2f * %.2f = %.2f\n\n", number1, number2, number1*number2)
		fmt.Println("\n\n\tDigit the digit 1 then confirm with 'enter' to set the number left.\n\tDigit the digit 2 then confirm with 'enter' to set the number right.\n\n")
		fmt.Println("\tOther digit or integer number then confirm with 'enter' to back to menu.\n")
		fmt.Println("\t\t\t\t\tMr. Z!©\n")
		fmt.Scanf("%s", &number)
		_, err := strconv.Atoi(number)
		if err == nil {
			if number[0] == '1' {
				fmt.Print("Number 1:")
				fmt.Scanf("%f", &number1)
			} else if number[0] == '2' {
				fmt.Print("Number 2:")
				fmt.Scanf("%f", &number2)
			} else {
				menu()
			}
		}
		clear()
	}
}

func division() {
	var (
		number  string
		number1 float64
		number2 float64
	)
	clear()
	for {
		fmt.Print("\t\t\t--\tSimple Calculator\t--\n\n\n")
		fmt.Print("\t\t\t\t Division!\n\n")
		fmt.Printf("\t\t\t\t %.2f / %.2f = %.2f\n\n", number1, number2, number1/number2)
		fmt.Println("\n\n\tDigit the digit 1 then confirm with 'enter' to set the number left.\n\tDigit the digit 2 then confirm with 'enter' to set the number right.\n\n")
		fmt.Println("\tOther digit or integer number then confirm with 'enter' to back to menu.\n")
		fmt.Println("\t\t\t\t\tMr. Z!©\n")
		fmt.Scanf("%s", &number)
		_, err := strconv.Atoi(number)
		if err == nil {
			if number[0] == '1' {
				fmt.Print("Number 1:")
				fmt.Scanf("%f", &number1)
			} else if number[0] == '2' {
				fmt.Print("Number 2:")
				fmt.Scanf("%f", &number2)
			} else {
				menu()
			}
		}
		clear()
	}
}

func power() {
	var (
		number  string
		number1 float64
		number2 float64
	)
	clear()
	for {
		fmt.Print("\t\t\t--\tSimple Calculator\t--\n\n\n")
		fmt.Print("\t\t\t\t Power!\n\n")
		fmt.Printf("\t\t\t\t %.2f ^ %.2f = %.2f\n\n", number1, number2, math.Pow(number1, number2))
		fmt.Println("\n\n\tDigit the digit 1 then confirm with 'enter' to set the number left.\n\tDigit the digit 2 then confirm with 'enter' to set the number right.\n\n")
		fmt.Println("\tOther digit or integer number then confirm with 'enter' to back to menu.\n")
		fmt.Println("\t\t\t\t\tMr. Z!©\n")
		fmt.Scanf("%s", &number)
		_, err := strconv.Atoi(number)
		if err == nil {
			if number[0] == '1' {
				fmt.Print("Number 1:")
				fmt.Scanf("%f", &number1)
			} else if number[0] == '2' {
				fmt.Print("Number 2:")
				fmt.Scanf("%f", &number2)
			} else {
				menu()
			}
		}
		clear()
	}
}

func squareRoot() {
	var (
		number  string
		number1 float64
	)
	clear()
	for {
		fmt.Print("\t\t\t--\tSimple Calculator\t--\n\n\n")
		fmt.Print("\t\t\t\t Square Root!\n\n")
		fmt.Printf("\t\t\t\t √%.2f = %.2f\n\n", number1, math.Sqrt(number1))
		fmt.Println("\n\n\tDigit the digit 1 then confirm with 'enter' to set the number.\n\n")
		fmt.Println("\tOther digit or integer number then confirm with 'enter' to back to menu.\n")
		fmt.Println("\t\t\t\t\tMr. Z!©\n")
		fmt.Scanf("%s", &number)
		_, err := strconv.Atoi(number)
		if err == nil {
			if number[0] == '1' {
				fmt.Print("Number:")
				fmt.Scanf("%f", &number1)
			} else {
				menu()
			}
		}
		clear()
	}
}

func saida() {
	clear()
	os.Exit(0)
}

func menu() {
	op := 0
	button = "0"
	clear()
	for {
		fmt.Print("\t\t\t--\tSimple Calculator\t--\n\n\n")
		fmt.Print("\t\t\t\tSelect the operation!\n\n")
		if op == 0 {
			fmt.Println("\t\t\t\t->1. Addition")
			fmt.Println("\t\t\t\t  2. Subtraction")
			fmt.Println("\t\t\t\t  3. Multiplication")
			fmt.Println("\t\t\t\t  4. Division")
			fmt.Println("\t\t\t\t  5. Power")
			fmt.Println("\t\t\t\t  6. Square Root")
			fmt.Println("\t\t\t\t  7. Exit")
		} else if op == 1 {
			fmt.Println("\t\t\t\t  1. Addition")
			fmt.Println("\t\t\t\t->2. Subtraction")
			fmt.Println("\t\t\t\t  3. Multiplication")
			fmt.Println("\t\t\t\t  4. Division")
			fmt.Println("\t\t\t\t  5. Power")
			fmt.Println("\t\t\t\t  6. Square Root")
			fmt.Println("\t\t\t\t  7. Exit")
		} else if op == 2 {
			fmt.Println("\t\t\t\t  1. Addition")
			fmt.Println("\t\t\t\t  2. Subtraction")
			fmt.Println("\t\t\t\t->3. Multiplication")
			fmt.Println("\t\t\t\t  4. Division")
			fmt.Println("\t\t\t\t  5. Power")
			fmt.Println("\t\t\t\t  6. Square Root")
			fmt.Println("\t\t\t\t  7. Exit")
		} else if op == 3 {
			fmt.Println("\t\t\t\t  1. Addition")
			fmt.Println("\t\t\t\t  2. Subtraction")
			fmt.Println("\t\t\t\t  3. Multiplication")
			fmt.Println("\t\t\t\t->4. Division")
			fmt.Println("\t\t\t\t  5. Power")
			fmt.Println("\t\t\t\t  6. Square Root")
			fmt.Println("\t\t\t\t  7. Exit")
		} else if op == 4 {
			fmt.Println("\t\t\t\t  1. Addition")
			fmt.Println("\t\t\t\t  2. Subtraction")
			fmt.Println("\t\t\t\t  3. Multiplication")
			fmt.Println("\t\t\t\t  4. Division")
			fmt.Println("\t\t\t\t->5. Power")
			fmt.Println("\t\t\t\t  6. Square Root")
			fmt.Println("\t\t\t\t  7. Exit")
		} else if op == 5 {
			fmt.Println("\t\t\t\t  1. Addition")
			fmt.Println("\t\t\t\t  2. Subtraction")
			fmt.Println("\t\t\t\t  3. Multiplication")
			fmt.Println("\t\t\t\t  4. Division")
			fmt.Println("\t\t\t\t  5. Power")
			fmt.Println("\t\t\t\t->6. Square Root")
			fmt.Println("\t\t\t\t  7. Exit")
		} else if op == 6 {
			fmt.Println("\t\t\t\t  1. Addition")
			fmt.Println("\t\t\t\t  2. Subtraction")
			fmt.Println("\t\t\t\t  3. Multiplication")
			fmt.Println("\t\t\t\t  4. Division")
			fmt.Println("\t\t\t\t  5. Power")
			fmt.Println("\t\t\t\t  6. Square Root")
			fmt.Println("\t\t\t\t->7. Exit")
		}
		fmt.Println("\n\n\tPress a digit from 1 to 7 and then confirm with 'enter' to point a option.\n\tPress the digit 8 and then confirm with 'enter' to select the operation.\n\n")
		fmt.Println("\t\t\t\t\tMr. Z!©\n")
		fmt.Scan(&button)
		if button[0] >= '1' && button[0] <= '7' {
			op = int(button[0]) - 49
		} else if button[0] == '8' {
			if op == 0 {
				addition()
			} else if op == 1 {
				subtraction()
			} else if op == 2 {
				multiplication()
			} else if op == 3 {
				division()
			} else if op == 4 {
				power()
			} else if op == 5 {
				squareRoot()
			} else if op == 6 {
				saida()
			}
		}
		clear()
	}
}

func main() {
	menu()
}
