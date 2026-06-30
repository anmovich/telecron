package handels

import (
	"encoding/json"

	notifications "micron/Notifications"
	timer "micron/Timer"
	"net/http"

)

func TimerHandler(w http.ResponseWriter, r *http.Request){
	var t timer.Time
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil{
    w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !timer.CheckType(t){
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Неверный тип времени"))
			return

	}
		go func(t *timer.Time){
		if err := notifications.SendNotificationTimer(t); err != nil{
			w.WriteHeader(http.StatusBadRequest)	
		  return
		} 	
		}(&t)
	w.Write([]byte("Таймер запущен успешно, ждите хуйню"))

}

