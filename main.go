package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	First    string `json:"first_name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}

var users []User

var maxID uint64

// Esta funcion es para inicializar valores, en este caso nuestro arreglo de usuarios
func init() {
	users = []User{
		{
			ID:       1,
			First:    "John",
			LastName: "Doe",
			Email:    "mateo@gmail.com",
		},
		{
			ID:       2,
			First:    "John",
			LastName: "Doe",
			Email:    "mateo@gmail.com",
		},
		{
			ID:       3,
			First:    "John",
			LastName: "Doe",
			Email:    "mateo@gmail.com",
		},
	}
	maxID = 3
}

func main() {

	http.HandleFunc("/users", UserServer)
	/*
		http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		})
	*/
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func UserServer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetAllUsers(w)
	case http.MethodPost:
		decode := json.NewDecoder(r.Body)
		var user User
		err := decode.Decode(&user)

		if err != nil {
			MsgResponse(w, http.StatusBadRequest, err.Error())
			return

		}
		PostUser(w, user)
	default:
		fmt.Fprintf(w, "NOT FOUND, %q", html.EscapeString(r.URL.Path))

	}
}

// Esta funcion es para obtener todos los usuarios
func GetAllUsers(w http.ResponseWriter) {
	DataResponse(w, http.StatusOK, users)
}

func PostUser(w http.ResponseWriter, data interface{}) {
	//casteamos al data que viene como interface a User
	user := data.(User)

	maxID++
	user.ID = int(maxID)
	users = append(users, user)

	DataResponse(w, http.StatusCreated, user)

}
func MsgResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status" : %d, "msg": %s}`, status, message)
}

func DataResponse(w http.ResponseWriter, status int, users interface{}) {
	//Marshal me trata de convertir una entidad a un json
	value, err := json.Marshal(users)
	if err != nil {
		MsgResponse(w, http.StatusBadRequest, err.Error())
		return

	}
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status" : %d, "data": %s}`, status, value)
}
