package main

import (
	"fmt"
	"strings"
)

// Функция для преобразования римских чисел в десятичные
func romanToDecimal(roman string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var result int
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value := romanMap[string(roman[i])]
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}

	return result
}

// Функция для преобразования десятичных чисел в римские
func decimalToRoman(decimal int) string {
	romanMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	var result strings.Builder
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for _, value := range values {
		for decimal >= value {
			result.WriteString(romanMap[value])
			decimal -= value
		}
	}

	return result.String()
}

func main() {
	var input1, input2 string
	var operator rune

	fmt.Println("Привет! Добро пожаловать в римский калькулятор.")
	fmt.Println("Введите первое римское число:")
	fmt.Scanln(&input1)
	fmt.Println("Введите оператор (+, -, *, /):")
	fmt.Scanln(&operator)
	fmt.Println("Введите второе римское число:")
	fmt.Scanln(&input2)

	number1 := romanToDecimal(input1)
	number2 := romanToDecimal(input2)

	var result int

	switch operator {
	case '+':
		result = number1 + number2
	case '-':
		result = number1 - number2
	case '*':
		result = number1 * number2
	case '/':
		if number2 == 0 {
			fmt.Println("Ошибка: деление на ноль")
			return
		}
		result = number1 / number2
	default:
		fmt.Println("Ошибка: неверный оператор")
		return
	}

	fmt.Println("Результат:", decimalToRoman(result))
}
