package controllers

import (
	"encoding/json"
	"forum/models"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("views/index.html"))
	if err := tpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("views/form.html"))
	if err := tpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Inscription de l'utilisateur
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Lire les données JSON de la requête
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Erreur lors de la lecture des données JSON", http.StatusBadRequest)
			return
		}

		// Insérer l'utilisateur dans la base de données
		err = models.InsertUser(user)
		if err != nil {
			http.Error(w, "Erreur lors de l'insertion de l'utilisateur dans la base de données", http.StatusInternalServerError)
			return
		}

		// Créer un objet réponse contenant des données JSON
		response := struct {
			Message string `json:"message"`
		}{
			Message: "Inscription réussie pour l'utilisateur: " + user.Username,
		}

		// Convertir la réponse en JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Erreur lors de la conversion de la réponse en JSON", http.StatusInternalServerError)
			return
		}

		// Définir le type de contenu de la réponse comme JSON
		w.Header().Set("Content-Type", "application/json")
		// Répondre avec un code de statut 200 et les données JSON
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}
