package timer

import (

	"time"
)

type Time struct{
	TimeType int `json:"type"`
	UnformatDuration int `json:"duration"`
	Duration time.Duration
	Text string `json:"text"`
}


func TimeConverter(tim int, typ int) (time.Duration){
	hui := []time.Duration{
		time.Second,
		time.Minute,
		time.Hour,
	}	
	return time.Duration(tim) * hui[typ -1]
}


