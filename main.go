package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Alimento struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}

type Frutas struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}

type Arana struct {
	ID     int    `json:"ID"`
	Tipo   string `json:"Tipo"`
	Nombre string `json:"Nombre"`
}

var alimentos = []Alimento{
	{ID: 1, Nombre: "Arepas de choclo", Description: "Hechas de maíz"},
	{ID: 1, Nombre: "Arepitas", Description: "Hechas de maicitos"},
	{ID: 1, Nombre: "Buñuelos", Description: "redondos"},
}

var aranas = []Arana{
	{ID: 1, Tipo: "Tarantula", Nombre: "kidd keo"},
	{ID: 2, Tipo: "Aranita chiquita", Nombre: "Ariana Grande"},
}

func getAlimento(a *gin.Context) {
	a.IndentedJSON(http.StatusOK, alimentos)
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

var fruits = []Frutas{
	{ID: 1, Nombre: "Sandia", Description: "Fruta tropical"},
	{ID: 2, Nombre: "Limón", Description: "Fruta ácida"},
	{ID: 3, Nombre: "Naranja", Description: "Fruta semi-ácida"},
}

func getFrutaByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de ID inválido"})
		return
	}

	// Encontrar la fruta por ID
	var frutaEncontrada *Frutas
	for _, c := range fruits {
		if c.ID == id {
			frutaEncontrada = &c
			break

		}
	}

	// Comprobar si se encontró la fruta
	// if frutaEncontrada.ID == nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "Fruta no encontrada"})
	// 	return
	// }

	// Devolver la fruta encontrada
	c.IndentedJSON(http.StatusOK, frutaEncontrada)
}

// Inicio Carambombos Steven
type Carambombo struct {
	ID          int    `json:"ID"`
	Nombre      string `json:"tipo"`
	Description string `json:"descripcion"`
}

var carambombos = []Carambombo{
	{ID: 1, Nombre: "Carambombo normal", Description: "es amarillo"},
	{ID: 1, Nombre: "Carambombo Verdoso", Description: "es verde"},
	{ID: 1, Nombre: "Carambombo Azu", Description: "es azul"},
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

	var carambos Carambombo
	for _, carambombos := range carambombos {
		if carambombos.ID == id {
			carambos = carambombos
			break
		}
	}

	y.JSON(http.StatusOK, carambos)
}
// Fin Carambombos Steven
func main() {
	router := gin.Default()
	router.GET("/alimentos", getAlimento)
	router.GET("/aranas", getAranas)
	router.GET("/aranas/:id", getAranaByID)
	router.GET("/carambombos", getCarambombo)
	router.GET("/carambombos/:id", getCarambomboByID)
	router.Run("localhost:4000")
}
