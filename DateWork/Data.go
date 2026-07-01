package datework

import (
	"errors"
	"time"
)

type DateJson struct{
	Year int `json:"year"`
	Month int `json:"month"`
	Day int `json:"day"`
	Hour int `json:"hour"`
	Minute int `json:"minute"`
	Text string `json:"text"`
} 

func RemakeDate(d DateJson)time.Time{
	return time.Date(
		d.Year,
		monthConverter(d.Month),
		d.Day,
		d.Hour,
		d.Minute,
		0,
		0,
		time.Local,
	)
}

func DateValidator(d DateJson) error{
	if d.Year < 0{
		return errors.New("Incorrect year")
	} else if d.Month <0 || d.Month > 12{
		return errors.New("Incorrect month")
	} else if d.Day < 0 || d.Day > 31{
		return errors.New("Incorrect day")
	} else if d.Hour < 0 || d.Hour > 24{
		return errors.New("Incorrect hour")
	} else if d.Minute < 0 || d.Minute > 60{
		return errors.New("Incorrect minute")
	} else{
		return nil
	} 
}

func monthConverter(month int) time.Month{
	m := []time.Month{
		time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August,
		time.September,
		time.October,
		time.November,
		time.December,
	}
	return m[month -1]
}
