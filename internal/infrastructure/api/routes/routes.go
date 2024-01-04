package infrastructure

import (
	"log"
	"net/http"

	controllers "github.com/beriloqueiroz/prontu-go/internal/infrastructure/api/controller"
	middleware "github.com/beriloqueiroz/prontu-go/internal/infrastructure/api/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var allowOrigins = []string{"*"} //todo specify

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/api/sessions", controllers.GetSessions).Methods("GET")
	r.HandleFunc("/api/sessions/professional/{id}", controllers.GetSessionsByProfessional).Methods("GET")
	r.HandleFunc("/api/sessions/professional/{professionalId}/{patientId}", controllers.GetSessionsByProfessionalAndPatient).Methods("GET")
	r.HandleFunc("/api/sessions/{id}", controllers.GetSessionById).Methods("GET")
	r.HandleFunc("/api/sessions/{id}", controllers.DeleteSession).Methods("DELETE")
	r.HandleFunc("/api/sessions/professional/{professionalId}/{id}", controllers.DeleteSessionByProfessional).Methods("DELETE")
	r.HandleFunc("/api/sessions", controllers.CreateSession).Methods("POST")
	r.HandleFunc("/api/sessions", controllers.UpdateSession).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins(allowOrigins))(r)))
}
