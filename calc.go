package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Функция для конвертации римских цифр в арабские
func romanToArabic(roman string) int {
	romanNumerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	arabic := 0
	for i := 0; i < len(roman); i++ {
		if i > 0 && romanNumerals[string(roman[i])] > romanNumerals[string(roman[i-1])] {
			arabic += romanNumerals[string(roman[i])] - 2*romanNumerals[string(roman[i-1])]
		} else {
			arabic += romanNumerals[string(roman[i])]
		}
	}
	return arabic
}

// Функция для конвертации арабских чисел в римские
func arabicToRoman(arabic int) string {
	romanNumerals := []struct {
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
	}
	var result strings.Builder
	for _, numeral := range romanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Простой Римский Калькулятор")
	fmt.Println("Введите выражение в формате: число оператор число (например: III + II)")

	for {
		fmt.Print(">> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Разделение ввода на числа и оператор
		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("Неправильный формат. Пожалуйста, введите выражение снова.")
			continue
		}

		num1 := romanToArabic(parts[0])
		operator := parts[1]
		num2 := romanToArabic(parts[2])

		// Выполнение операции
		var result int
		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			result = num1 / num2
		default:
			fmt.Println("Неподдерживаемый оператор.")
			continue
		}

		// Вывод результата в римских цифрах
		fmt.Println(arabicToRoman(result))
	}
}
