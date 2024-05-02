package initializers

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	// Ouvrir la connexion à la base de données
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		return nil, err
	}

	// Créer la table pour stocker les informations des utilisateurs
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS utilisateur (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				nom TEXT NOT NULL,
				prenom TEXT NOT NULL,
				photo BLOB,
				username TEXT NOT NULL,
				mot_de_passe TEXT NOT NULL,
				email TEXT NOT NULL
			);
			CREATE TABLE IF NOT EXISTS Token_utilisateur (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				token TEXT NOT NULL,
				date_expiration TEXT NOT NULL,
				id_utilisateur INTEGER NOT NULL,
				FOREIGN KEY (id_utilisateur) REFERENCES utilisateur (id)
			);
			CREATE TABLE IF NOT EXISTS Session (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				id_utilisateur INTEGER NOT NULL,
				date_connexion TEXT NOT NULL,
				actif INTEGER DEFAULT 1,
				FOREIGN KEY (id_utilisateur) REFERENCES utilisateur (id)
			);
    `)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetDb() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return nil, err
	}
	return
}
