package task2

import (
	"encoding/json" // Библиотека для работы с JSON
	"fmt"
	"log"
	"os"
)

type Grade struct {
	Subject string `json:"subject"`
	Grades []int `json:"grades"`
}

type Student struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	StudyGroup string `json:"study_group"`
	Performance []Grade `json:"performance"`
}

type StudentAvgMarks struct {
	ID int `json:"id"`
	FullName string `json:"full_name"`
	OverallAverage float64 `json:"overall_average"`
	SubjectsAverages map[string]float64 `json:"subjects_average"`
}

func calculateAvg(grades []int) float64 {
	if len(grades) == 0 {
		return 0.0
	}
	sum := 0
	for _, val := range grades {
		sum += val
	}
	return float64(sum) / float64(len(grades))
}

func Task2() {
	// Открываем файл с входными данными
	inputFile, err := os.Open("task2/data.json")
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %v", err)
	}
	defer inputFile.Close()

	// Декодируем JSON из файла в массив студентов
	var students []Student
	decoder := json.NewDecoder(inputFile)
	err = decoder.Decode(&students)
	if err != nil {
		log.Fatalf("Ошибка при чтении JSON: %v", err)
	}

	// Начало обработки данных
	var results []StudentAvgMarks // Массив для хранения рассчитанных средних оценок для каждого студента
	for _, student := range students {
		subjectsAverages := map[string]float64{}
		allGrades := []int{}
		
		for _, performance := range student.Performance {
			// Расчет средней по каждому предмету отдельно
			avg := calculateAvg(performance.Grades)
			subjectsAverages[performance.Subject] = avg
			allGrades = append(allGrades, performance.Grades...)
		}

		result := StudentAvgMarks { // Формирование результата
			ID: student.ID,
			FullName: fmt.Sprintf("%s %s", student.FirstName, student.LastName),
			OverallAverage: calculateAvg(allGrades), // Расчет общей средней по всем предметам
			SubjectsAverages: subjectsAverages,
		}
		results = append(results, result)
	}

	// Создаем файл с результатами
	outputFile, err := os.Create("task2/result.json")
	if err != nil {
		log.Fatalf("Ошибка при создании файла с ответом: %v", err)
	}
	defer outputFile.Close()

	// Кодируем результаты в файл
	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(results); err != nil {
		log.Fatalf("Ошибка при кодировании JSON: %v", err)
	}

	fmt.Println("Обработка файла прошла успешно. Результат записан в файл result.json")

}
