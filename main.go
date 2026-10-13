package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Привет, это умный калькулятор(почти)")
	var number1 float64
	var number2 float64
	var action string

	fmt.Print("Введи 1 число: ")
	_, err := fmt.Scan(&number1)
	if err != nil {
		fmt.Println("Введено не число")
	}
	if number1 == math.Trunc(number1) {
		fmt.Printf("Введено целое число: %.0f\n", number1)
	} else {
		fmt.Printf("Введено дробное число: %g\n", number1)
	}

	fmt.Print("Введите 2 число: ")
	_, err = fmt.Scan(&number2)
	if err != nil {
		fmt.Println("Введено не число")
	}
	if number2 == math.Trunc(number2) {
		fmt.Printf("Введено целое число: %.0f\n", number2)
	} else {
		fmt.Printf("Введено дробное число: %g\n", number2)
	}

	fmt.Print("Выберите действие над числами(*, /, +, -): ")
	fmt.Scan(&action)

	switch action {
	case "*":
		fmt.Printf("%g * %g = %g\n", number1, number2, number1*number2)
	case "/":
		if number2 == 0 {
			fmt.Println("Ошибка, деление на ноль")
		} else {
			fmt.Printf("%g / %g = %g\n", number1, number2, number1/number2)
		}
	case "+":
		fmt.Printf("%g + %g = %g\n", number1, number2, number1+number2)
	case "-":
		fmt.Printf("%g - %g = %g\n", number1, number2, number1-number2)
	default:
		fmt.Println("Неизвестное действие")
	}
}
