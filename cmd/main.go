package main

import (
	"fmt"

	"github.com/23vugarr/content-sharing-app/handlers"
	"github.com/23vugarr/content-sharing-app/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var err error

func main() {
	router := gin.Default()

	params := models.DBparams{
		DBHost:   "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Dbname:   "postgres",
	}

	db := models.NewDB(params)
	err = db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()

	err = db.DB.Ping()
	if err != nil {
		panic(err)
	}

	router.GET("/login", handlers.PostLogin)
	router.POST("/register", handlers.PostRegister)

	router.Run("localhost:8080")
	fmt.Println("started app on localhost 8080")
}
