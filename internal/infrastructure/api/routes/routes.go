package infrastructure

import (
	"log"
	"net/http"

	controllers "github.com/beriloqueiroz/prontu-go/internal/infrastructure/api/controller"
	middleware "github.com/beriloqueiroz/prontu-go/internal/infrastructure/api/middleware"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/api/sessions", controllers.GetSessions).Methods("GET")
	r.HandleFunc("/api/sessions/{id}", controllers.GetSessionById).Methods("GET")
	r.HandleFunc("/api/sessions/{id}", controllers.DeleteSession).Methods("DELETE")
	r.HandleFunc("/api/sessions", controllers.CreateSession).Methods("POST")
	r.HandleFunc("/api/sessions", controllers.UpdateSession).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
