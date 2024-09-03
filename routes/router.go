package routes

import (
	"api-go-rest/controllers"
	"api-go-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest()  {
	r := mux.NewRouter()
	r.Use(middleware.AddContentType)

	r.HandleFunc("/", controllers.Home).Methods("GET")

	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonalityById).Methods("GET")

	r.HandleFunc("/api/personalities", controllers.InsertPersonality).Methods("POST")
	r.HandleFunc("/api/personalities/{id}", controllers.EditPersonality).Methods("PUT")
	r.HandleFunc("/api/personalities/{id}", controllers.RemovePersonality).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
