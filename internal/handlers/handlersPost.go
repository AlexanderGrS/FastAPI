package handlers

import (
	"FastAPI/config"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	description := r.FormValue("description")
	ingredients := r.FormValue("ingredients")
	cooking_steps := r.FormValue("cooking_steps")
	cooking_time := r.FormValue("cooking_time")
	recipe_rating := r.FormValue("recipe_rating")

	cfg := config.GetConfig()

	db := setupDB(cfg.Storage)

	printMessage("Inserting movie into DB")

	fmt.Println(name, description, ingredients, cooking_steps, cooking_time, recipe_rating, "Insert this")

	var lastInsertID int
	err := db.QueryRow(cfg.DBqueries.CreateRecipes, name, description, ingredients, cooking_steps, cooking_time, recipe_rating).Scan(&lastInsertID)

	checkErr(err)

	var response = JsonResponse{Type: "success", Message: "Insert new recipe"}

	json.NewEncoder(w).Encode(response)
}

func ChangeRecipe(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	description := r.FormValue("description")
	ingredients := r.FormValue("ingredients")
	cooking_steps := r.FormValue("cooking_steps")
	cooking_time := r.FormValue("cooking_time")
	recipe_rating := r.FormValue("recipe_rating")

	var response = JsonResponse{}

	if id == "" {
		response = JsonResponse{Type: "error", Message: "You need to insert id, id is null"}
		json.NewEncoder(w).Encode(response)
	}

	var oldRecipe Recipe

	cfg := config.GetConfig()

	db := setupDB(cfg.Storage)

	row := db.QueryRow(cfg.DBqueries.GetRecipe, id)

	err := row.Scan(&oldRecipe.Id, &oldRecipe.Name, &oldRecipe.Description, &oldRecipe.Ingredients, &oldRecipe.Cooking_steps, &oldRecipe.Cooking_time, &oldRecipe.Recipe_rating)

	checkErr(err)

	printMessage("Changing recipe")

	fmt.Println(id, name, description, ingredients, cooking_steps, cooking_time, recipe_rating, "Insert this")

	if name == "" {
		name = oldRecipe.Name
	}
	if description == "" {
		description = oldRecipe.Description
	}
	if ingredients == "" {
		ingredients = oldRecipe.Ingredients
	}
	if cooking_steps == "" {
		cooking_steps = oldRecipe.Cooking_steps
	}
	if cooking_time == "" {
		cooking_time = oldRecipe.Cooking_time
	}
	if recipe_rating == "" {
		recipe_rating = fmt.Sprintf("%.2f", oldRecipe.Recipe_rating)
	}
	fmt.Println(id, name, description, ingredients, cooking_steps, cooking_time, recipe_rating, "Insert this")

	_, err = db.Exec(cfg.DBqueries.ChangeRecipe, name, description, ingredients, cooking_steps, cooking_time, recipe_rating, id)

	checkErr(err)

	response = JsonResponse{Type: "success", Message: fmt.Sprintf("recipe with Id = %s successfuly changed", id)}

	json.NewEncoder(w).Encode(response)
}

func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	var response = JsonResponse{}

	if id == "" {
		response = JsonResponse{Type: "error", Message: "You need to insert id, id is null"}
		json.NewEncoder(w).Encode(response)
	}

	cfg := config.GetConfig()

	db := setupDB(cfg.Storage)

	printMessage("Deleting recipe")

	fmt.Println(id, "Delete this")

	_, err := db.Exec(cfg.DBqueries.DeleteRecipe, id)

	checkErr(err)

	response = JsonResponse{Type: "success", Message: fmt.Sprintf("recipe with Id = %s successfuly deleted", id)}

	json.NewEncoder(w).Encode(response)
}

func SortingRecipesByRating(w http.ResponseWriter, r *http.Request) {
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

	_, err = db.Exec(cfg.DBqueries.DeleteTable)

	checkErr(err)

	_, err = db.Exec(cfg.DBqueries.CreateTable)

	checkErr(err)

	for i := 0; i < len(recipes); i++ {
		var lastInsertID int
		err := db.QueryRow(cfg.DBqueries.CreateRecipes, recipes[i].Name, recipes[i].Description, recipes[i].Ingredients, recipes[i].Cooking_steps, recipes[i].Cooking_time, recipes[i].Recipe_rating).Scan(&lastInsertID)

		checkErr(err)
	}

	var response = JsonResponse{Type: "success", Message: "Table recipes sorted by rating"}

	json.NewEncoder(w).Encode(response)

}

func SortingRecipesByCookingTime(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConfig()

	db := setupDB(cfg.Storage)

	printMessage("Sorted by Time")

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

	recipes := sortRecipesByTime(oldRecipes)

	_, err = db.Exec(cfg.DBqueries.DeleteTable)

	checkErr(err)

	_, err = db.Exec(cfg.DBqueries.CreateTable)

	checkErr(err)

	for i := 0; i < len(recipes); i++ {
		var lastInsertID int
		err := db.QueryRow(cfg.DBqueries.CreateRecipes, recipes[i].Name, recipes[i].Description, recipes[i].Ingredients, recipes[i].Cooking_steps, recipes[i].Cooking_time, recipes[i].Recipe_rating).Scan(&lastInsertID)

		checkErr(err)
	}

	var response = JsonResponse{Type: "success", Message: "Table recipes sorted by cooking time"}

	json.NewEncoder(w).Encode(response)

}
