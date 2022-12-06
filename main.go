package main

import (
	"fmt"
	"unittest/db/queries"
	"unittest/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config, err := LoadConfig(".env")

	if err != nil {
		panic(err)
	}

	connUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", config.MySQLUser, config.MySQLPassword, config.MySQLHost, config.MySQLPort, config.MySQLDatabase)
	db, err := gorm.Open(mysql.Open(connUrl))

	if err != nil {
		panic(err)
	}

	userHandler := handlers.NewUserHandler(queries.NewUserQueries(db))

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/users", userHandler.CreateUser)
		v1.GET("/users/:id", userHandler.GetUser)
		v1.GET("/users", userHandler.GetUsers)
	}

	r.Run()
}
