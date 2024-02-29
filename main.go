package main

import(
"github.com/gin-gonic/gin"
"net/http"
"strconv"
)


type Spider struct {
	ID          string `json:"id"`
  Nombre      string `json:"nombre"`
  Descripcion string `json:"descripcion"`
}
//YERSON
type Oso struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

var osos = []Oso{
	{ID: 1, Nombre: "Oso Negro Americano", Descripcion: "Habita en Norteamérica."},
	{ID: 2, Nombre: "Oso Polar", Descripcion: "Habita en el Ártico."},
	{ID: 3, Nombre: "Oso Panda", Descripcion: "Caracterizado por su pelaje blanco y negro."},
}

func getOso(c *gin.Context) {
    c.IndentedJSON(http.StatusOK,osos)
}

func getOsoPorID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error al convertir ID a entero"})
        return
    }

    for _, oso := range osos {
        if oso.ID == id {
            c.JSON(http.StatusOK, oso)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Oso no encontrado"})
}

type Leon struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
}


var spider = []Spider{
	{ID: "1", Nombre: "Tarantulas", Description: "arañas migalomorfas de gran tamaño con el cuerpo cubierto por pelos llamados sedas."},
	{ID: "2", Nombre: "Viudas Negras", Description: "tiene un cuerpo negro y brillante con una forma de reloj de arena rojo en la zona ventral"},
	{ID: "3", Nombre: "Licósidos", Description: "arañas que vagan en el suelo, excavando pequeñas galerías verticales u ocupando grietas naturales desde las que acechan a sus presas, cuya presencia detectan por las vibraciones del suelo."},
}

func getSpider(a *gin.Context) {
	a.IndentedJSON(http.StatusOK, spider)
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

	router.GET("/spider", getSpider)
	router.GET("/spider/:id", getSpiderByID)
	router.GET("/leones", getleon)
	router.GET("/leones/:id", getLeonesID)
	router.GET("/aranas", getAranas)
	router.GET("/aranas/:id", getAranaByID)
	router.GET("/carambombo", getCarambombo)
	router.GET("/carambombo/:id", getCarambomboByID)
  router.GET("/osos", getOso)
  router.GET("/osos/:id", getOsoPorID)


  
  
  
  router.Run("localhost:4000")
}


  



