package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	database "github.com/beriloqueiroz/prontu-go/internal/infrastructure/persistence"
	models "github.com/beriloqueiroz/prontu-go/internal/infrastructure/persistence/models"
	"github.com/gorilla/mux"
)

func GetSessions(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")
	fmt.Printf("page: %v\n", page)
	fmt.Printf("size: %v\n", size)
	pageInt, errP := strconv.Atoi(page)
	sizeInt, errS := strconv.Atoi(size)

	if errP != nil || errS != nil {
		pageInt = 0
		sizeInt = 100
	}

	var sessions []models.Session
	database.DB.Limit(sizeInt).Offset(pageInt).Find(&sessions)
	fmt.Println(sessions)
	json.NewEncoder(w).Encode(sessions)
}

func GetSessionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var session models.Session
	database.DB.First(&session, id)
	json.NewEncoder(w).Encode(session)
}
