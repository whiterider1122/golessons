// потенциально может работать с римскими числами влоть до 3999, в настоящий момент ввод каждого числа ограничен 10, значит максимальное римское число 100(С)
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isroman(str1 string, str2 string) bool {
	re := regexp.MustCompile(`^[IVXLCDM]+$`)

	if re.MatchString(str1) && re.MatchString(str2) {
		return true
	} else {
		return false
	}
}

func isdigit(str1 string, str2 string) bool {
	re := regexp.MustCompile(`^[0-9]+$`)

	if re.MatchString(str1) && re.MatchString(str2) {
		return true
	} else {
		return false
	}
}
func int2roman(num int) string {
	if num <= 0 || num > 3999 {
		return "Недопустимое значение"
	}

	// Таблица соответствия арабских и римских чисел
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := strings.Builder{}

	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			num -= vals[i]
			result.WriteString(romans[i])
		}
	}

	return result.String()
}
func roman2int(roman string) int {
	r2i_map := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	for i := 0; i < len(roman); i++ {
		if i == 0 || r2i_map[roman[i]] <= r2i_map[roman[i-1]] {
			res += r2i_map[roman[i]]
		} else {
			res += r2i_map[roman[i]] - 2*r2i_map[roman[i-1]]
		}
	}
	return res
}
func calc(oper1 int, oper2 int, operator rune) int {

	out := -1
	if oper1 > 10 || oper2 > 10 {
		panic("Значение одного из чисел больше 10")
	}

	switch operator {
	case '+':
		out = int(oper1 + oper2)
	case '-':
		out = int(oper1 - oper2)
	case '*':
		out = int(oper1 * oper2)
	case '/':
		out = int(oper1 / oper2)
	}
	return out
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите выражение (exit для выхода):")
		// Считываем строку, пока не будет введена новая строка (\n)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка при чтении строки:", err)
			return
		}

		// Удаляем пробельные символы в начале и в конце строки
		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println("Выход из программы.")
			break
		}
		operators := "+-*/"
		counts := make(map[rune]int)
		var operator rune

		// Подсчитываем количество операторов в строке
		for _, char := range input {
			if strings.ContainsRune(operators, char) {
				counts[char]++
				operator = char
			}
		}

		// Проверяем количество операторов
		operatorCount := 0
		for _, count := range counts {
			operatorCount += count
			if count > 1 {
				fmt.Println("Ошибка: один из операторов встречается больше одного раза.")
				return
			}
		}

		if operatorCount == 1 {
			// Разделяем строку по найденному оператору
			parts := strings.Split(input, string(operator))
			if len(parts) == 2 {
				oper1 := strings.TrimSpace(parts[0])
				oper2 := strings.TrimSpace(parts[1])

				if isroman(oper1, oper2) {
					res := calc(roman2int(oper1), roman2int(oper2), operator)
					if res == -1 {
						fmt.Println("Ошибка: Римские числа не могут быть отрицательными.")
						return
					}
					fmt.Println("Результат:", oper1, string(operator), oper2, "=", int2roman(res))
				} else if isdigit(oper1, oper2) {
					num1, err := strconv.Atoi(oper1)
					if err != nil {
						fmt.Println("Ошибка:", err)
						return
					}
					num2, err := strconv.Atoi(oper2)
					if err != nil {
						fmt.Println("Ошибка:", err)
						return
					}
					fmt.Println("Результат:", oper1, string(operator), oper2, "=", calc(num1, num2, operator))
				} else {
					fmt.Println("Ошибка: Оба числа должны быть либо римскими либо арабскими.")
				}
			} else {
				fmt.Println("Должно быть только два числа.")
			}
		} else {
			fmt.Println("Ошибка: строка должна содержать ровно один оператор (+, -, *, /).")
		}
	}
}
