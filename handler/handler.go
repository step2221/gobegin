package handler

import (
	"fmt"
	"net/http"
	"restapi/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

//ErrorResponse .
type ErrorResponse struct {
	Message string `json:"message"`
}

//Handler  .
type Handler struct {
	storage storage.Storage
}

//NewHandler .
func NewHandler(storage storage.Storage) *Handler {
	return &Handler{storage: storage}
}

//GetEmployee .
func (h *Handler) GetEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ID"))

	fmt.Printf("%d", id)
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	employee, err := h.storage.Get(id)
	if err != nil {
		fmt.Printf("failed to get employee %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, employee)
}

//DeleteEmployee .
func (h *Handler) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	print(id)
	h.storage.Delete(id)

	c.String(http.StatusOK, "employee deleted")
}

//CreateEmployee .
func (h *Handler) CreateEmployee(c *gin.Context) {
	var employee storage.Employee

	if err := c.BindJSON(&employee); err != nil {
		fmt.Printf("failed to bind employee: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.storage.Insert(&employee)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": employee.ID,
	})
}

//UpdateEmployee .
func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ID"))
	fmt.Printf("Изменяем запись с id = %d", id)
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var employee storage.Employee

	if err := c.BindJSON(&employee); err != nil {
		fmt.Printf("failed to bind employee: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.storage.Update(id, employee)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": employee.ID,
	})
}
