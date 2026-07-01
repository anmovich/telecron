package handels

import (
	"encoding/json"
	"fmt"

	notifications "micron/Notifications"
	timer "micron/Timer"
	"net/http"
)

//Функция для запуска таймера
func TimerHandler(w http.ResponseWriter, r *http.Request){
	method := r.Method

	//Описание метода POST, запуск таймера
	if method == http.MethodPost{
	var t timer.Time
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil{
    w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	ok, err := t.ValidatorTimer()
	if !ok{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return

	}
		go func(t *timer.Time){
		if err := notifications.SendNotificationTimer(t); err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
		  return
		} 	
		}(&t)
		
	hResponse, err := json.MarshalIndent(t, "", "    ")
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Не удалось перевести структуру")
	} else{
		w.WriteHeader(http.StatusCreated)
		w.Write(hResponse)
	}
	}
}

