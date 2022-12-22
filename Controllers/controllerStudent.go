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
		"data": students})
}

func GetStudent(c *gin.Context) {
	db := models.DB
	// Get model if exist
	var student models.Students
	if err := db.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Siswa Dengan nama" + student.Nama + " tidak Ada"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": student})
}

func PostStudent(c *gin.Context) {

	db := models.DB
	var input models.Students
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := db.Where("nama = ?", input.Nama).First(&input).Error; err != nil {
		student := models.Students{
			Nama: input.Nama,
		}
		db.Create(&student)
		c.JSON(http.StatusOK, gin.H{
			"data": student})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama Siswa Tersebut Sudah Ada !!!"})
		return
	}

}

func PutStudent(c *gin.Context) {
	db := models.DB
	// Get model if exist
	var input, temp models.Students
	if err := c.ShouldBind(&temp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := db.Where("id = ?", c.Param("id")).First(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Siswa Dengan nama" + input.Nama + " tidak Ada"})
		return
	}
	newName := temp.Nama
	err := db.Where("nama = ?", newName).First(&temp).Error
	if err != nil {
		student := models.Students{
			Nama: temp.Nama,
		}
		db.Model(&input).Updates(student)
		c.JSON(http.StatusOK, gin.H{
			"data": input,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama Siswa Sudah Ada !!!"})
		return
	}

}

func DeleteStudent(c *gin.Context) {
	db := models.DB
	// Get model if exist
	var input models.Students
	if err := db.Where("id = ?", c.Param("id")).First(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Siswa dengan nama " + input.Nama + " Tidak Ada !"})
		return
	}
	temp := input.Nama
	db.Delete(&input)

	c.JSON(http.StatusOK, gin.H{
		"data": "Siswa Dengan Nama " + temp + " Berhasil Di Hapus !!!",
	})
}
