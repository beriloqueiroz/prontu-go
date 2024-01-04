package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"

	database "github.com/beriloqueiroz/prontu-go/internal/infrastructure/persistence"
	models "github.com/beriloqueiroz/prontu-go/internal/infrastructure/persistence/models"
	"github.com/gorilla/mux"

	"github.com/google/uuid"
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
	database.DB.Limit(sizeInt).Offset(pageInt).Preload("Patients").Preload("Professionals").Find(&sessions)
	json.NewEncoder(w).Encode(sessions)
}

func GetSessionsByProfessional(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")
	pageInt, errP := strconv.Atoi(page)
	sizeInt, errS := strconv.Atoi(size)

	if errP != nil || errS != nil {
		pageInt = 0
		sizeInt = 100
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var sessions []models.Session
	database.DB.Limit(sizeInt).Preload("Patients").Offset(pageInt).Preload("Professionals").Joins("JOIN professionals p ON p.session_id = sessions.id AND p.professional_id = ?", id).Find(&sessions)
	json.NewEncoder(w).Encode(sessions)
}

func GetSessionsByProfessionalAndPatient(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")
	pageInt, errP := strconv.Atoi(page)
	sizeInt, errS := strconv.Atoi(size)

	if errP != nil || errS != nil {
		pageInt = 0
		sizeInt = 100
	}

	vars := mux.Vars(r)
	professionalId := vars["professionalId"]
	patientId := vars["patientId"]

	var sessions []models.Session
	database.DB.Limit(sizeInt).Offset(pageInt).Preload("Patients").Preload("Professionals").Joins("JOIN professionals p ON p.session_id = sessions.id").Joins("JOIN patients pat ON pat.session_id = sessions.id").Where("p.professional_id=? AND pat.patient_id=?", professionalId, patientId).Find(&sessions)
	json.NewEncoder(w).Encode(sessions)
}

func GetSessionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var session models.Session
	database.DB.Preload("Patients").Preload("Professionals").First(&session, "id = ?", id)
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
	newSession.Id = uuid.NewString()
	database.DB.Create(&newSession)
	json.NewEncoder(w).Encode(newSession)
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var session models.Session
	database.DB.Delete(&session, "id = ?", id)
	json.NewEncoder(w).Encode(session)
}

func DeleteSessionByProfessional(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	professionalId := vars["professionalId"]
	var session models.Session
	database.DB.Preload("Professionals").First(&session, "id = ?", id)

	fmt.Println(session)

	if !slices.ContainsFunc(session.Professionals, func(p models.Professional) bool {
		return p.ProfessionalId == professionalId
	}) {
		json.NewEncoder(w).Encode(DefaultError{Title: "400 - Something bad happened! - professional inválido"})
		return
	}

	database.DB.Delete(&session, "id = ?", id)
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
