package main

import(
"github.com/gin-gonic/gin"
"net/http"
"strconv"
)

type Arana struct {
	ID           int     `json:ID`
	Tipo string `json:ID`
	Nombre string `json:Nombre`
}


type Carambombo struct {
	ID int `json:"ID"`
	Nombre string `json:"tipo"`
	Descripcion string `json:"descripcion"`
}

var aranas = []Arana{
	{ID: 1, Tipo: "Tarantula", Nombre: "kidd keo"},
	{ID: 2, Tipo: "Aranita chiquita", Nombre: "Ariana Grande"},
}

var carambombos = []Carambombo{
	{ID: 1, Nombre:"Carambombo normal", Descripcion:"es amarillo"},
	{ID: 1, Nombre:"Carambombo Verdoso", Descripcion:"es verde"},
	{ID: 1, Nombre:"Carambombo Azu", Descripcion:"es azul"},
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

func main() {
	router := gin.Default()
	router.GET("/aranas", getAranas)
	router.GET("/aranas/:id", getAranaByID)
	router.GET("/carambombo", getCarambombo)
	router.GET("/carambombo/:id", getCarambomboByID)










	router.Run("localhost:4000")
}

