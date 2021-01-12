package main

import (
	"log"
	"net/http"

	"github.com/challenge/pkg/auth"
	"github.com/challenge/pkg/controller"
	"github.com/challenge/pkg/database"
	"github.com/challenge/pkg/jwt"
	"github.com/challenge/pkg/service"

	_ "github.com/mattn/go-sqlite3"
)

const (
	ServerPort       = "8080"
	CheckEndpoint    = "/check"
	UsersEndpoint    = "/users"
	LoginEndpoint    = "/login"
	MessagesEndpoint = "/messages"
)

var (
	handler      controller.Handler
	validateUser func(http.HandlerFunc) http.HandlerFunc
)

func main() {
	conn := database.NewConnection("chat.sqlite")
	defer conn.Close()
	database.InitChatDatabase()

	handler = controller.Handler{
		Service: service.NewService(database.NewSQLiteDB(conn), jwt.New()),
	}
	validateUser = auth.NewValidateUserHandler(jwt.New())

	configureEndpoints()

	// Start server
	log.Println("Server started at port " + ServerPort)
	log.Fatal(http.ListenAndServe(":"+ServerPort, nil))
}

func configureEndpoints() {
	// Health
	http.HandleFunc(CheckEndpoint, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		handler.Check(w, r)
	})

	// Users
	http.HandleFunc(UsersEndpoint, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		handler.CreateUser(w, r)
	})

	// Auth
	http.HandleFunc(LoginEndpoint, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		handler.Login(w, r)
	})

	// Messages
	http.HandleFunc(MessagesEndpoint, validateUser(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetMessages(w, r)
		case http.MethodPost:
			handler.SendMessage(w, r)
		default:
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
	}))

}
