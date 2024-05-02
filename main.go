package main

import (
	"fmt"
	"forum/controllers"
	"forum/initializers"
	"log"
	"net/http"
)

func main() {
	_, err := initializers.InitDB()
	if err != nil {
		panic(err)
	}
	// Définir les gestionnaire
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	// Définir les autres routes
	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/form", controllers.FormHandler)
	http.HandleFunc("/register", controllers.RegisterUserHandler)
	http.HandleFunc("/connexion", controllers.ConnexionUserHandler)
	http.HandleFunc("/forum/acceuil", controllers.ForumAcceuilHandler)
	// Démarrer le serveur HTTP sur le port 8080
	fmt.Println("Server running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
