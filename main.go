package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringRomanToInt(roman string) (int, error) {
	romanMap := map[byte]int{'I': 1, 'V': 5, 'X': 10}

	if len(roman) == 0 {
		return 0, fmt.Errorf("Пустая строка")
	}

	for _, char := range roman {
		if _, ok := romanMap[byte(char)]; !ok {
			return 0, fmt.Errorf("Не римсккая цифра")
		}
	}

	result := romanMap[roman[len(roman)-1]]
	for i := len(roman) - 2; i >= 0; i-- {
		if romanMap[roman[i]] < romanMap[roman[i+1]] {
			result -= romanMap[roman[i]]
		} else {
			result += romanMap[roman[i]]
		}
	}

	return result, nil
}

func intToRoman(num int) string {
	romanValues := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	intValues := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := ""
	i := 0
	for num > 0 {
		for intValues[i] <= num {
			result += romanValues[i]
			num -= intValues[i]
		}
		i++
	}

	return result
}

func calc(num1 int, operator string, num2 int) int {
	if num1 < 1 || num1 > 10 {
		panic("число 1 меньше 1 или больше 10")
	}
	if num2 < 1 || num2 > 10 {
		panic("число 2 меньше 1 или больше 10")
	}

	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "/":
		return num1 / num2
	case "*":
		return num1 * num2
	default:
		panic("Это действие нельзя сделать")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	equation, _ := reader.ReadString('\n')

	terms := strings.Split(equation, " ")
	if len(terms) != 3 {
		panic("Арефметическая операция не удовлетворяет условиям")
	}
	operand1Str := terms[0]
	operator := terms[1]
	operand2Str := strings.TrimSpace(terms[2])
	equationRoman := false
	var operand1int int
	var operand2int int
	operand1int, err := strconv.Atoi(operand1Str)
	if err != nil {
		operand1int, err = stringRomanToInt(operand1Str)
		if err != nil {
			panic("Ведите корректное слагаймое или целое число!")
		}
		equationRoman = true
	}
	if equationRoman {
		roman1Num, err1 := stringRomanToInt(operand1Str)
		roman2Num, err2 := stringRomanToInt(operand2Str)
		if err1 != nil || err2 != nil {
			panic("Не римские цифры")
		}
		resultInt := calc(roman1Num, operator, roman2Num)
		if resultInt < 1 {
			panic("Ответ в риских цифрах должен быть положительным")
		}
		resultRoman := intToRoman(resultInt)
		fmt.Printf("result: %s", resultRoman)
	}
	if equationRoman == false {
		operand2int, err = strconv.Atoi(operand2Str)
		if err != nil {
			panic("Вторая переменная не является целым числом")
		}
		result := calc(operand1int, operator, operand2int)
		fmt.Printf("result: %d", result)
	}
}
