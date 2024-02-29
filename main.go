package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

  
//maryuri
type Animales struct {
	ID  int  `json:ID`
	Tipo string `json:ID`
	Nombre string `json:Nombre`
}
//Maryuri
var pulgas = []Animales{
	{ID: 1, Tipo: "Tungidae.", Nombre: "En casa"},
	{ID: 2, Tipo: "Pulga de perros", Nombre: "En caballo"},
	{ID: 3, Tipo: "Pulga de gatos", Nombre: "En caballo"},
	{ID: 4, Tipo: "Pulga de santi", Nombre: "En caballo"},
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

type AnimalAve struct{
ID int `json:"id"`
Nombre string `json:"nombre"`
Ubicacion string `json:"ubicacion"`
Habitat string `json:"habitat"`
}// yesica zemanate
// yesica zemanate
var ave =[]AnimalAve{
	{ID: 1, Nombre: "Águila calva",Ubicacion:"América del Norte", Habitat: " Bosques cercanos a cuerpos de agua, como ríos, lagos y costas marinas"},
	{ID: 2, Nombre :"Colibrí esmeralda ", Ubicacion: "América Central y del Sur", Habitat: "Bosques tropicales y subtropicales, jardines y áreas con flores"},
	{ID:3, Nombre:"Cóndor de los Andes ", Ubicacion: "Cordilleras de los Andes en América del Sur", Habitat: " Cordilleras montañosas, acantilados y cielos abiertos"},
}// yesica zemanate

func getAve(a *gin.Context){	
  a.IndentedJSON(http.StatusOK, ave)
}

func getAveId(a *gin.Context){
	ids:= a.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		a.JSON(http.StatusNotFound, gin.H{"error": "ID no válido"})
		return
	}
	for _, c := range ave{
		if c.ID == id{
		a.IndentedJSON(http.StatusOK,c )	
		return
	  }
  }
}// yesica zemanate


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


 type Mango struct {
  ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Description string `json:"description"`
 }

  
var mangos = []Mango{
	{ID: 1, Nombre: "Mango verde", Description: "Es un mango que todavía está verde y no es jugoso"},
	{ID: 2, Nombre: "Mango rojo", Description: "Es un mango que es muy jugoso ya que su color rojo lo demuestra"},
	{ID: 3, Nombre: "Mango amarillo", Description: "Es un mango que está entre jugoso y verde"},
}

func getMangoById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID debe ser un número entero"})
		return
	}

	var mango Mango
	for _, item := range mangos {
		if item.ID == id {
			mango = item
			break
		}
	}

	if mango.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mango no encontrado"})
		return
	}

	c.IndentedJSON(http.StatusOK, mango)
}

func getAllMangos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mangos)
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
	router.GET("/leones", getleon)
	router.GET("/leones/:id", getLeonesID)
	router.GET("/aranas", getAranas)
	router.GET("/aranas/:id", getAranaByID)
	router.GET("/carambombo", getCarambombo)
	router.GET("/carambombo/:id", getCarambomboByID)
	router.GET("/naranja", getNaranjas)
	router.GET("/naranja/:id", getNaranjaByID)
	router.GET("/osos", getOso)
	router.GET("/osos/:id", getOsoPorID)
  router.GET("/ave", getAve)
	router.GET("/ave/:id",getAveId)
  router.GET("/pulgasM", getPulga)
  router.GET("/pulgasM/:id", getByIdPulga)
  router.GET("/frutamango/:id", getMangoById)
	router.GET("/frutamango", getAllMangos)

  
  
  router.Run("localhost:4000")
}
