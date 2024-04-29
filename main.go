package main

import (
	"fmt"
	"forum/controllers"
	"log"
	"net/http"
)

func main() {
	// Définir les gestionnaire
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	// Définir les autres routes
	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/form", controllers.FormHandler)
	// Démarrer le serveur HTTP sur le port 8080
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
