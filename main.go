package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type MinetestMessage struct {
	Message string
}

func main() {
	fmt.Println("Start")

	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("API / has been hit")
		w.Write([]byte("welcome"))
	})
	r.Post("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("API /test has been hit")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Print(err.Error())
		}
		var m MinetestMessage

		error := json.Unmarshal(body, &m)
		if error != nil {
			fmt.Print(error.Error())
		}
		fmt.Println(m.Message)
		json.NewEncoder(w).Encode(testFunc(r))
	})

	http.ListenAndServe(":3001", r)
}

// unmarshalToPerson : takes r http.Request and returns Person object
func testFunc(r *http.Request) string {
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }
	return "message from go API"
}
