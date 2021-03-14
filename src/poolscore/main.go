package main

import (
	"net/http"
	h "poolscore/handlers"
	. "poolscore/utils"
)

var c Config

func main() {
	c.LoadConfig()
	http.HandleFunc("/", h.ShowScores)
	http.ListenAndServe(":8080", nil)
}
