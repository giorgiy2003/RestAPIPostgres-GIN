package main

import (
	Handler "app/internal/handlers"
	Repository "app/internal/repository"
	"log"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	err := Repository.OpenTable()
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.GET("/person", Handler.GetPersons)
	router.GET("/person/:id", Handler.GetById)
	router.POST("/person", Handler.PostPerson)
	router.DELETE("/person/:id", Handler.DeleteById)
	router.PUT("/person/:id", Handler.UpdatePersonById)
	_ = router.Run("localhost:8080")
}
