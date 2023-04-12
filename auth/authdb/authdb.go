package authdb

import (
	"FastAPI/config"
	"FastAPI/helpers"
	"FastAPI/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func setupDB(cfg config.StorageConfig) *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cfg.Username, cfg.Password, cfg.Database)
	db, err := sql.Open("postgres", dbinfo)

	helpers.CheckErr(err)

	return db
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	db_user := r.FormValue("db_user")
	password := r.FormValue("password")

	cfg := config.GetConfig()
	db := setupDB(cfg.Storage)

	helpers.PrintMessage("Creating New db_user")

	encryptedPassword, err := hashPassword(password)
	helpers.CheckErr(err)

	db.QueryRow(cfg.DBqueries.SignUp, db_user, encryptedPassword)

	var response = models.JsonResponse{Type: "success", Message: "New db_user created"}

	json.NewEncoder(w).Encode(response)
}

func VerifyUserPass(username, password string) bool {

	cfg := config.GetConfig()
	db := setupDB(cfg.Storage)

	row := db.QueryRow(cfg.DBqueries.SignIn, username)
	if row == nil {
		return false
	}

	var db_password string
	err := row.Scan(&username, &db_password)
	helpers.CheckErr(err)

	if compare := bcrypt.CompareHashAndPassword([]byte(db_password), []byte(password)); compare == nil {
		return true
	}
	return false
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
