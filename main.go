package main
import(
"github.com/gin-gonic/gin"
"net/http"
)

type Alimento struct{
	ID int `json:"id"`
	Nombre string `json:"nombre"`
	Description string `json:"description"`
}

var alimentos = []Alimento{
	 {ID: 1, Nombre:"Arepas de choclo", Description: "Hechas de maíz"},
	{ID: 1, Nombre:"Arepitas", Description: "Hechas de maicitos"},
	{ID: 1, Nombre:"Buñuelos", Description: "redondos"},
}
func getAlimento(a *gin.Context){
a.IndentedJSON(http.StatusOK, alimentos)
}
func main(){
router := gin.Default()
router.GET("/alimentos", getAlimento)
router.Run("localhost:4000")
}