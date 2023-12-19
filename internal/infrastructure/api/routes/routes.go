package infrastructure

import (
	"log"
	"net/http"

	controllers "github.com/beriloqueiroz/prontu-go/internal/infrastructure/api/controller"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/api/sessions", controllers.GetSessions).Methods("GET")
	r.HandleFunc("/api/sessions/{id}", controllers.GetSessionById).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
