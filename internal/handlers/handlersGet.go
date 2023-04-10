package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"testApi/config"

	_ "github.com/lib/pq"
)

type Recipe struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Ingredients   string  `json:"ingredients"`
	Cooking_steps string  `json:"cooking_steps"`
	Cooking_time  string  `json:"cooking_time"`
	Recipe_rating float32 `json:"recipe_rating"`
}

type JsonResponse struct {
	Type    string   `json:"type"`
	Data    []Recipe `json:"data"`
	Message string   `json:"message"`
}

func setupDB(cfg config.StorageConfig) *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cfg.Username, cfg.Password, cfg.Database)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

func GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConfig()

	db := setupDB(cfg.Storage)

	printMessage("Getting recipes...")

	rows, err := db.Query(cfg.DBqueries.GetAllRecipes)

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var recipes []Recipe

	// Foreach movie
	for rows.Next() {
		var id int
		var name string
		var description string
		var ingredients string
		var cooking_steps string
		var cooking_time string
		var recipe_rating float32

		err = rows.Scan(&id, &name, &description, &ingredients, &cooking_steps, &cooking_time, &recipe_rating)

		// check errors
		checkErr(err)

		recipes = append(recipes, Recipe{Id: id, Name: name, Description: description, Ingredients: ingredients, Cooking_steps: cooking_steps, Cooking_time: cooking_time, Recipe_rating: recipe_rating})
	}

	var response = JsonResponse{Type: "success", Data: recipes, Message: "Got all recipes"}

	json.NewEncoder(w).Encode(response)
}

func GetRecipe(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConfig()

	db := setupDB(cfg.Storage)

	printMessage("Getting one recipe")

	ReqId := r.FormValue("id")

	var response = JsonResponse{}

	if ReqId == "" {
		response = JsonResponse{Type: "error", Message: "You need to insert id, id is null"}
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

	// check errors
	checkErr(err)

	recipe := Recipe{Id: id, Name: name, Description: description, Ingredients: ingredients, Cooking_steps: cooking_steps, Cooking_time: cooking_time, Recipe_rating: recipe_rating}

	//response = JsonResponse{Type: "error",Data: recipe, Message: "You need to insert id, id is null"}

	json.NewEncoder(w).Encode(recipe)
}

func GetRecipesSortedByIngredients(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConfig()

	db := setupDB(cfg.Storage)

	printMessage("Sorted by Ingridients")

	rows, err := db.Query(cfg.DBqueries.GetRecipesSortedByIngredients)

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var recipes []Recipe

	// Foreach movie
	for rows.Next() {
		var id int
		var name string
		var description string
		var ingredients string
		var cooking_steps string
		var cooking_time string
		var recipe_rating float32

		err = rows.Scan(&id, &name, &description, &ingredients, &cooking_steps, &cooking_time, &recipe_rating)

		// check errors
		checkErr(err)

		recipes = append(recipes, Recipe{Id: id, Name: name, Description: description, Ingredients: ingredients, Cooking_steps: cooking_steps, Cooking_time: cooking_time, Recipe_rating: recipe_rating})
	}

	var response = JsonResponse{Type: "success", Data: recipes, Message: "recipes filtred by ingridients"}

	json.NewEncoder(w).Encode(response)

}

func GetRecipesSortedByCookingTime(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConfig()

	db := setupDB(cfg.Storage)

	printMessage("Sorted by Time")

	rows, err := db.Query(cfg.DBqueries.GetAllRecipes)

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var oldRecipes []Recipe

	// Foreach movie
	for rows.Next() {
		var id int
		var name string
		var description string
		var ingredients string
		var cooking_steps string
		var cooking_time string
		var recipe_rating float32

		err = rows.Scan(&id, &name, &description, &ingredients, &cooking_steps, &cooking_time, &recipe_rating)

		// check errors
		checkErr(err)

		oldRecipes = append(oldRecipes, Recipe{Id: id, Name: name, Description: description, Ingredients: ingredients, Cooking_steps: cooking_steps, Cooking_time: cooking_time, Recipe_rating: recipe_rating})
	}

	recipes := sortRecipesByTime(oldRecipes)

	var response = JsonResponse{Type: "success", Data: recipes, Message: "recipes filtred by cooking time"}

	json.NewEncoder(w).Encode(response)

}

func GetRecipesSortedByRating(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConfig()

	db := setupDB(cfg.Storage)

	printMessage("Sorted by Rating")

	rows, err := db.Query(cfg.DBqueries.GetAllRecipes)

	checkErr(err)

	var oldRecipes []Recipe

	for rows.Next() {
		var id int
		var name string
		var description string
		var ingredients string
		var cooking_steps string
		var cooking_time string
		var recipe_rating float32

		err = rows.Scan(&id, &name, &description, &ingredients, &cooking_steps, &cooking_time, &recipe_rating)

		checkErr(err)

		oldRecipes = append(oldRecipes, Recipe{Id: id, Name: name, Description: description, Ingredients: ingredients, Cooking_steps: cooking_steps, Cooking_time: cooking_time, Recipe_rating: recipe_rating})
	}

	recipes := sortRecipesByRating(oldRecipes)

	var response = JsonResponse{Type: "success", Data: recipes, Message: "recipes filtred by rating"}

	json.NewEncoder(w).Encode(response)

}
