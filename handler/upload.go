package handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"dfs/pkg"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	file, header, err := r.FormFile("file") // r.FormFile считывает данные с именем формы "file", где находятся file, header, err
	if err != nil {                         // если код ошибки не равен nill, то срабатывает функция обработки ошибок
		w.WriteHeader(http.StatusBadRequest) //так как мы считали header(мета данные файла, его имя и т.д), выводим ошибку StatusBadRequest, что значит неккоректный запроос
		response := map[string]string{       //формат выведение json ответа
			"error": "don't read",
		}
		json.NewEncoder(w).Encode(response) // вывод json ответа
		return                              // останавливаем функцию
	}
	defer file.Close() //закрываем файл

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	column := r.FormValue("column")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{
			"answer": "couldn't read CSV file",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	columnIndex := -1

	for i, name := range records[0] {
		if name == column {
			columnIndex = i
			break
		}
	}

	var numbers []float64

	for i, record := range records[1:] { //наичнаем со второй строки т.к первая у нас заголовки

		if columnIndex >= len(record) {
			w.WriteHeader(http.StatusBadRequest)
			response := map[string]string{
				"answer": "missing value in CSV row",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		number, err := strconv.ParseFloat(strings.TrimSpace(record[columnIndex]), 64) // достаем значения из колонки записывая в numbers и превращаем str в float64

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := map[string]string{
				"error": fmt.Sprintf("invalid number in row %d", i+2), //ошибка I+2 строка тип
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		numbers = append(numbers, number)

		if len(numbers) == 0 {
			fmt.Println("чисел нет")
			return
		}
	}

	response := map[string]string{ //формат выведение json ответа
		"message":  "File succsesfully taken!",
		"filename": header.Filename,
	}
	json.NewEncoder(w).Encode(response)                          // вывод json ответа
	fmt.Printf("Файл получен, имя файла: %s\n", header.Filename) //вывод названия файла в терминале

	count := len(numbers)
	minimum := pkg.Min(numbers)
	maximum := pkg.Max(numbers)
	mean := pkg.Sum(numbers) / float64(count)

	// Формируем JSON-ответ
	responsend := map[string]any{
		"message":  "File successfully processed",
		"filename": header.Filename,
		"column":   column,
		"count":    count,
		"min":      minimum,
		"max":      maximum,
		"mean":     mean,
	}

	json.NewEncoder(w).Encode(responsend)

	fmt.Printf("Файл обработан: %s\n", header.Filename)

}
