package main

import (
	"restapi/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	memoryStorage := storage.NewMemoryStorage()
	handler := NewHandler(memoryStorage)

	router := gin.Default()
	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)

	router.Run()

}
