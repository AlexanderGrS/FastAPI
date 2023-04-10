package main

import (
	"FastAPI/internal/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/recipes/GetAllRecipes", handlers.GetAllRecipes).Methods("GET")
	router.HandleFunc("/recipes/GetRecipe", handlers.GetRecipe).Methods("GET")
	router.HandleFunc("/recipes/GetRecipesSortedByIngredients", handlers.GetRecipesSortedByIngredients).Methods("GET")
	router.HandleFunc("/recipes/GetRecipesSortedByCookingTime", handlers.GetRecipesSortedByCookingTime).Methods("GET")
	router.HandleFunc("/recipes/GetRecipesSortedByRating", handlers.GetRecipesSortedByRating).Methods("GET")

	router.Headers().Methods()
	router.HandleFunc("/recipes/CreateRecipe", handlers.CreateRecipe).Methods("POST")
	router.HandleFunc("/recipes/ChangeRecipe", handlers.ChangeRecipe).Methods("POST")
	router.HandleFunc("/recipes/DeleteRecipe", handlers.DeleteRecipe).Methods("POST")
	router.HandleFunc("/recipes/SortingRecipesByCookingTime", handlers.SortingRecipesByCookingTime).Methods("POST")
	router.HandleFunc("/recipes/SortingRecipesByRating", handlers.SortingRecipesByRating).Methods("POST")

	// serve the app
	fmt.Println("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
