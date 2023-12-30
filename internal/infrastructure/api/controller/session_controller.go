package infrastructure

import (
	"encoding/json"
	"net/http"
	"strconv"

	database "github.com/beriloqueiroz/prontu-go/internal/infrastructure/persistence"
	models "github.com/beriloqueiroz/prontu-go/internal/infrastructure/persistence/models"
	"github.com/gorilla/mux"
)

func GetSessions(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")
	pageInt, errP := strconv.Atoi(page)
	sizeInt, errS := strconv.Atoi(size)

	if errP != nil || errS != nil {
		pageInt = 0
		sizeInt = 100
	}

	var sessions []models.Session
	database.DB.Limit(sizeInt).Offset(pageInt).Find(&sessions)
	json.NewEncoder(w).Encode(sessions)
}

func GetSessionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var session models.Session
	database.DB.First(&session, id)
	json.NewEncoder(w).Encode(session)
}

func CreateSession(w http.ResponseWriter, r *http.Request) {
	var newSession models.Session
	json.NewDecoder(r.Body).Decode(&newSession)
	valid, msg := newSession.Validate()
	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(DefaultError{Title: "400 - Something bad happened! - " + msg})
		return
	}
	database.DB.Create(&newSession)
	json.NewEncoder(w).Encode(newSession)
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var session models.Session
	database.DB.Delete(&session, id)
	json.NewEncoder(w).Encode(session)
}

func UpdateSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var session models.Session
	database.DB.First(&session, id)
	json.NewDecoder(r.Body).Decode(&session)
	valid, msg := session.Validate()
	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(DefaultError{Title: "400 - Something bad happened! - " + msg})
		return
	}
	database.DB.Save(&session)
	json.NewEncoder(w).Encode(&session)
}
