package main

import(
"github.com/gin-gonic/gin"
"net/http"
"strconv"
)
//Brayn Delgado//
 type Server struct {
	ID          int
	Nombre      string
	Description string
	// Otros campos que puedas necesitar
}
type Motos struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}


var frutass = []Server{
	{ID: 1, Nombre: "Manzana", Description: "fruta que se da en climas frios"},
	{ID: 2, Nombre: "Uva", Description: "fruta dulce"},
	{ID: 3, Nombre: "Banano", Description: "fruta con mucha proteina"},
	{ID: 4, Nombre: "Mango", Description: "fruta tropical"},
}
var motos = []Motos{
	{ID: "1", Nombre: "yamaha", Description: "velocidad 1"},
	{ID: "2", Nombre: "honda", Description: "velocidad 2"},
	{ID: "3", Nombre: "cbr", Description: "velocidad 3"},
}


func buscarServerPorID(idBuscado int) *Server {
	for _, server := range frutass {
		if server.ID == idBuscado {
			return &server // Retorna una referencia al objeto encontrado
		}
	}
	return nil // Retorna nil si no se encuentra el objeto
}

func getServerByID(c *gin.Context) {
	// Obtener el ID del objeto desde la URL
	idStr := c.Param("id")

	// Convertir el ID de string a int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Buscar el objeto Server por ID
	serverEncontrado := buscarServerPorID(id)

	// Verificar si el objeto fue encontrado
	if serverEncontrado != nil {
		// Devolver el objeto encontrado como JSON
		c.JSON(http.StatusOK, serverEncontrado)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Objeto no encontrado"})
	}
}

func getMotos(a *gin.Context) {
	a.IndentedJSON(http.StatusOK, motos)
}

func getMotosByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range motos {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}

func main() {
	router := gin.Default()
	router.GET("/server/:id", getServerByID)
	router.GET("/motos", getMotos)
	router.GET("/motos/:id", getMotosByID)
	router.Run("localhost:4000")
}


  



