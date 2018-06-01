package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

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
	Token string
}

// AuthResponse : stuff
type AuthResponse struct {
	Response   string
	PlayerData Player
}

func main() {
	fmt.Println("Start")

	// test data
	var spawnPoint SpawnPoint
	spawnPoint.X = 0
	spawnPoint.Y = 10
	spawnPoint.Z = 0

	var spawnPoint2 SpawnPoint
	spawnPoint2.X = 5
	spawnPoint2.Y = 10
	spawnPoint2.Z = 22

	var boundary Boundary
	boundary.X1 = -10
	boundary.X2 = 10
	boundary.Z1 = -10
	boundary.Z2 = 10

	var boundary2 Boundary
	boundary2.X1 = -10
	boundary2.X2 = 10
	boundary2.Z1 = 12
	boundary2.Z2 = 32

	var aGroup Group
	aGroup.GroupName = "Test Group 1"
	aGroup.GroupBoundary = boundary
	aGroup.GroupSpawnPoint = spawnPoint

	var aGroup2 Group
	aGroup2.GroupName = "Test Group 2"
	aGroup2.GroupBoundary = boundary2
	aGroup2.GroupSpawnPoint = spawnPoint2

	var aPlayer Player
	aPlayer.Name = "blah-1234"
	aPlayer.Group = aGroup
	aPlayer.Token = "123456"

	var aPlayer2 Player
	aPlayer2.Name = "blah-1234"
	aPlayer2.Group = aGroup2
	aPlayer2.Token = "123123"
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
	r.Post("/auth", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("API /auth has been hit")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Print(err.Error())
		}
		var t PasswordToken
		var response AuthResponse
		error := json.Unmarshal(body, &t)
		if error != nil {
			fmt.Print(error.Error())
		}
		if t.Token == "123456" {
			response.Response = "authenticated"
			response.PlayerData = aPlayer
			json.NewEncoder(w).Encode(response)
		} else if t.Token == "123123" {
			response.Response = "authenticated"
			response.PlayerData = aPlayer2
			json.NewEncoder(w).Encode(response)
		} else {
			response.Response = "failed"
			json.NewEncoder(w).Encode(response)
		}

	})

	http.ListenAndServe(":3001", r)
}
