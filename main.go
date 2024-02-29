package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
type Carambombo struct {
	ID int `json:"ID"`
	Nombre string `json:"tipo"`
	Descripcion string `json:"descripcion"`
}
var carambombos = []Carambombo{
	{ID: 1, Nombre:"Carambombo normal", Descripcion:"es amarillo"},
	{ID: 1, Nombre:"Carambombo Verdoso", Descripcion:"es verde"},
	{ID: 1, Nombre:"Carambombo Azu", Descripcion:"es azul"},
}
func getCarambombo(z *gin.Context) {
	z.IndentedJSON(http.StatusOK, carambombos)
}
func getCarambomboByID(y *gin.Context) {
	idStr := y.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		y.JSON(http.StatusBadRequest, gin.H{"error": "ID no válido"})
		return
	}

	var carambo Carambombo
	for _, carambombo := range carambombos {
		if carambombo.ID == id {
			carambo = carambombo
			break
		}
	}

	y.JSON(http.StatusOK, carambo)
}

func main() {
	router := gin.Default()
	router.GET("/carambombo", getCarambombo)
	router.GET("/carambombo/:id", getCarambomboByID)









	router.Run("localhost:4000")
}
