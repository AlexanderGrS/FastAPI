package handlers

import (
	"FastAPI/config"
	"FastAPI/helpers"
	"FastAPI/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func setupDB(cfg config.StorageConfig) *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cfg.Username, cfg.Password, cfg.Database)
	db, err := sql.Open("postgres", dbinfo)

	helpers.CheckErr(err)

	return db
}

func GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConfig()

	db := setupDB(cfg.Storage)

	helpers.PrintMessage("Getting recipes...")

	rows, err := db.Query(cfg.DBqueries.GetAllRecipes)

	helpers.CheckErr(err)

	var recipes []models.Recipe

	for rows.Next() {
		var id int
		var name string
		var description string
		var ingredients string
		var cooking_steps string
		var cooking_time string
		var recipe_rating float32

		err = rows.Scan(&id, &name, &description, &ingredients, &cooking_steps, &cooking_time, &recipe_rating)

		helpers.CheckErr(err)

		recipes = append(recipes, models.Recipe{Id: id, Name: name, Description: description, Ingredients: ingredients, Cooking_steps: cooking_steps, Cooking_time: cooking_time, Recipe_rating: recipe_rating})
	}

	var response = models.JsonResponse{Type: "success", Data: recipes, Message: "Got all recipes"}

	json.NewEncoder(w).Encode(response)
}

func GetRecipe(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConfig()

	db := setupDB(cfg.Storage)

	helpers.PrintMessage("Getting one recipe")

	ReqId := r.FormValue("id")

	var response = models.JsonResponse{}

	if ReqId == "" {
		response = models.JsonResponse{Type: "error", Message: "You need to insert id, id is null"}
		json.NewEncoder(w).Encode(response)
	}

	row := db.QueryRow(cfg.DBqueries.GetRecipe, ReqId)

	var id int
	var name string
	var description string
	var ingredients string
	var cooking_steps string
	var cooking_time string
	var recipe_rating float32

	err := row.Scan(&id, &name, &description, &ingredients, &cooking_steps, &cooking_time, &recipe_rating)

	helpers.CheckErr(err)
	recipe := models.Recipe{Id: id, Name: name, Description: description, Ingredients: ingredients, Cooking_steps: cooking_steps, Cooking_time: cooking_time, Recipe_rating: recipe_rating}

	json.NewEncoder(w).Encode(recipe)
}

func GetRecipesSortedByIngredients(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConfig()

	db := setupDB(cfg.Storage)

	helpers.PrintMessage("Sorted by Ingridients")

	rows, err := db.Query(cfg.DBqueries.GetRecipesSortedByIngredients)

	helpers.CheckErr(err)

	var recipes []models.Recipe

	for rows.Next() {
		var id int
		var name string
		var description string
		var ingredients string
		var cooking_steps string
		var cooking_time string
		var recipe_rating float32

		err = rows.Scan(&id, &name, &description, &ingredients, &cooking_steps, &cooking_time, &recipe_rating)

		helpers.CheckErr(err)

		recipes = append(recipes, models.Recipe{Id: id, Name: name, Description: description, Ingredients: ingredients, Cooking_steps: cooking_steps, Cooking_time: cooking_time, Recipe_rating: recipe_rating})
	}

	var response = models.JsonResponse{Type: "success", Data: recipes, Message: "recipes filtred by ingridients"}

	json.NewEncoder(w).Encode(response)

}

func GetRecipesSortedByCookingTime(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConfig()

	db := setupDB(cfg.Storage)

	helpers.PrintMessage("Sorted by Time")

	rows, err := db.Query(cfg.DBqueries.GetAllRecipes)

	helpers.CheckErr(err)

	var oldRecipes []models.Recipe

	for rows.Next() {
		var id int
		var name string
		var description string
		var ingredients string
		var cooking_steps string
		var cooking_time string
		var recipe_rating float32

		err = rows.Scan(&id, &name, &description, &ingredients, &cooking_steps, &cooking_time, &recipe_rating)

		helpers.CheckErr(err)

		oldRecipes = append(oldRecipes, models.Recipe{Id: id, Name: name, Description: description, Ingredients: ingredients, Cooking_steps: cooking_steps, Cooking_time: cooking_time, Recipe_rating: recipe_rating})
	}

	recipes := helpers.SortRecipesByTime(oldRecipes)

	var response = models.JsonResponse{Type: "success", Data: recipes, Message: "recipes filtred by cooking time"}

	json.NewEncoder(w).Encode(response)

}

func GetRecipesSortedByRating(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConfig()

	db := setupDB(cfg.Storage)

	helpers.PrintMessage("Sorted by Rating")

	rows, err := db.Query(cfg.DBqueries.GetAllRecipes)

	helpers.CheckErr(err)

	var oldRecipes []models.Recipe

	for rows.Next() {
		var id int
		var name string
		var description string
		var ingredients string
		var cooking_steps string
		var cooking_time string
		var recipe_rating float32

		err = rows.Scan(&id, &name, &description, &ingredients, &cooking_steps, &cooking_time, &recipe_rating)

		helpers.CheckErr(err)

		oldRecipes = append(oldRecipes, models.Recipe{Id: id, Name: name, Description: description, Ingredients: ingredients, Cooking_steps: cooking_steps, Cooking_time: cooking_time, Recipe_rating: recipe_rating})
	}

	recipes := helpers.SortRecipesByRating(oldRecipes)

	var response = models.JsonResponse{Type: "success", Data: recipes, Message: "recipes filtred by rating"}

	json.NewEncoder(w).Encode(response)

}
