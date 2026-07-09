package handels

import (
	"context"
	"encoding/json"
	"fmt"

	bd "micron/Bd"
	datework "micron/DateWork"
	notifications "micron/Notifications"
	run "micron/Run"
	timer "micron/Timer"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	DB  *pgxpool.Pool
	Ctx context.Context
}

// Функция для запуска таймера
func (h *Handler) TimerHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	//Описание метода POST, запуск таймера
	if method == http.MethodPost {
		var t timer.Time
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		ok, err := t.ValidatorTimer()
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return

		}
		go func(t *timer.Time) {
			if err := notifications.SendNotificationTimer(t); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println(err)
				return
			}
		}(&t)

		hResponse, err := json.MarshalIndent(t, "", "    ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Не удалось перевести структуру в json")
		} else {
			w.WriteHeader(http.StatusCreated)
			w.Write(hResponse)
		}
	}
}

// Date query
func (h *Handler) DateHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	//Method POST
	if method == http.MethodPost {
		var Date datework.DateJson //Создание даты для парсинга в нее
		//Parsing json to variable
		if err := json.NewDecoder(r.Body).Decode(&Date); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		go func(Date datework.DateJson) {
			if err := run.ExecCommand(Date); err != nil {
				fmt.Println(err)
			}
		}(Date)

		if err := bd.WriteTimer(h.Ctx, h.DB, Date); err != nil {
			fmt.Println(err)
		}

		//Marshall date to write response
		hResponse, err := json.MarshalIndent(Date, "", "    ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Таймер успешно запущен"))
		w.Write(hResponse)
	}
}
