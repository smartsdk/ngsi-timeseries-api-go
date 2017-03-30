package 

import (
	"net/http"
)

type APIEntryPoint struct {

}

func Retrieve API Resources(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

