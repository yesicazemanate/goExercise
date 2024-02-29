package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Arana struct {
	ID     int    `json:ID`
	Tipo   string `json:ID`
	Nombre string `json:Nombre`
}

type Carambombo struct {
	ID          int    `json:"ID"`
	Nombre      string `json:"tipo"`
	Descripcion string `json:"descripcion"`
}

var aranas = []Arana{
	{ID: 1, Tipo: "Tarantula", Nombre: "kidd keo"},
	{ID: 2, Tipo: "Aranita chiquita", Nombre: "Ariana Grande"},
}

var carambombos = []Carambombo{
	{ID: 1, Nombre: "Carambombo normal", Descripcion: "es amarillo"},
	{ID: 1, Nombre: "Carambombo Verdoso", Descripcion: "es verde"},
	{ID: 1, Nombre: "Carambombo Azu", Descripcion: "es azul"},
}

func getAranas(a *gin.Context) {
	a.IndentedJSON(http.StatusOK, aranas)
}

func getAranaByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID no válido"})
		return
	}

	var arana Arana
	for _, a := range aranas {
		if a.ID == id {
			arana = a
			break
		}
	}

	if arana.Nombre == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Arana no encontrada"})
		return
	}

	c.JSON(http.StatusOK, arana)
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

// Luisa Villacorte
type Naranjas struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}

var naranjas = []Naranjas{
	{ID: 1, Nombre: "Toronja", Description: "Fruta cítrica de sabor dulce originaria del Caribe."},
	{ID: 2, Nombre: "Sanguinello", Description: "Pulpa roja y color intenso que puede ir del naranja al rojo."},
	{ID: 3, Nombre: "Tarocco", Description: "Su pulpa puede tener vetas rojas debido a la presencia de antocianinas y casi no tienen semillas. Es ideal para zumos dulces."},
}

func getNaranjas(a *gin.Context) {
	a.IndentedJSON(http.StatusOK, naranjas)
}
func getNaranjaByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID no válido"})
		return
	}
	var naranjaEncontrada Naranjas
	for _, naranja := range naranjas {
		if naranja.ID == id {
			naranjaEncontrada = naranja
			break
		}
	}
	if naranjaEncontrada.Nombre == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Naranja no encontrada"})
		return
	}

	c.JSON(http.StatusOK, naranjaEncontrada)
}

//Luisa Villacorte

func main() {
	router := gin.Default()
	router.GET("/aranas", getAranas)
	router.GET("/aranas/:id", getAranaByID)
	router.GET("/carambombo", getCarambombo)
	router.GET("/carambombo/:id", getCarambomboByID)
	router.GET("/naranja", getNaranjas)
	router.GET("/naranja/:id", getNaranjaByID)
	router.Run("localhost:4000")
}
