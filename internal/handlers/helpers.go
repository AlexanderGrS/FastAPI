package handlers

import (
	"fmt"
	"strconv"
)

const (
	hourInKirilic   = 135
	minuteInKirilic = 188
	secondInKirilic = 129
)

func sortRecipesByRating(oldRecipes []Recipe) (recipes []Recipe) {
	var orderedIndexes []int
	var oldRecipeSlice []Recipe
	for i := 0; i < len(oldRecipes); i++ {
		oldRecipeSlice = append(oldRecipeSlice, oldRecipes[i])
	}
	for i := 0; i < len(oldRecipeSlice); i++ {
		var maxElement float32
		var currentId, needToDeleteId int
		for j := 0; j < len(oldRecipeSlice); j++ {
			if maxElement < oldRecipeSlice[j].Recipe_rating {
				maxElement = oldRecipeSlice[j].Recipe_rating
				currentId = oldRecipeSlice[j].Id
				needToDeleteId = j
			}
			if j == len(oldRecipeSlice)-1 {
				oldRecipeSlice[needToDeleteId].Recipe_rating = -1
			}
		}
		orderedIndexes = append(orderedIndexes, currentId)
	}
	o := oldRecipes
	for i := 0; i < len(orderedIndexes); i++ {
		for j := 0; j < len(o); j++ {
			if o[j].Id == orderedIndexes[i] {
				recipes = append(recipes, Recipe{Id: o[j].Id, Name: o[j].Name, Description: o[j].Description, Ingredients: o[j].Ingredients, Cooking_steps: o[j].Cooking_steps, Cooking_time: o[j].Cooking_time, Recipe_rating: o[j].Recipe_rating})
			}
		}
	}
	return recipes
}

func sortRecipesByTime(oldRecipes []Recipe) (recipes []Recipe) {
	recipesWithIdAndTime := parseCookingTime(oldRecipes)
	var orderedIndexes []int
	for len(recipesWithIdAndTime) != 0 {
		var maxElement, currentId int
		for i, el := range recipesWithIdAndTime {
			if maxElement < el {
				maxElement = el
				currentId = i
			}
		}
		orderedIndexes = append(orderedIndexes, currentId)
		delete(recipesWithIdAndTime, currentId)
	}
	o := oldRecipes
	for i := len(orderedIndexes) - 1; i >= 0; i-- {
		for j := 0; j < len(o); j++ {
			if o[j].Id == orderedIndexes[i] {
				recipes = append(recipes, Recipe{Id: o[j].Id, Name: o[j].Name, Description: o[j].Description, Ingredients: o[j].Ingredients, Cooking_steps: o[j].Cooking_steps, Cooking_time: o[j].Cooking_time, Recipe_rating: o[j].Recipe_rating})
			}
		}
	}
	return recipes
}

func parseCookingTime(recipes []Recipe) map[int]int {
	recipesWithIdAndTime := make(map[int]int)
	for i := 0; i < len(recipes); i++ {
		var timeInSeconds int
		if recipes[i].Cooking_time == "" {
			recipesWithIdAndTime[recipes[i].Id] = 0
		}
		if recipes[i].Cooking_time != "" {
			timeStampStart, timeStampEnd := 0, 1
			strCookingTime := recipes[i].Cooking_time[1 : len(recipes[i].Cooking_time)-1]
			for j := 1; j < len(strCookingTime)-2; j++ {
				switch {
				case strCookingTime[j+2] == hourInKirilic && string(strCookingTime[j]) == " ":
					timeInFloat, _ := strconv.ParseFloat(strCookingTime[timeStampStart:timeStampEnd], 32)
					timeInSeconds += int(timeInFloat * 3600)
				case strCookingTime[j+2] == minuteInKirilic && string(strCookingTime[j]) == " ":
					timeInFloat, _ := strconv.ParseFloat(strCookingTime[timeStampStart:timeStampEnd], 32)
					timeInSeconds += int(timeInFloat * 60)
				case strCookingTime[j+2] == secondInKirilic && string(strCookingTime[j]) == " ":
					timeInFloat, _ := strconv.ParseFloat(strCookingTime[timeStampStart:timeStampEnd], 32)
					timeInSeconds += int(timeInFloat)
				}
				if string(strCookingTime[j]) != " " && (string(strCookingTime[j-1]) == " " || string(strCookingTime[j-1]) == ",") {
					timeStampStart = j
				}
				if string(strCookingTime[j]) != " " && string(strCookingTime[j+1]) == " " {
					timeStampEnd = j + 1
				}
			}
			recipesWithIdAndTime[recipes[i].Id] = timeInSeconds
		}
		timeInSeconds = 0
	}
	return recipesWithIdAndTime
}

func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
