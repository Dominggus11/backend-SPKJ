package routes

import (
	controllers "spkj/Controllers"
	models "spkj/Models"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.Use(CORS)
	models.ConnectDatabase()
	r.GET("/", controllers.Developer)

	// untuk API Students
	r.GET("/student", controllers.GetStudents)
	r.GET("/student/:id", controllers.GetStudent)
	r.POST("/student", controllers.PostStudent)
	r.PUT("/student/:id", controllers.PutStudent)
	r.DELETE("/student/:id", controllers.DeleteStudent)

	// untuk API Nilai Siswa
	// r.POST("/nilai/:id", controllers.PostNilai)

	// untuk API SAW
	r.GET("/normalisasi", controllers.GetCi)

	// untuk API nilai
	r.GET("nilai/:id", controllers.GetNilai)
	r.GET("normalisasi/:id", controllers.GetNormalisasi)

	// untuk API Criterias
	r.GET("/kriteria", controllers.GetCriterias)
	r.POST("/kriteria", controllers.PostCriteria)
	r.GET("/kriteria/:id", controllers.GetCriteria)
	r.PUT("/kriteria/:id", controllers.PutCriteria)
	r.DELETE("/kriteria/:id", controllers.DeleteCriteria)

	// untuk API User
	r.POST("/registrasi", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/user", controllers.GetUsers)
	r.Run()
}

func CORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, "+
		"Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}
