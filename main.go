package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	// Получение имени входного файла и имени файла для вывода результатов из аргументов командной строки
	if len(os.Args) < 3 {
		fmt.Println("Необходимо указать имя входного файла и имя файла для вывода результатов")
		return
	}
	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	// Чтение содержимого входного файла
	inputContent, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		fmt.Printf("Ошибка при чтении входного файла: %v\n", err)
		return
	}

	// Парсинг математических выражений и вычисление результатов
	results := parseAndCalculateExpressions(string(inputContent))

	// Запись результатов в файл вывода
	err = writeResultsToFile(outputFileName, results)
	if err != nil {
		fmt.Printf("Ошибка при записи результатов: %v\n", err)
		return
	}

	fmt.Println("Результаты успешно записаны в файл", outputFileName)
}

func parseAndCalculateExpressions(input string) map[string]string {
	results := make(map[string]string)

	// Паттерн для поиска математических выражений в формате "число1+число2=?"
	expressionPattern := regexp.MustCompile(`(\d+)\s*([\+\-\*\/])\s*(\d+)\s*=\?`)

	// Поиск всех математических выражений во входном тексте
	matches := expressionPattern.FindAllStringSubmatch(input, -1)

	// Вычисление результатов и сохранение их в словаре
	for _, match := range matches {
		number1 := match[1]
		operator := match[2]
		number2 := match[3]
		result := calculateExpression(number1, operator, number2)
		expression := fmt.Sprintf("%s%s%s=?", number1, operator, number2)
		results[expression] = result
	}

	return results
}

func calculateExpression(number1, operator, number2 string) string {
	// Простой калькулятор для вычисления математического выражения
	switch operator {
	case "+":
		return fmt.Sprintf("%d", parseInt(number1)+parseInt(number2))
	case "-":
		return fmt.Sprintf("%d", parseInt(number1)-parseInt(number2))
	case "*":
		return fmt.Sprintf("%d", parseInt(number1)*parseInt(number2))
	case "/":
		return fmt.Sprintf("%.2f", parseFloat(number1)/parseFloat(number2))
	default:
		return ""
	}
}

func parseInt(str string) int {
	var result int
	_, _ = fmt.Sscanf(str, "%d", &result)
	return result
}

func parseFloat(str string) float64 {
	var result float64
	_, _ = fmt.Sscanf(str, "%f", &result)
	return result
}

func writeResultsToFile(filename string, results map[string]string) error {
	// Открытие файла вывода для записи (с очисткой содержимого, если файл уже существует)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Запись результатов в файл
	for expression, result := range results {
		output := fmt.Sprintf("%s%s\n", expression, result)
		_, err := writer.WriteString(output)
		if err != nil {
			return err
		}
	}

	// Очистка буфера и сохранение данных на диск
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}