package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api-gin/database"
	"api-gin/models"
)

func Hello(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"Saudações": nome + " " + "tudo bem?",
	})
}

func GetAllData(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)

}

func CreateNewData(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := models.ValidateData(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func GetByID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "id não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)

}

func DeleteData(c *gin.Context) {
	var deleted models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&deleted, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "ok",
	})

}

func EditiData(c *gin.Context) {
	var edite models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&edite, id)
	if err := c.ShouldBindJSON(&edite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	if err := models.ValidateData(&edite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	database.DB.Model(&edite).UpdateColumns(edite)
	c.JSON(http.StatusOK, edite)

}

func FindTodata(c *gin.Context) {
	var CPFData models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&CPFData)
	if CPFData.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, CPFData)
}

func HtmlHandler(c *gin.Context) { // função para renderizar paginas html
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func HandlerRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")   // informa o local onde o arquivo html está registrado
	r.Static("/assets", "./assets") // informa o local onde o arquivo css está registrado
	r.GET("/alunos", GetAllData)
	r.GET("/:nome", Hello)
	r.POST("/alunos", CreateNewData)
	r.GET("/alunos/:id", GetByID)
	r.DELETE("/alunos/:id", DeleteData)
	r.PATCH("/alunos/:id", EditiData)
	r.GET("/alunos/cpf/:cpf", FindTodata)
	r.GET("/index", HtmlHandler)
	r.Run()
}
