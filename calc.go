package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция для конвертации римских цифр в арабские
func romanToArabic(roman string) (int, error) {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	arabic := 0
	prevValue := 0
	for _, r := range roman {
		value, found := romanNumerals[r]
		if !found {
			panic(fmt.Sprintf("некорректные римские цифры: %s", roman))
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
		return "", fmt.Errorf("число вне диапазона (1-3999): %d", arabic)
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
	fmt.Println("Для выхода введите exit")

	for {
		fmt.Print(">> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Программа завершена.")
			return
		}

		// Разделение ввода на числа и оператор
		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("Неправильный формат. Пожалуйста, введите выражение снова.")
			continue
		}

		// Проверка, что вводимые числа либо оба арабские, либо оба римские
		isNum1Roman := isRoman(parts[0])
		isNum2Roman := isRoman(parts[2])
		if (isNum1Roman && !isNum2Roman) || (!isNum1Roman && isNum2Roman) {
			fmt.Println("Ошибка: используйте либо только арабские, либо только римские числа.")
			continue
		}

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			num1, err = romanToArabic(parts[0])
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}
		}

		operator := parts[1]

		num2, err := strconv.Atoi(parts[2])
		if err != nil {
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
		if isNum1Roman {
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

// Функция для проверки, является ли строка римским числом
func isRoman(s string) bool {
	for _, r := range s {
		if r != 'I' && r != 'V' && r != 'X' && r != 'L' && r != 'C' && r != 'D' && r != 'M' {
			return false
		}
	}
	return true
}
