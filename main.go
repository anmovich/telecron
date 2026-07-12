package main

import (
	"context"
	bd "micron/Bd"
	handels "micron/Handels"
	initilization "micron/Initilization"
	"net/http"
)

func main() {
	ctx := context.Background()
	pool, err := bd.CreateConnection(ctx)

	if err != nil {
		panic(err)
	}
	defer pool.Close()
	h := handels.Handler{
		DB:  pool,
		Ctx: ctx,
	}
	ini := initilization.Init{BD: pool, Ctx: ctx}
	ini.CheckDbGo()
	if err := bd.CreateDatabase(ctx, pool); err != nil {
		panic(err)
	}
	http.HandleFunc("/timer", h.TimerHandler)
	http.HandleFunc("/date", h.DateHandler)
	http.ListenAndServe(":6767", nil)
}
