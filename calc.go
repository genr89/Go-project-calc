package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 2 {
		panic("Usage: ./calc \"expression\"")
	}

	expression := os.Args[1]
	var result string
	var err error

	// Убираем кавычки вокруг строк и разбиваем выражение на части
	tokens := parseExpression(expression)

	if len(tokens) != 3 {
		panic("Invalid expression format")
	}

	switch tokens[1] {
	case "+":
		result, err = add(tokens[0], tokens[2])
	case "-":
		result, err = subtract(tokens[0], tokens[2])
	case "*":
		result, err = multiply(tokens[0], tokens[2])
	case "/":
		result, err = divide(tokens[0], tokens[2])
	default:
		panic("Unsupported operation")
	}

	if err != nil {
		panic(err)
	}

	if len(result) > 40 {
		result = result[:40] + "..."
	}

	fmt.Printf("%s\n", result)
}

func parseExpression(expression string) []string {
	tokens := make([]string, 0, 3)
	var token strings.Builder
	inQuotes := false

	for _, r := range expression {
		switch {
		case r == '"':
			inQuotes = !inQuotes
			if !inQuotes {
				tokens = append(tokens, token.String())
				token.Reset()
			}
		case unicode.IsSpace(r) && !inQuotes:
			continue
		case (r == '+' || r == '-' || r == '*' || r == '/') && !inQuotes:
			if token.Len() > 0 {
				tokens = append(tokens, token.String())
				token.Reset()
			}
			tokens = append(tokens, string(r))
		default:
			token.WriteRune(r)
		}
	}

	if token.Len() > 0 {
		tokens = append(tokens, token.String())
	}

	return tokens
}

func add(a, b string) (string, error) {
	if isNumeric(a) && isNumeric(b) {
		num1, _ := strconv.Atoi(a)
		num2, _ := strconv.Atoi(b)
		return strconv.Itoa(num1 + num2), nil
	}

	if isValidString(a) && isValidString(b) {
		return a + b, nil
	}

	return "", fmt.Errorf("Invalid string format")
}

func subtract(a, b string) (string, error) {
	if isNumeric(a) && isNumeric(b) {
		num1, _ := strconv.Atoi(a)
		num2, _ := strconv.Atoi(b)
		return strconv.Itoa(num1 - num2), nil
	}

	if isValidString(a) && isValidString(b) {
		return strings.Replace(a, b, "", -1), nil
	}

	return "", fmt.Errorf("Invalid string format")
}

func multiply(a, b string) (string, error) {
	if isNumeric(a) && isNumeric(b) {
		num1, _ := strconv.Atoi(a)
		num2, _ := strconv.Atoi(b)
		return strconv.Itoa(num1 * num2), nil
	}

	if isValidString(a) {
		n, err := strconv.Atoi(b)
		if err != nil || n < 1 || n > 10 {
			return "", fmt.Errorf("Invalid number format")
		}
		return strings.Repeat(a, n), nil
	}

	return "", fmt.Errorf("Invalid string format")
}

func divide(a, b string) (string, error) {
	if isNumeric(a) && isNumeric(b) {
		num1, _ := strconv.Atoi(a)
		num2, _ := strconv.Atoi(b)
		if num2 == 0 {
			return "", fmt.Errorf("Division by zero")
		}
		return strconv.Itoa(num1 / num2), nil
	}

	if isValidString(a) {
		n, err := strconv.Atoi(b)
		if err != nil || n < 1 || n > 10 {
			return "", fmt.Errorf("Invalid number format")
		}

		substrLen := len(a) / n
		return a[:substrLen], nil
	}

	return "", fmt.Errorf("Invalid string format")
}

func isValidString(s string) bool {
	return len(s) <= 10
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
