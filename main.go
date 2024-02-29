package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Serpiente struct {
	ID     int    `json:ID`
	Tipo   string `json:"Tipo"`
	Nombre string `json:"Nombre"`
}

var serpiente = []Serpiente{
	{ID: 1, Tipo: "Cobra Real", Nombre: "Ophiophagus"},
	{ID: 2, Tipo: "Mamba Negra ", Nombre: "Dendroaspis "},
}

func getSerpiente(s *gin.Context) {
	s.IndentedJSON(http.StatusOK, serpiente)
}

func getSerpienteByID(s *gin.Context) {
	idse := s.Param("id")
	idint, err := strconv.Atoi(idse)
	if err != nil {
		s.IndentedJSON(http.StatusNotFound, gin.H{"message": "Serpiente no encontrada"})
	}

	for _, e := range serpiente {
		if e.ID == idint {
			s.IndentedJSON(http.StatusOK, e)
		}
	}
}

func main() {
	router := gin.Default()

	router.GET("/serpiente", getSerpiente)
	router.GET("/serpiente/:id", getSerpienteByID)

	router.Run("localhost:4000")
}
