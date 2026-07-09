package main

import (
	"context"
	bd "micron/Bd"
	handels "micron/Handels"
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
	if err := bd.CreateDatabase(ctx, pool); err != nil {
		panic(err)
	}
	http.HandleFunc("/timer", h.TimerHandler)
	http.HandleFunc("/date", h.DateHandler)
	http.ListenAndServe(":6767", nil)
}
