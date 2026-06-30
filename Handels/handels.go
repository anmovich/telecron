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
		go func(t *timer.Time){
		if err := notifications.SendNotificationTimer(t); err != nil{
			w.WriteHeader(http.StatusInternalServerError)
		} else{
			w.Write([]byte("Таймер запущен успешно, ждите хуйню"))
		}
	}(&t)

}

