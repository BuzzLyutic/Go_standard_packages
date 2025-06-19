package task1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func Task1 () {
	file, err := os.Open("task1/task1_text.txt") // Открываем файл
	if err != nil { // Обработка ошибки открытия файла
		log.Fatal(err)
	}
	defer file.Close() // Гарантированное закрытие файла по окончании работы функции

	scanner := bufio.NewScanner(file) // Используем bufio.Scanner, потому что он считывает информацию порциями (буфферами), экономя память
	scanner.Split(bufio.ScanWords) // ScanWords обрабатывает юникод-разделители (пробелы, табуляции, новые строки) и работает быстрее ручного разбиения.
	words := make(map[string]int) // Карта для хранения количества вхождений слов
	for scanner.Scan() { // Итерация по токенам (словам). 
		token := scanner.Text() // Считанный токен доступен через Scanner.Bytes или Scanner.Text

		trimmed := strings.TrimFunc(token, func(r rune) bool { // Обрезаем пунктуацию
			return unicode.IsPunct(r) && r != '\''
		})

		if trimmed == "" { // Пропускаем пустые строки после обрезки (например, "---" → "")
			continue
		}

		word := strings.ToLower(trimmed) // Приводим к нижнему регистру
		words[word]++ // Увеличиваем счетчик для данного слова
	}

	if err := scanner.Err(); err != nil { // Ошибки сканирования
		log.Fatal(err)
	}

	fmt.Println("\nСтатистика употребления слов:")
    fmt.Println("============================")

	for word, count := range words {
    fmt.Printf("%-20s: %5d\n", word, count)
	}
}
