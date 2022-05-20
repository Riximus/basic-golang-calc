package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	num1, num2, op := input()
	result := calc(num1, num2, op)
	fmt.Println(result)
}

// TODO learn how to give multiple returns for the calc-function
func input() (float64, float64, rune) {
	scanner := bufio.NewScanner(os.Stdin)
	rune := bufio.NewReader(os.Stdin)

	fmt.Println("CALCULATOR")
	fmt.Println("------------")

	fmt.Print("Number -> ")
	scanner.Scan()
	num1, _ := strconv.ParseFloat(scanner.Text(), 0)
	println(num1)

	fmt.Print("Operator -> ")
	op, _, err := rune.ReadRune()
	println(op)
	if err != nil {
		fmt.Println(err)
	}
	if _, err := strconv.Atoi(string(op)); err == nil {
		fmt.Printf("This is not an operator %c \n", op)
		os.Exit(0)
	}

	fmt.Print("Number -> ")
	scanner.Scan()
	num2, _ := strconv.ParseFloat(scanner.Text(), 0)
	println(num2)
	return num1, num2, op
}

func calc(num1 float64, num2 float64, op rune) float64 {
	var result float64

	switch op {
	case '+':
		result = num1 + num2
		break
	case '-':
		result = num1 - num2
		break
	case '*':
		result = num1 * num2
		break
	case '/':
		result = num1 / num2
		break
	default:
		print("default")

	}

	return result
}
