// modele/user.go
package models

import (
	"fmt"
	"forum/initializers"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Nom      string `json:"nom"`
	Prenom   string `json:"prenom"`
	Photo    string `json:"photo"`
}

// requette pour enregistrer un utilisateur
func InsertUser(user User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	fmt.Println("oooooooooooooooooooooooooooo")
	db, errs := initializers.GetDb()
	if errs != nil {
		return errs
	}
	// Préparer la requête d'insertion
	stmt, err := db.Prepare(`
			INSERT INTO utilisateur (nom, prenom, photo, username, mot_de_passe, email) 
			VALUES (?, ?, ?, ?, ?, ?)
    `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Exécuter la requête d'insertion avec les valeurs de l'utilisateur
	_, err = stmt.Exec(user.Nom, user.Prenom, user.Photo, user.Username, hashedPassword, user.Email)
	if err != nil {
		return err
	}

	fmt.Println("Utilisateur inséré avec succès")
	return nil
}

// requette pour recuperer tous les utilisateurs
func FindAllUser(user User, id int) error {
	// Préparer la requête d'insertion
	return nil
}

// requette pour recuperer un seul user par ID
func FindOneUser(user User, id int) error {
	// Préparer la requête d'insertion
	return nil
}

// requette pour modifier un seul user par ID
func UpdateUser(user User, id int) error {
	// Préparer la requête d'insertion
	return nil
}

// requette pour modifier un seul user par ID
func DeleteUser(user User, id int) error {
	// Préparer la requête d'insertion
	return nil
}
