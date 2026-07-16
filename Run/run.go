package run

import (
	"context"
	bd "micron/Bd"
	datework "micron/DateWork"
	"os/exec"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Execute struct {
	Ctx context.Context
	DB  *pgxpool.Pool
}

// func to exec commands by date
func (e *Execute) ExecCommand(d datework.DateJson) error {
	delay := time.Until(d.Date)
	time.Sleep(delay)
	cmd := exec.Command(d.Command, d.Args...)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (e *Execute) ExecOnce(d datework.DateJson) error {
	if err := e.ExecCommand(d); err != nil {
		return err
	}
	if err := bd.UpdateTaskDone(d, e.Ctx, e.DB); err != nil {
		return err
	}
	return nil
}

func (e *Execute) ExecRepeat(d *datework.DateJson) error {
	for {
		if err := e.ExecCommand(*d); err != nil {
			return err
		}
		d.Date = d.Date.AddDate(0, 0, 1)
		bd.UpdateTask(*d, e.Ctx, e.DB)
	}

}
