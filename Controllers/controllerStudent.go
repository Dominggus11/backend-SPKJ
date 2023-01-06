package controllers

import (
	"fmt"
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
	} else {
		input.Ci_UjianSekolah, input.Ci_RerataRaport, input.Ci_IPA, input.Ci_IPS, input.Ci_Minat = BeforeNormalisasi(temp.UjianSekolah, temp.RerataRaport, temp.IPA, temp.IPS, temp.Minat)
		student := models.Students{
			Nama:            input.Nama,
			NISN:            input.NISN,
			UjianSekolah:    temp.UjianSekolah,
			RerataRaport:    temp.RerataRaport,
			IPA:             temp.IPA,
			IPS:             temp.IPS,
			Minat:           temp.Minat,
			Ci_UjianSekolah: input.Ci_UjianSekolah,
			Ci_RerataRaport: input.Ci_RerataRaport,
			Ci_IPA:          input.Ci_IPA,
			Ci_IPS:          input.Ci_IPS,
			Ci_Minat:        input.Ci_Minat,
		}
		fmt.Println(input.Nama, input.NISN, temp.UjianSekolah, temp.RerataRaport, temp.IPA, temp.IPS, temp.Minat)
		db.Model(&input).Updates(student)
		c.JSON(http.StatusOK, gin.H{
			"message":  student,
			"response": "200"})
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

func ConversiNilai(nilai float64) float64 {
	var ci float64
	if nilai > 91 {
		ci = 5
	} else if nilai > 80 {
		ci = 4
	} else if nilai > 70 {
		ci = 3
	} else if nilai > 60 {
		ci = 2
	} else {
		ci = 1
	}
	return ci
}

func ConversiJurusan(jurusan string) float64 {
	var ci float64
	if jurusan == "IPA" {
		ci = 5
	} else {
		ci = 1
	}
	return ci
}

func GetCi(c *gin.Context) {
	// get data siswas
	siswas := GetDataSiswa()

	// normalisasi data siswas
	Normalisasi(c, siswas)

	// get kriterias
	criterias := GetDataKriteria()

	// result dari SAW
	ResultSAW(c, criterias, siswas)

	var students []models.Students
	models.DB.Find(&students)

	c.JSON(http.StatusOK, gin.H{
		"message":  students,
		"response": "200"})

}
