package controllers

import (
	"fmt"
	"net/http"
	models "spkj/Models"

	"github.com/gin-gonic/gin"
)

func Developer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Developer": "Roy Dominggus Andornov Malau",
		"Project":   "Sistem Pendukung Keputusan Jurusan",
		"Version":   "1.0",
	})
}

func GetCriterias(c *gin.Context) {
	var criterias []models.Criterias
	models.DB.Find(&criterias)

	c.JSON(http.StatusOK, gin.H{
		"message":  criterias,
		"response": "200"})
}

func GetCriteria(c *gin.Context) {
	db := models.DB
	// Get model if exist
	var criteria models.Criterias
	if err := db.Where("id = ?", c.Param("id")).First(&criteria).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Kriteria Tidak Ditemukan",
			"response": "409"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": criteria})
}

func PostCriteria(c *gin.Context) {

	db := models.DB
	var input models.Criterias
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": "409",
		})
		return
	}
	if err := db.Where("nama_kriteria = ?", input.NamaKriteria).First(&input).Error; err != nil {
		criteria := models.Criterias{
			NamaKriteria:  input.NamaKriteria,
			BobotKriteria: input.BobotKriteria,
		}
		db.Create(&criteria)
		c.JSON(http.StatusOK, gin.H{
			"data":     criteria,
			"response": "200"})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Jenis Kriteria Sudah Ada",
			"response": "409"})
		return
	}

}

func PutCriteria(c *gin.Context) {
	db := models.DB
	// Get model if exist
	var input, temp models.Criterias
	if err := c.ShouldBind(&temp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": "409",
		})
		return
	}
	if err := db.Where("id = ?", c.Param("id")).First(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Kriteria Tidak Ditemukan",
			"response": "409"})
		return
	}
	// newName := temp.NamaKriteria
	// err := db.Where("nama_kriteria = ?", newName).First(&temp).Error
	// if err != nil {
	criteria := models.Criterias{
		BobotKriteria: temp.BobotKriteria,
	}
	db.Model(&input).Updates(criteria)
	c.JSON(http.StatusOK, gin.H{
		"data":     input,
		"response": "200",
	})
	// } else {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error":    "Kriteria Sudah Ada",
	// 		"response": "409",
	// 	})
	// 	return
	// }

}

func PutCriteriaIsactive(c *gin.Context) {
	db := models.DB
	// Get model if exist
	var input, temp models.Criterias
	if err := c.ShouldBind(&temp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"response": "409",
		})
		return
	}
	if err := db.Where("id = ?", c.Param("id")).First(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Kriteria Tidak Ditemukan",
			"response": "409"})
		return
	}
	criteria := models.Criterias{
		Is_active: temp.Is_active,
	}
	fmt.Println(temp.Is_active)
	db.Model(&input).Updates(criteria)
	c.JSON(http.StatusOK, gin.H{
		"data":     input,
		"response": "200",
	})
}

func DeleteCriteria(c *gin.Context) {
	db := models.DB
	// Get model if exist
	var input models.Criterias
	if err := db.Where("id = ?", c.Param("id")).First(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Kriteria Tidak Ditemukan",
			"response": "409"})
		return
	}
	temp := input.NamaKriteria
	db.Delete(&input)

	c.JSON(http.StatusOK, gin.H{
		"data":     "Kriteria " + temp + " Berhasil Di Hapus",
		"response": "200",
	})
}
