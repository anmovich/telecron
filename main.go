package main

import (
	handels "micron/Handels"
	"net/http"
)

func main(){
	http.HandleFunc("/timer", handels.TimerHandler)
	http.ListenAndServe(":6767", nil)
}

