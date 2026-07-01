package main

import (
	handels "micron/Handels"
	"net/http"
)

func main(){
	http.HandleFunc("/timer", handels.TimerHandler)
	http.HandleFunc("/date", handels.DateHandler)
	http.ListenAndServe(":6767", nil)
}

