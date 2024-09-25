package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Введите выражение (например, I + II или 2 + 3): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	parts := strings.Fields(input)
	if len(parts) != 3 {
		fmt.Println("Неверный формат ввода. Используйте 'a оператор b'.")
		return
	}

	aStr, operator, bStr := parts[0], parts[1], parts[2]

	a, errA := strconv.Atoi(aStr)
	if errA != nil {
		fmt.Printf("Ошибка преобразования числа A: %v\n", errA)
		return
	}
	b, errB := strconv.Atoi(bStr)
	if errB != nil {
		fmt.Printf("Ошибка преобразования числа B: %v\n", errB)
		return
	}

	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль.")
			return
		}
		result = a / b
	default:
		fmt.Println("Неизвестный оператор.")
		return
	}

	fmt.Println("Результат:", result)
}
