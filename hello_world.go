package main

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
)

type User struct {
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Email 		string `json:"email"`
	CreatedAt 	time.Time `json:"created_at"`
}

type HomePageHandler struct{}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	user := new(User)
	json.NewDecoder(r.Body).Decode(user)
	user.CreatedAt = time.Now()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user)
	w.Write(data)
}

func main(){
	http.Handle("/", &HomePageHandler{})

	barHandler := func (w http.ResponseWriter, r *http.Request){
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Bar"
		}
		fmt.Fprintf(w, "Hello %s!", name )
	}
	http.HandleFunc("/bar", barHandler)

	http.ListenAndServe(":3000", nil)
}

