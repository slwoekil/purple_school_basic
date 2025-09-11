package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operationType := getOperationType()

	operations := getOperations()

	switch operationType {
	case "SUM":
		SUM := calculateSUM(operations)
		fmt.Println("Сумма чисел:", SUM)
	case "AVG":
		AVG := calculateAVG(operations)
		fmt.Println("Среднее значение:", AVG)
	case "MED":
		MED := calculateMED(operations)
		fmt.Println("Медианное значение:", MED)
	}
}

func getOperationType() string {
	var operationType string

	for {
		fmt.Println("Какую операцию вы бы хотели выполнить? (AVG/SUM/MED)")
		if _, err := fmt.Scanln(&operationType); err != nil {
			fmt.Println("Ошибка ввода", err)
			break
		}

		operationType = strings.ToUpper(operationType)

		switch operationType {
		case "AVG", "SUM", "MED":
		default:
			fmt.Println("Неизвестная операция")
			continue
		}
		break
	}
	return operationType

}

func getOperations() []float64 {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите числа через запятую")

		scanner.Scan()

		input := scanner.Text()
		input = strings.ReplaceAll(input, ",", " ")

		parts := strings.Fields(input)

		operations := make([]float64, 0, len(parts))

		for _, value := range parts {
			value, err := strconv.ParseFloat(value, 64)
			if err != nil {
				fmt.Println("Ошибка обработки значения")
				continue
			}
			operations = append(operations, value)
		}

		return operations
	}
}

func calculateSUM(operations []float64) float64 {
	sum := 0.0

	for _, value := range operations {
		sum += value
	}
	return sum
}
func calculateAVG(operations []float64) float64 {
	sum := calculateSUM(operations)
	return sum / float64(len(operations))
}

func calculateMED(operations []float64) float64 {
	sort.Float64s(operations)
	if len(operations)%2 == 0 {
		return (operations[len(operations)/2] + operations[len(operations)/2+1]) / 2
	} else {
		return operations[len(operations)/2]
	}
}
