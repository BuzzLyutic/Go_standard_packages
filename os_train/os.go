package os_train

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Пакет os предоставляет платформонезависимый интерфейс для работы с операционной системой.
// Пакет io предоставляет базовые интерфейсы для ввода/вывода данных.

func CreateFile() {
	newFile, err := os.Create("NewFile.txt")
	if err != nil {
		// Обработка ошибки
		return
	}
	fmt.Println("Файл создан успешно!")
	defer newFile.Close()

	newFile.WriteString("Данный файл и его содержимое созданы при помощи пакета os языка Go!")
	fmt.Println("Новая запись в файл...")
	file, err := os.Open("NewFile.txt")
	if err != nil {
		// Обработка ошибки
	}
	content, err := io.ReadAll(file)
	if err != nil {
		// Обработка ошибки
	}
	fmt.Printf("Содержимое файла NewFile: %s\n", content)
	// Проверка содержимого
	contains := strings.Contains(string(content), "Go")
	fmt.Println(contains)
	hasPrefix := strings.HasPrefix("Тестовая строка", "Тест")
	fmt.Println(hasPrefix)
	// Преобразования строк
	//upper := strings.ToUpper(string(content))
	//lower := strings.ToLower(upper)
	// Разделение и объединение
	//split := strings.Split("abc", " ")
	//join := strings.Join([]string{"a", "b", "c"}, ",")
	// Замена
	//replace := strings.Replace("ZZZZ", "Z", "X", 3) // XXXZ
	//replaceAll := strings.ReplaceAll("YYYY", "Y", "Z")
}
