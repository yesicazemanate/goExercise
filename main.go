package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
//Brayan Delgado
type Server struct {
    ID          int
    Nombre      string
    Description string
    
}

// Lista de objetos de ejemplo
type Alimento struct{
	ID int `json:"id"`
	Nombre string `json:"nombre"`
	Description string `json:"description"`
}

type Frutas struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}

type Arana struct {
	ID           int     `json:ID`
	Tipo string `json:ID`
	Nombre string `json:Nombre`
}
type Animales struct {
	ID         int     `json:ID`
	Tipo string `json:ID`
	Nombre string `json:Nombre`
}

var aranas = []Arana{
	{ID: 1, Tipo: "Tarantula", Nombre: "kidd keo"},
	{ID: 2, Tipo: "Aranita chiquita", Nombre: "Ariana Grande"},
}
//Maryuri
var pulgas = []Animales{
	{ID: 1, Tipo: "Tungidae.", Nombre: "En casa"},
	{ID: 2, Tipo: "Pulga de perros", Nombre: "En caballo"},
	{ID: 3, Tipo: "Pulga de gatos", Nombre: "En caballo"},
	{ID: 4, Tipo: "Pulga de santi", Nombre: "En caballo"},
}

func getAranas(a *gin.Context) {
	a.IndentedJSON(http.StatusOK, aranas)
}

//Get Maryuri
func getPulga(m *gin.Context) {
	m.IndentedJSON(http.StatusOK, pulgas)
}
//GET by Id MARYURI
func getByIdPulga(m * gin.Context){
	idStr := m.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
			m.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found"})
	}
	for _, ma := range pulgas{
		if ma.ID == id{
			m.IndentedJSON(http.StatusOK, ma)
			return
		}
	}

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
////////SANTIAGO NARVAEZ LASSO


// LUISA VILLACORTE
var fruits = []Frutas{
	{ID: 1, Nombre: "Sandia", Description: "Fruta tropical"},
	{ID: 2, Nombre: "Limón", Description: "Fruta ácida"},
	{ID: 3, Nombre: "Naranja", Description: "Fruta semi-ácida"},
}
func getFrutas(a *gin.Context) {
	a.IndentedJSON(http.StatusOK, fruits)
}

func getFrutaByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID no válido"})
		return
	}

	var frutaEncontrada Frutas
	for _, fruta := range fruits {
		if fruta.ID == id {
			frutaEncontrada = fruta
			break
		}
	}

	if frutaEncontrada.Nombre == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fruta no encontrada"})
		return
	}

	c.JSON(http.StatusOK, frutaEncontrada)
}

<<<<<<< HEAD
//Darbin
type Animal struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}

var animales = []Animal{
	{ID: 1, Nombre: "Leon verde", Description: "es muy grande"},
	{ID: 2, Nombre: "Leon rojo", Description: "es muy grande"},
	{ID: 3, Nombre: "Leon negro", Description: "es muy grande"},
	{ID: 4, Nombre: "Leon gris", Description: "es muy grande"},

}

func getAnimales(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, animales)
}

func getAnimal(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}


	var animal Animal
	for _, a := range animales {
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
// BRAYAN DELGADO//
var frutas = []*Server{
	{ID: 1, Nombre: "Manzana", Description: "fruta que se da en climas frios"},
	{ID: 2, Nombre: "Uva", Description: "fruta dulce"},
	{ID: 3, Nombre: "Banano", Description: "fruta con mucha proteina"},
	{ID: 4, Nombre: "Mango", Description: "fruta tropical"},
}

// Función para buscar una fruta por su ID en una lista de frutas
func buscarFrutaPorID(idBuscado int) *Fruta {
	for _, fruta := range frutas {
		if fruta.ID == idBuscado {
			return &fruta // Retorna una referencia a la fruta encontrada
		}
	}
	return nil // Retorna nil si no se encuentra la fruta
}

// Función handler para obtener una fruta por su ID
func getFrutaByIDD(c *gin.Context) {
	// Obtener el ID de la fruta desde la URL
	idStr := c.Param("id")

	// Convertir el ID de string a int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Buscar la fruta por ID
	frutaEncontrada := buscarFrutaPorID(id)

	// Verificar si la fruta fue encontrada
	if frutaEncontrada != nil {
		// Devolver la fruta encontrada como JSON
		c.JSON(http.StatusOK, frutaEncontrada)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Fruta no encontrada"})
	}
}
=======
// LUISA VILLACORTE

>>>>>>> Frutas-Lu

func main(){
	router := gin.Default()
	router.GET("/alimentos", getAlimento)
	router.GET("/aranas", getAranas)
	router.GET("/aranas/:id", getAranaByID)
<<<<<<< HEAD
<<<<<<< HEAD
	router.GET("/pulgasM", getPulga)
	router.GET("/pulgasM/:id", getByIdPulga)
=======
	router.GET("/animales", getAnimales)
	router.GET("/animales/:id", getAnimal)
	router.GET("/frutas/:id", getFrutaByIDD)
>>>>>>> 7016b22bc4ffaafc86cc9991cadfd072d52471dc
=======
	router.GET("/frutas", getFrutas)
	router.GET("/frutas/:id", getFrutaByID)
>>>>>>> Frutas-Lu
	router.Run("localhost:4000")
}