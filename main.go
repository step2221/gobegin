package main

import (
	"restapi/handler"
	"restapi/storage"

	"github.com/gin-gonic/gin"
)

//TestFunc .
func TestFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {

	MemoryStorage := storage.NewMemoryStorage()
	handler := handler.NewHandler(MemoryStorage)

	router := gin.Default()
	router.GET("/test", TestFunc)
	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:ID", handler.GetEmployee)
	router.PUT("/employee/:ID", handler.UpdateEmployee)
	router.DELETE("/employee/:ID", handler.DeleteEmployee)

	router.Run()

}
