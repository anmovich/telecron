package main

import (
	"context"
	bd "micron/Bd"
	handels "micron/Handels"
	"net/http"
)

func main() {
	ctx := context.Background()
	conn, err := bd.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	if err := bd.CreateDatabase(ctx, conn); err != nil {
		panic(err)
	}
	http.HandleFunc("/timer", handels.TimerHandler)
	http.HandleFunc("/date", handels.DateHandler)
	http.ListenAndServe(":6767", nil)
}
