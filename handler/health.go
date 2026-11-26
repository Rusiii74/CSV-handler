package handler

import (
	"encoding/json"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) { //функция-обработчик вызывающяяся запросом маршрутизатора, создает ответ клиенту

	w.Header().Set("Content-Type", "application/json") // Content-Type сообщает клиенту, в каком формате сервер отправляет данные. application/json объявляет в каком формате сервер отправляет данные

	response := map[string]string{ // то что будет находиться в ответе,
		"status": "ok",
	}

	json.NewEncoder(w).Encode(response) // json.NewEncoder(w) создаёт кодировщик JSON, который записывает данные в w.Encode преобразует Go-значение в JSON и отправляет его в HTTP-ответ. записывает полученный JSON непосредственно в HTTP-ответ.
}
