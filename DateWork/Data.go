package datework

import (
	"errors"
	"time"
)

type DateJson struct {
	ID          int
	Date        time.Time `json:"time"`
	Description string    `json:"description"`
	Command     string    `json:"command"`
	Args        []string  `json:"args"`
	Repeat      bool      `json:"repeating"`
	TimeCreated time.Time `json:"time_created"`
	Done        bool      `json:"done"`
}

func (d DateJson) DateValidator() error {
	delay := time.Until(d.Date)
	if delay < 0 {
		return errors.New("Время не может быть в прошлом")
	}

	return nil
}
