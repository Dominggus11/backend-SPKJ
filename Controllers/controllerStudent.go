package controllers

import (
	"net/http"
	models "spkj/Models"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Students
	models.DB.Find(&students)
	result := models.DB.Order("id").Find(&students)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  students,
		"response": "200"})
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
	temp_minat := strings.ToUpper(input.Minat)
	if err := db.Where("nisn = ?", input.NISN).First(&input).Error; err != nil {
		student := models.Students{
			Nama:         input.Nama,
			NISN:         input.NISN,
			UjianSekolah: input.UjianSekolah,
			RerataRaport: input.RerataRaport,
			IPA:          input.IPA,
			Minat:        temp_minat,
		}
		db.Create(&student)
		var students []models.Students
		models.DB.Find(&students)
		c.JSON(http.StatusOK, gin.H{
			"message":  students,
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
	if err := c.ShouldBindJSON(&input); err != nil { // Menggunakan ShouldBindJSON untuk membaca input dari JSON
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": "400",
		})
		return
	}
	if err := db.Where("id = ?", c.Param("id")).First(&temp).Error; err == nil {

		if input.NISN != temp.NISN { // Memeriksa apakah NISN dari input berbeda dengan NISN di basis data
			var existingStudent models.Students
			if err := db.Where("nisn = ?", input.NISN).First(&existingStudent).Error; err == nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":    "Failed Update Student",
					"message":  "NISN already exists",
					"response": "400"})
				return
			}
		}

		temp_minat := strings.ToUpper(input.Minat)
		student := models.Students{
			Nama:         input.Nama,
			NISN:         input.NISN,
			UjianSekolah: input.UjianSekolah,
			RerataRaport: input.RerataRaport,
			IPA:          input.IPA,
			Minat:        temp_minat,
		}
		db.Model(&temp).Updates(student)
		var students []models.Students
		db.Find(&students) // Menggunakan db.Find untuk membaca data dari basis data
		c.JSON(http.StatusOK, gin.H{
			"message":  students,
			"response": "200"})
		return

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Failed Update Student",
			"message":  "Data siswa tidak tersedia",
			"response": "400"})
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
	var students []models.Students
	models.DB.Find(&students)
	c.JSON(http.StatusOK, gin.H{
		"message":  students,
		"data":     "Siswa dengan nama " + temp + " berhasil di hapus",
		"response": "200",
	})
}

func ConversiNilai(nilai float64) float64 {
	var ci float64
	if nilai >= 90 {
		ci = 5
	} else if nilai >= 80 {
		ci = 4
	} else if nilai >= 70 {
		ci = 3
	} else if nilai >= 60 {
		ci = 2
	} else {
		ci = 1
	}
	return ci
}

func GetCi(c *gin.Context) {
	// get data siswas
	siswas := GetDataSiswa()
	// before normalisasi
	BeforeNormalisasi(c, siswas)
	// normalisasi data siswas
	siswas = GetDataSiswa()
	Normalisasi(c, siswas)

	siswas = GetDataSiswa()
	NormalisasiSMART(c, siswas)
	// get kriterias
	criterias := GetDataKriteria()
	// result dari SAW
	siswas = GetDataSiswa()
	ResultSAW(c, criterias, siswas)
	siswas = GetDataSiswa()
	ResultSMART(c, criterias, siswas)
	var students []models.Students
	models.DB.Find(&students)

	result := models.DB.Order("id").Find(&students)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  students,
		"response": "200"})

}
