package main

import (
	"FastAPI/auth/authdb"
	"FastAPI/auth/middleware"
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

	router.Handle("/recipes/CreateRecipe", middleware.BasicAuth(http.HandlerFunc(handlers.CreateRecipe))).Methods("POST")
	router.Handle("/recipes/ChangeRecipe", middleware.BasicAuth(http.HandlerFunc(handlers.ChangeRecipe))).Methods("POST")
	router.Handle("/recipes/DeleteRecipe", middleware.BasicAuth(http.HandlerFunc(handlers.DeleteRecipe))).Methods("POST")
	router.Handle("/recipes/SortingRecipesByCookingTime", middleware.BasicAuth(http.HandlerFunc(handlers.SortingRecipesByCookingTime))).Methods("POST")
	router.Handle("/recipes/SortingRecipesByRating", middleware.BasicAuth(http.HandlerFunc(handlers.SortingRecipesByRating))).Methods("POST")

	router.HandleFunc("/auth/SignUp", authdb.SignUp).Methods("POST")

	fmt.Println("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
