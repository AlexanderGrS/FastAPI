package models

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
