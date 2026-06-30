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


func TimeConverter(tim int, typ int) (time.Duration, error){
	hui := []time.Duration{
		time.Second,
		time.Minute,
		time.Hour,
	}	
	return time.Duration(tim) * hui[typ -1], nil
}

func CheckType(t Time) bool{
	if t.TimeType < 1 || t.TimeType > 3{
		return false
	}
	return true
}
