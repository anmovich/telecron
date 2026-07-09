package datework

import (
	"time"
)

type DateJson struct {
	Date        time.Time `json:"time"`
	Description string    `json:"description"`
	Command     string    `json:"command"`
	Args        []string  `json:"args"`
	Repeat      bool      `json:"repeat"`
}
