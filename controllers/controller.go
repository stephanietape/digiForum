package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/models"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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

func ForumAcceuilHandler(w http.ResponseWriter, r *http.Request) {

	tpl := template.Must(template.ParseFiles("views/acceuil.html"))
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
		var email = user.Email
		// Vérifier si l'email ou le nom d'utilisateur est déjà présent dans la base de données
		if userExists(email, "email") || userExists(user.Username, "username") {
			fmt.Println("Email ou nom d'utilisateur déjà utilisé")
			http.Error(w, "Email ou nom d'utilisateur déjà utilisé", http.StatusBadRequest)
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

//Connexion

func ConnexionUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("connexionnnnn")
	if r.Method == "POST" {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Erreur lors de la lecture des données JSON", http.StatusBadRequest)
			return
		}

		// Vérifier si l'utilisateur existe dans la base de données
		db, err := sql.Open("sqlite3", "forum.db")
		if err != nil {
			http.Error(w, "Erreur de connexion à la base de données", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var dbPassword string
		err = db.QueryRow("SELECT mot_de_passe FROM utilisateur WHERE  username = ?", user.Username).Scan(&dbPassword)
		if err != nil {
			fmt.Println("Utilisateur non trouvé")
			http.Error(w, "Utilisateur non trouvé", http.StatusUnauthorized)
			return
		}
		fmt.Println("dbPassword", dbPassword)

		// Comparer les mots de passe
		// Comparer le mot de passe fourni avec le hash stocké dans la base de données
		err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(user.Password))
		if err != nil {
			fmt.Println("Mot de passe incorrect")
			http.Error(w, "Mot de passe incorrect", http.StatusUnauthorized)
			return
		}

		// Authentification réussie
		// Vous pouvez renvoyer une réponse avec un code de statut 200 pour indiquer une connexion réussie
		// ou renvoyer des données supplémentaires sur l'utilisateur connecté si nécessaire
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Connexion réussie"))

	}
}

// fonction pour verifier si un utilisateur exist dans la base de donnée
func userExists(value string, column string) bool {
	// Ouvrir la connexion à la base de données
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		// Gérer l'erreur de connexion à la base de données
		return true // Supposons que l'utilisateur existe pour éviter de créer une nouvelle inscription
	}
	defer db.Close()

	// Préparer la requête SQL pour vérifier si la valeur existe déjà
	query := "SELECT COUNT(*) FROM utilisateur WHERE " + column + " = ?"
	var count int
	err = db.QueryRow(query, value).Scan(&count)
	if err != nil {
		// Gérer l'erreur de requête SQL
		return true // Supposons que l'utilisateur existe pour éviter de créer une nouvelle inscription
	}

	return count > 0 // Si count > 0, cela signifie que l'utilisateur existe déjà, sinon non
}
