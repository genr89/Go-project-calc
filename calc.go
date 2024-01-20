package main

import (
	"fmt"
)

// Римские числа и их значения
var romanNumerals = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// Функция для преобразования римских чисел в арабские
func romanToArabic(roman string) int {
	result := 0

	for i := 0; i < len(roman); i++ {
		currentValue := romanNumerals[rune(roman[i])]

		if i+1 < len(roman) {
			nextValue := romanNumerals[rune(roman[i+1])]

			if currentValue < nextValue {
				result += nextValue - currentValue
				i++ // Пропустить следующий символ, так как уже учтен
			} else {
				result += currentValue
			}
		} else {
			result += currentValue
		}
	}

	return result
}

// Функция для преобразования арабских чисел в римские
func arabicToRoman(arabic int) string {
	result := ""

	for _, numeral := range []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	} {
		for arabic >= numeral.Value {
			result += numeral.Symbol
			arabic -= numeral.Value
		}
	}

	return result
}

// Функция для выполнения операции сложения римских чисел
func addRomanNumbers(roman1, roman2 string) string {
	arabic1 := romanToArabic(roman1)
	arabic2 := romanToArabic(roman2)
	sum := arabic1 + arabic2

	return arabicToRoman(sum)
}

func main() {
	// Пример использования
	romanNumber1 := "II"
	romanNumber2 := "II"

	sum := addRomanNumbers(romanNumber1, romanNumber2)

	fmt.Printf("%s + %s = %s\n", romanNumber1, romanNumber2, sum)
}
