package handlers

import (
	"encoding/json"
	config "main/configs"
	"net/http"

	"main/internal/entities"
)

func GetDogs(w http.ResponseWriter, r *http.Request) {
	var dogs []entities.Dog

	config.Database.Find(&dogs)

	json.NewEncoder(w).Encode(dogs)
}

func GetDog(w http.ResponseWriter, r *http.Request) {
	var dog entities.Dog

	id := r.URL.Query().Get("id")
	config.Database.First(&dog, id)

	json.NewEncoder(w).Encode(dog)
}

func CreateDog(w http.ResponseWriter, r *http.Request) {
	var dog entities.Dog

	if err := json.NewDecoder(r.Body).Decode(&dog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	config.Database.Create(&dog)

	json.NewEncoder(w).Encode(dog)
}

func UpdateDog(w http.ResponseWriter, r *http.Request) {
	var dog entities.Dog

	if err := json.NewDecoder(r.Body).Decode(&dog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := r.URL.Query().Get("id")
	config.Database.Where("id = ?", id).Updates(&dog)

	json.NewEncoder(w).Encode(dog)
}

func DeleteDog(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	config.Database.Delete(&entities.Dog{}, id)

	json.NewEncoder(w).Encode("Dog deleted")
}
