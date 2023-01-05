package controllers

import (
	"net/http"
	models "spkj/Models"

	"github.com/gin-gonic/gin"
)

func PostNilai(c *gin.Context) {
	db := models.DB
	var input models.NilaiSiswas
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": "409",
		})
		return
	}
	if err := db.Where("id = ?", input.StudentID).First(&input).Error; err == nil {
		nilai := models.NilaiSiswas{
			// StudentID:    input.StudentID,
			UjianSekolah: input.UjianSekolah,
			RerataRaport: input.RerataRaport,
		}
		db.Create(&nilai)
		c.JSON(http.StatusOK, gin.H{
			"message":  nilai,
			"response": "200"})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Failed To Create Student",
			"message":  "Data Siswa Tidak Tersedia",
			"response": "409"})
		return
	}
}
