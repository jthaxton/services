package handlers

import (
	"fmt"
	"net/http"
	// "fmt"
	"main/dblocal/schemas"
)

// IEventHandler is implement all the handlers
// type IEventHandler interface {
// 	GetOther(w http.ResponseWriter, r *http.Request)
// 	GetUser(w http.ResponseWriter, r *http.Request)
// }

func GetOther(w http.ResponseWriter, r *http.Request) {

	a := &schemas.Test{Args: "hithere"}
	b := MakeResponse(a) //[]byte{102, 97, 108, 99, 111, 110}
	// response := Response{Json(): data, StatusCode(): 200}
	// b, _ := json.Marshal(&a)
	// fmt.Println("HERE NOW")
	// fmt.Println(string(b))
	WriteResponse(w, b)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := schemas.GetUser(7)
	encoded_user := MakeResponse(user)
	WriteResponse(w,encoded_user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	succ := schemas.CreateUser("joe", "thaxton")
	fmt.Println("HIT ENDPOINT")
	if succ {
		user := schemas.GetUser(1)
		encoded_user := MakeResponse(user)
		WriteResponse(w,encoded_user)
	}

}