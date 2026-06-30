package notifications

import (
	"micron/Timer"
	"os/exec"
	"time"
)

func SendNotificationTimer(t *timer.Time) error{
	var err error
	t.Duration, err = timer.TimeConverter(t.UnformatDuration, t.TimeType)
	if err != nil{
		return err
	}
	time.Sleep(t.Duration)
	cmd := exec.Command("notify-send",
							 "Уведомление от программы",
							 t.Text,
  )
	err = cmd.Run()
return err
}
