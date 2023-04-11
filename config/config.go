package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Storage   StorageConfig   `yaml:"storage"`
	DBqueries DBqueriesConfig `yaml:"DBqueries"`
}

type StorageConfig struct {
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type DBqueriesConfig struct {
	GetAllRecipes                 string `json:"getallrecipes"`
	CreateRecipes                 string `json:"createrecipes"`
	GetRecipe                     string `json:"getrecipe"`
	ChangeRecipe                  string `json:"changerecipe"`
	DeleteRecipe                  string `json:"deleterecipe"`
	GetRecipesSortedByIngredients string `json:"getrecipessortedbyingredients"`
	DeleteTable                   string `json:"deletetable"`
	CreateTable                   string `json:"createtable"`
	SignUp                        string `json:"signup"`
	SignIn                        string `json:"signin"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		err := cleanenv.ReadConfig("config.yml", instance)
		if err != nil {
			panic(err)
		}
	})
	return instance
}
