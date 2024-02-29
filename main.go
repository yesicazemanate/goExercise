package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
type Leon struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}

var leones = []Leon{
	{ID: 1, Nombre: "Leon verde", Description: "es muy grande"},
	{ID: 2, Nombre: "Leon rojo", Description: "es muy grande"},
	{ID: 3, Nombre: "Leon negro", Description: "es muy grande"},
	{ID: 4, Nombre: "Leon gris", Description: "es muy grande"},

}

func getleon(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, leones)
}

func getLeonesID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}


	var animal Leon
	for _, a := range leones{
		if a.ID == id {
			animal = a
			break
		}
	}


	if animal.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Animal no encontrado"})
		return
	}


	c.IndentedJSON(http.StatusOK, animal)
}

func main() {
	router := gin.Default()
	router.GET("/leones", getleon)
	router.GET("/leones/:id", getLeonesID)









	router.Run("localhost:4000")
}
