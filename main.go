package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Определить какие цифры арабские и какие римские
var romanNumerals = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"XX":   20,
	"XXX":  30,
	"XL":   40,
	"L":    50,
	"LX":   60,
	"LXX":  70,
	"LXXX": 80,
	"XC":   90,
	"C":    100,
}

var romanNumeralsReverse = map[int]string{
	1:   "I",
	2:   "II",
	3:   "III",
	4:   "IV",
	5:   "V",
	6:   "VI",
	7:   "VII",
	8:   "VIII",
	9:   "IX",
	10:  "X",
	20:  "XX",
	30:  "XXX",
	40:  "XL",
	50:  "L",
	60:  "LX",
	70:  "LXX",
	80:  "LXXX",
	90:  "XC",
	100: "C",
}

// Преобразование арабских чисел в римские цифры
func arabicToRoman(arabic int) string {
	if arabic <= 0 {
		return "Ошибка: отрицательный результат для римских цифр"
	}

	if roman, ok := romanNumeralsReverse[arabic]; ok {
		return roman
	}
	return strconv.Itoa(arabic)
}

// Разобрать строку input на числа. Разобраться с ошибками.
func parseOperand(input string) (int, error) {
	if arabic, err := strconv.Atoi(input); err == nil {
		if arabic >= 1 && arabic <= 10 {
			return arabic, nil
		} else {
			return 0, fmt.Errorf("невозможно принять число: %d, введите число от 1 до 10 включительно", arabic)
		}
	}

	if roman, ok := romanNumerals[input]; ok {
		return roman, nil
	}

	return 0, fmt.Errorf("невозможно преобразовать в число: %s", input)
}

// Проверить строку на соотв римским числам
func isRomanNumeral(input string) bool {
	for _, r := range input {
		if _, ok := romanNumerals[string(r)]; !ok {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите выражение (или 'exit' для завершения): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Fields(input)
	if len(parts) != 3 {
		fmt.Println("Некорректные данные")
		return
	}

	num1, err := parseOperand(parts[0])
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	operator := parts[1]

	num2, err := parseOperand(parts[2])
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Проверка на операции между римскими и арабскими цифрами
	if (isRomanNumeral(parts[0]) && isRomanNumeral(parts[2])) || (!isRomanNumeral(parts[0]) && !isRomanNumeral(parts[2])) {
		var result int
		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 != 0 {
				result = num1 / num2
			} else {
				fmt.Println("Ошибка: деление на ноль")
				return
			}
		default:
			fmt.Println("Ошибка: некорректный оператор")
			return
		}

		if isRomanNumeral(parts[0]) {
			fmt.Println(arabicToRoman(result))
		} else {
			fmt.Println(result)
		}
	} else {
		fmt.Println("Ошибка: операции между римскими и арабскими цифрами запрещены")
	}
}
