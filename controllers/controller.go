package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/models"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("cAR35t342MvhTt"))

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

	session, err := store.Get(r, "digiforum")
	if err != nil {
		http.Error(w, "Erreur de session", http.StatusInternalServerError)
		return
	}

	// Vérifier si l'utilisateur est authentifié
	userID, ok := session.Values["username"].(string)
	if !ok {
		// L'utilisateur n'est pas authentifié, rediriger vers la page de connexion
		http.Redirect(w, r, "/form", http.StatusFound)
		return
	}
	fmt.Println(userID)

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

// Connexion
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

		// Créer une session pour l'utilisateur

		session, err := store.Get(r, "digiforum")
		if err != nil {
			http.Error(w, "Erreur de session", http.StatusInternalServerError)
			return
		}

		session.Values["username"] = user.Username
		// Définir d'autres valeurs de session si nécessaire

		// Définir l'expiration de la session
		session.Options = &sessions.Options{
			MaxAge:   1800, // durée de vie de la session en secondes (ici,30 min)
			HttpOnly: true,
			Secure:   true,
		}

		// Enregistrer la session
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, "Erreur de session", http.StatusInternalServerError)
			return
		}
		// Authentification réussie
		// Vous pouvez renvoyer une réponse avec un code de statut 200 pour indiquer une connexion réussie
		// ou renvoyer des données supplémentaires sur l'utilisateur connecté si nécessaire
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Connexion réussie"))

	}
}

// deconnexion
func DeconnexionHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer la session de l'utilisateur
	session, err := store.Get(r, "digiforum")
	if err != nil {
		http.Error(w, "Erreur de session", http.StatusInternalServerError)
		return
	}

	// Supprimer le cookie de session en fixant la durée de vie à une valeur négative
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Erreur de session", http.StatusInternalServerError)
		return
	}

	// Rediriger l'utilisateur vers la page d'accueil ou une autre page
	http.Redirect(w, r, "/", http.StatusFound)
}

// reset_password
/*func reset_passwordHandler(w http.ResponseWriter, r *http.Request) {
	// Vérifier la méthode HTTP
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	// Récupérer les données du formulaire
	ancienMotDePasse := r.FormValue("ancien_mot_de_passe")
	nouveauMotDePasse := r.FormValue("nouveau_mot_de_passe")
	// Vérifier l'authentification de l'utilisateur
	// (vous pouvez utiliser la session ou toute autre méthode d'authentification que vous avez mise en place)

	// Vérifier si l'ancien mot de passe correspond au mot de passe actuel de l'utilisateur
	// (vous devrez adapter cette vérification à votre modèle de données)
	if ancienMotDePasse != motDePasseActuel {
		http.Error(w, "Ancien mot de passe incorrect", http.StatusBadRequest)
		return
	}

	// Mettre à jour le mot de passe de l'utilisateur avec le nouveau mot de passe
	// (vous devrez adapter cette mise à jour à votre modèle de données)
	err := MettreAJourMotDePasseUtilisateur(userID, nouveauMotDePasse)
	if err != nil {
		http.Error(w, "Erreur lors de la mise à jour du mot de passe", http.StatusInternalServerError)
		return
	}

	// Rediriger l'utilisateur vers une page de confirmation ou une autre page appropriée
	http.Redirect(w, r, "/profil", http.StatusFound)
}*/

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
