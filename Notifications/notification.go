package notifications

import (
	"micron/Timer"
	"os/exec"
	"time"
)

func SendNotificationTimer(t *timer.Time) error{
	t.Duration = timer.TimeConverter(t.UnformatDuration, t.TimeType)
	time.Sleep(t.Duration)
	cmd := exec.Command("notify-send",
							 "Уведомление от программы",
							 t.Text,
  )
	err := cmd.Run()
return err
}
