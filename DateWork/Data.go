package datework

import (
	"fmt"
	"os/exec"
	"time"
)

func GetDate(year *int, month *int, day *int, hour *int, minute *int) time.Time{
	fmt.Println("Введите год через пробел: год месяц день:")
	fmt.Scanf("%d %d %d", year, month, day)
	fmt.Println("Введите час минуты:")
	fmt.Scanf("%d %d", hour, minute)
	return CreateDate(*year, *month, *day, *hour, *minute)
} 

func CreateDate(year int, month int, day int, hour int, minute int)time.Time{
	return time.Date(
		year,
		monthConverter(month),
		day,
		hour,
		minute,
		0,
		0,
		time.Local,
	)
}

func NotificationOnDate(t time.Time, text string) error{
	delay := time.Until(t)
	time.Sleep(delay)
	cmd := exec.Command("notify-send", 
											"Уведомление по дате",
											text)
	err := cmd.Run()
	return err
}

func DateMenu(){
	var text string
	fmt.Print("Введите текст уведомления: ")
	fmt.Scan(&text)
	var year, month, day, hour, minute int
	t := GetDate(&year, &month, &day, &hour, &minute)
	go func(t time.Time, text string){
		if err := NotificationOnDate(t, text); err != nil{
			fmt.Println(err)
		}
	}(t, text)
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
