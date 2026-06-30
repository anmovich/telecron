package handels

import (
	"encoding/json"
	"io"
	notifications "micron/Notifications"
	timer "micron/Timer"
	"net/http"

)

func TimerHandler(w http.ResponseWriter, r *http.Request){
	reqBody, err := io.ReadAll(r.Body)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var t timer.Time
	if err := json.Unmarshal(reqBody, &t); err != nil{
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

