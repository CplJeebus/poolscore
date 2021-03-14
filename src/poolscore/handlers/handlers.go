package handlers

import (
	"fmt"
	"net/http"
	. "poolscore/utils"
)

var c Config

func ShowScores(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, Pretty)
}
