package handlers

import (
	"clickercopro/app/logic"
	"encoding/json"
	"fmt"
	"net/http"
)

var PointData []byte

func HandlerClickevent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var Data map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&Data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		PointData, err = json.Marshal(Data)
		if err != nil {
			fmt.Println("Ошибка сериализации jsona", err)
			return
		}

		// Обработка данных
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "Успех", "message": "Данные получены!"})
		fmt.Println("Данные получены")
	} else {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
	}
	logic.Point(PointData)
}
