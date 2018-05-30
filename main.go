package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// MinetestMessage : test message
type MinetestMessage struct {
	Message string
}

// PasswordToken : Password Token
type PasswordToken struct {
	Token string
}

// SpawnPoint : coordinates for spawn point.
type SpawnPoint struct {
	X int
	Y int
	Z int
}

// Boundary : coordinates for edges of player boundary
type Boundary struct {
	X1 int
	X2 int
	Z1 int
	Z2 int
}

// Group : player group
type Group struct {
	GroupName       string
	GroupBoundary   Boundary
	GroupSpawnPoint SpawnPoint
}

// Player : structure to represent each player
type Player struct {
	Name  string
	Group Group
}

func main() {
	fmt.Println("Start")

	// test data
	var spawnPoint SpawnPoint
	spawnPoint.X = 0
	spawnPoint.Y = 10
	spawnPoint.Z = 0

	var boundary Boundary
	boundary.X1 = -10
	boundary.X2 = 10
	boundary.Z1 = -10
	boundary.Z2 = 10

	var aGroup Group
	aGroup.GroupName = "Test Group 1"
	aGroup.GroupBoundary = boundary
	aGroup.GroupSpawnPoint = spawnPoint

	var aPlayer Player
	aPlayer.Name = "blah-1234"
	aPlayer.Group = aGroup
	//end test data

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
	r.Post("/player", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("API /player has been hit")
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
		fmt.Println(aPlayer)
		json.NewEncoder(w).Encode(aPlayer)
	})
	r.Post("/auth", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("API /auth has been hit")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Print(err.Error())
		}
		var t PasswordToken

		error := json.Unmarshal(body, &t)
		if error != nil {
			fmt.Print(error.Error())
		}
		fmt.Println(t.Token)
		if t.Token == "123456" {
			fmt.Println("authenticated")
			json.NewEncoder(w).Encode("authenticated")
		} else {
			fmt.Println("failed")
			json.NewEncoder(w).Encode("failed")
		}

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
