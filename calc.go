package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var conversions = []struct {
	value int
	digit string
}{
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{8, "VIII"},
	{7, "VII"},
	{6, "VI"},
	{5, "V"},
	{4, "IV"},
	{3, "III"},
	{2, "II"},
	{1, "I"},
}
var a, b int
var operators = map[string]func() int{
	"+": func() int { return a + b },
	"-": func() int { return a - b },
	"*": func() int { return a * b },
	"/": func() int { return a / b },
}

func main() {
	fmt.Println("Введите значение:")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.TrimSpace(strings.ToUpper(input))
	input = strings.ReplaceAll(input, " ", "")
	var operator string
	numbers := make([]int, 0)
	romans := make([]string, 0)
	romToInt := make([]int, 0)
	var bezoper []string
	var stringElem, intToRom int
	var roman strings.Builder
	for _, val := range input {
		for idx := range operators {
			if idx == string(val) {
				operator += idx
				bezoper = strings.Split(input, operator)
			}
		}
	}
	if len(operator) > 1 {
		panic("Формат математической операции не удовлетворяет заданию - два операнда и один оператор")
	} else if len(operator) < 1 {
		panic("Строка не является математической операцией")
	}
	for _, val := range bezoper {
		num, err := strconv.Atoi(val)
		if err != nil {
			stringElem++
			romans = append(romans, val)
		} else {
			numbers = append(numbers, num)
		}
	}
	if stringElem == 0 {
		if val, ok := operators[operator]; ok && numbers[0] > 0 && numbers[0] < 11 && numbers[1] > 0 && numbers[1] < 11 {
			a, b = numbers[0], numbers[1]
			fmt.Println(val())
		} else {
			panic("Числа должны быть от 1 до 10")
		}
	} else if stringElem == 1 {
		panic("Используются одновременно разные системы счисления")
	} else if stringElem == 2 {
		for _, elem := range romans {
			for _, conversion := range conversions {
				if elem == conversion.digit {
					romToInt = append(romToInt, conversion.value)
				}
			}
		}
		if len(romToInt) < 2 || romToInt[0] > 10 || romToInt[1] > 10 {
			panic("Числа должны быть от 1 до 10")
		}
		if val, ok := operators[operator]; ok {
			a, b = romToInt[0], romToInt[1]
			intToRom = val()
		}
		if intToRom <= 0 {
			panic("Результат работы меньше единицы")
		}
		for _, conversion := range conversions {
			for intToRom >= conversion.value {
				roman.WriteString(conversion.digit)
				intToRom -= conversion.value
			}
		}
		fmt.Println(roman.String())
	}
}
