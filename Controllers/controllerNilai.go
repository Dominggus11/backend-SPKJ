package controllers

import (
	"net/http"
	models "spkj/Models"

	"github.com/gin-gonic/gin"
)

func GetNilai(c *gin.Context) {
	db := models.DB
	// Get model if exist
	var student models.Students
	if err := db.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Error Get Data",
			"message":  "Data Siswa Tidak Tersedia",
			"response": "409"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"nama":          student.Nama,
		"nisn":          student.NISN,
		"ujiansekolah":  student.UjianSekolah,
		"rerata raport": student.RerataRaport,
		"ipa":           student.IPA,
		"ips":           student.IPS,
		"minat":         student.Minat,
		"response":      "200"})
}

func GetNormalisasi(c *gin.Context) {
	db := models.DB
	// Get model if exist
	var student models.Students
	if err := db.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Error Get Data",
			"message":  "Data Siswa Tidak Tersedia",
			"response": "409"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"nama":          student.Nama,
		"nisn":          student.NISN,
		"ujiansekolah":  student.Ci_UjianSekolah,
		"rerata raport": student.Ci_RerataRaport,
		"ipa":           student.Ci_IPA,
		"ips":           student.Ci_IPS,
		"minat":         student.Ci_Minat,
		"response":      "200"})
}
