package controllers

import (

)

func Register(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")

	if (*r).Method == "POST":
		
}