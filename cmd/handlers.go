package main

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := RenderPage(w, r, "home.jet", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
