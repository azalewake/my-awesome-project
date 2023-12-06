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
	"I":        1,
	"II":       2,
	"III":      3,
	"IV":       4,
	"V":        5,
	"VI":       6,
	"VII":      7,
	"VIII":     8,
	"IX":       9,
	"X":        10,
	"XI":       11,
	"XII":      12,
	"XIII":     13,
	"XIV":      14,
	"XV":       15,
	"XVI":      16,
	"XVII":     17,
	"XVIII":    18,
	"XIX":      19,
	"XX":       20,
	"XXI":      21,
	"XXII":     22,
	"XXIII":    23,
	"XXIV":     24,
	"XXV":      25,
	"XXVI":     26,
	"XXVII":    27,
	"XXVIII":   28,
	"XXIX":     29,
	"XXX":      30,
	"XXXI":     31,
	"XXXII":    32,
	"XXXIII":   33,
	"XXXIV":    34,
	"XXXV":     35,
	"XXXVI":    36,
	"XXXVII":   37,
	"XXXVIII":  38,
	"XXXIX":    39,
	"XL":       40,
	"XLI":      41,
	"XLII":     42,
	"XLIII":    43,
	"XLIV":     44,
	"XLV":      45,
	"XLVI":     46,
	"XLVII":    47,
	"XLVIII":   48,
	"XLIX":     49,
	"L":        50,
	"LI":       51,
	"LII":      52,
	"LIII":     53,
	"LIV":      54,
	"LV":       55,
	"LVI":      56,
	"LVII":     57,
	"LVIII":    58,
	"LIX":      59,
	"LX":       60,
	"LXI":      61,
	"LXII":     62,
	"LXIII":    63,
	"LXIV":     64,
	"LXV":      65,
	"LXVI":     66,
	"LXVII":    67,
	"LXVIII":   68,
	"LXIX":     69,
	"LXX":      70,
	"LXXI":     71,
	"LXXII":    72,
	"LXXIII":   73,
	"LXXIV":    74,
	"LXXV":     75,
	"LXXVI":    76,
	"LXXVII":   77,
	"LXXVIII":  78,
	"LXXIX":    79,
	"LXXX":     80,
	"LXXXI":    81,
	"LXXXII":   82,
	"LXXXIII":  83,
	"LXXXIV":   84,
	"LXXXV":    85,
	"LXXXVI":   86,
	"LXXXVII":  87,
	"LXXXVIII": 88,
	"LXXXIX":   89,
	"XC":       90,
	"XCI":      91,
	"XCII":     92,
	"XCIII":    93,
	"XCIV":     94,
	"XCV":      95,
	"XCVI":     96,
	"XCVII":    97,
	"XCVIII":   98,
	"XCIX":     99,
	"C":        100,
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
	11:  "XI",
	12:  "XII",
	13:  "XIII",
	14:  "XIV",
	15:  "XV",
	16:  "XVI",
	17:  "XVII",
	18:  "XVIII",
	19:  "XIX",
	20:  "XX",
	21:  "XXI",
	22:  "XXII",
	23:  "XXIII",
	24:  "XXIV",
	25:  "XXV",
	26:  "XXVI",
	27:  "XXVII",
	28:  "XXVIII",
	29:  "XXIX",
	30:  "XXX",
	31:  "XXXI",
	32:  "XXXII",
	33:  "XXXIII",
	34:  "XXXIV",
	35:  "XXXV",
	36:  "XXXVI",
	37:  "XXXVII",
	38:  "XXXVIII",
	39:  "XXXIX",
	40:  "XL",
	41:  "XLI",
	42:  "XLII",
	43:  "XLIII",
	44:  "XLIV",
	45:  "XLV",
	46:  "XLVI",
	47:  "XLVII",
	48:  "XLVIII",
	49:  "XLIX",
	50:  "L",
	51:  "LI",
	52:  "LII",
	53:  "LIII",
	54:  "LIV",
	55:  "LV",
	56:  "LVI",
	57:  "LVII",
	58:  "LVIII",
	59:  "LIX",
	60:  "LX",
	61:  "LXI",
	62:  "LXII",
	63:  "LXIII",
	64:  "LXIV",
	65:  "LXV",
	66:  "LXVI",
	67:  "LXVII",
	68:  "LXVIII",
	69:  "LXIX",
	70:  "LXX",
	71:  "LXXI",
	72:  "LXXII",
	73:  "LXXIII",
	74:  "LXXIV",
	75:  "LXXV",
	76:  "LXXVI",
	77:  "LXXVII",
	78:  "LXXVIII",
	79:  "LXXIX",
	80:  "LXXX",
	81:  "LXXXI",
	82:  "LXXXII",
	83:  "LXXXIII",
	84:  "LXXXIV",
	85:  "LXXXV",
	86:  "LXXXVI",
	87:  "LXXXVII",
	88:  "LXXXVIII",
	89:  "LXXXIX",
	90:  "XC",
	91:  "XCI",
	92:  "XCII",
	93:  "XCIII",
	94:  "XCIV",
	95:  "XCV",
	96:  "XCVI",
	97:  "XCVII",
	98:  "XCVIII",
	99:  "XCIX",
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