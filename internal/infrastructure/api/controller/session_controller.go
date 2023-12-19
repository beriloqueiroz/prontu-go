package infrastructure

import (
	"encoding/json"
	"net/http"

	models "github.com/beriloqueiroz/prontu-go/internal/infrastructure/persistence/models"
	"github.com/gorilla/mux"
)

func GetSessions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Sessions)
}

func GetSessionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, v := range models.Sessions {
		if v.Id == id {
			json.NewEncoder(w).Encode(v)
		}
	}
}
