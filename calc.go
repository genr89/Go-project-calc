package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция для конвертации римских цифр в арабские
func romanToArabic(roman string) (int, error) {
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
	prevValue := 0
	for _, r := range roman {
		value, found := romanNumerals[string(r)]
		if !found {
			return 0, errors.New("некорректные римские цифры")
		}
		if prevValue < value {
			arabic += value - 2*prevValue
		} else {
			arabic += value
		}
		prevValue = value
	}
	return arabic, nil
}

// Функция для конвертации арабских чисел в римские
func arabicToRoman(arabic int) (string, error) {
	if arabic <= 0 || arabic > 3999 {
		return "", errors.New("число вне диапазона (1-3999)")
	}

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
	return result.String(), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Простой Римский/Арабский Калькулятор")
	fmt.Println("Введите выражение в формате: число оператор число (например: III + II или 3 + 2)")

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

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			// Попробуем интерпретировать ввод как римские числа
			num1, err = romanToArabic(parts[0])
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}
		}

		operator := parts[1]

		num2, err := strconv.Atoi(parts[2])
		if err != nil {
			// Попробуем интерпретировать ввод как римские числа
			num2, err = romanToArabic(parts[2])
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}
		}

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
			if num2 == 0 {
				fmt.Println("Ошибка: деление на ноль")
				continue
			}
			result = num1 / num2
		default:
			fmt.Println("Неподдерживаемый оператор.")
			continue
		}

		// Вывод результата в римских цифрах, если ввод был римским
		if _, err := strconv.Atoi(parts[0]); err != nil {
			romanResult, err := arabicToRoman(result)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}
			fmt.Println(romanResult)
		} else {
			fmt.Println(result)
		}
	}
}
