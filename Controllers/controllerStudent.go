package controllers

import (
	"net/http"
	models "spkj/Models"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Students
	models.DB.Find(&students)

	c.JSON(http.StatusOK, gin.H{
		"message":  students,
		"response": "200"})

	siswas := GetDataSiswa()
	Normalisasi(c, siswas)
}

func GetStudent(c *gin.Context) {
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
		"message":  student,
		"response": "200"})
}

func PostStudent(c *gin.Context) {

	db := models.DB
	var input models.Students
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": "409",
		})
		return
	}
	input.Ci_UjianSekolah, input.Ci_RerataRaport, input.Ci_IPA, input.Ci_IPS, input.Ci_Minat = BeforeNormalisasi(input.UjianSekolah, input.RerataRaport, input.IPA, input.IPS, input.Minat)
	if err := db.Where("nisn = ?", input.NISN).First(&input).Error; err != nil {
		student := models.Students{
			Nama:            input.Nama,
			NISN:            input.NISN,
			UjianSekolah:    input.UjianSekolah,
			RerataRaport:    input.RerataRaport,
			IPA:             input.IPA,
			IPS:             input.IPS,
			Minat:           input.Minat,
			Ci_UjianSekolah: input.Ci_UjianSekolah,
			Ci_RerataRaport: input.Ci_RerataRaport,
			Ci_IPA:          input.Ci_IPA,
			Ci_IPS:          input.Ci_IPS,
			Ci_Minat:        input.Ci_Minat,
		}
		db.Create(&student)
		c.JSON(http.StatusOK, gin.H{
			"message":  student,
			"response": "200"})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Failed To Create Student",
			"message":  "NISN Sudah Terdaftar",
			"response": "409"})
		return
	}

}

func PutStudent(c *gin.Context) {
	db := models.DB
	// Get model if exist
	var input, temp models.Students
	if err := c.ShouldBind(&temp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": "409",
		})
		return
	}
	if err := db.Where("id = ?", c.Param("id")).First(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Failed Update Student",
			"message":  "Data Siswa Tidak Tersedia",
			"response": "409"})
		return
	}
	newName := temp.NISN
	err := db.Where("NISN = ?", newName).First(&temp).Error
	if err != nil {
		student := models.Students{
			Nama: temp.Nama,
		}
		db.Model(&input).Updates(student)
		c.JSON(http.StatusOK, gin.H{
			"message":  input,
			"response": "200",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":  "NISN Sudah Terdaftar",
			"response": "409"})
		return
	}

}

func DeleteStudent(c *gin.Context) {
	db := models.DB
	// Get model if exist
	var input models.Students
	if err := db.Where("id = ?", c.Param("id")).First(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Failed To Delete Student",
			"message":  "Data Siswa Tidak Tersedia",
			"response": "409"})
		return
	}
	temp := input.Nama
	db.Delete(&input)

	c.JSON(http.StatusOK, gin.H{
		"data":     "Siswa Dengan Nama " + temp + " Berhasil Di Hapus",
		"response": "200",
	})
}
