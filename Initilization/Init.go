package initilization

import (
	"context"
	"fmt"
	bd "micron/Bd"
	datework "micron/DateWork"
	run "micron/Run"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Init struct {
	BD  *pgxpool.Pool
	Ctx context.Context
}

func (ini *Init) CheckDbGo() error {
	shit, err := bd.CheckDbNotDone(ini.Ctx, ini.BD)
	if err != nil {
		return err
	}
	e := run.Execute{DB: ini.BD, Ctx: ini.Ctx}
	for _, v := range shit {
		if v.Repeat == false {
			go func(v datework.DateJson) {
				if err := e.ExecOnce(v); err != nil {
					fmt.Println(err)
				}
			}(v)
		} else {
			go func(v datework.DateJson) {
				if err := e.ExecRepeat(&v); err != nil {
					fmt.Println(err)
				}
			}(v)
		}
	}
	return nil
}
