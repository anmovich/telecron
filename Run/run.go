package run

import (
	datework "micron/DateWork"
	"os/exec"
	"time"
)

// func to exec commands by date
func ExecCommand(d datework.DateJson) error {
	delay := time.Until(d.Date)
	time.Sleep(delay)
	cmd := exec.Command(d.Command, d.Args...)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
