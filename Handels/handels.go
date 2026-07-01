package handels

import (
	"encoding/json"
	"fmt"
	"time"

	datework "micron/DateWork"
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
		fmt.Println("Не удалось перевести структуру в json")
	} else{
		w.WriteHeader(http.StatusCreated)
		w.Write(hResponse)
	}
	}
}

func DateHandler(w http.ResponseWriter, r *http.Request){
	method := r.Method
	if method == http.MethodPost{
		var jsonDate datework.DateJson
		if err := json.NewDecoder(r.Body).Decode(&jsonDate); err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if err := datework.DateValidator(jsonDate); err != nil{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		clearDate := datework.RemakeDate(jsonDate)
		go func(clearDate time.Time){
			if err := notifications.SendNotificationOnDate(clearDate, jsonDate.Text); err != nil{
				fmt.Println(err)
			}
		}(clearDate)

		hResponse, err := json.MarshalIndent(clearDate, "", "    ")
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Таймер успешно запущен"))
		w.Write(hResponse)
	}
}

