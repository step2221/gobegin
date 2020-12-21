package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	MemoryStorage := NewMemoryStorage()
	handler := NewHandler(MemoryStorage)

	router := gin.Default()
	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)

	router.Run()

}
