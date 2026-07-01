package timer

import (
	"errors"
	"time"
)

type Time struct{
	TimeType int `json:"type"`
	UnformatDuration int `json:"duration"`
	Duration time.Duration `json:"formatDuration"`
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

//Validating timer struct
func (t Time) ValidatorTimer() (bool, error){
	if t.TimeType < 1 || t.TimeType > 3{
		return false, errors.New("Incorrect time type")
	}
	if t.UnformatDuration < 0{
		return false, errors.New("Incorrect duration")
	}
	return true, nil
}
