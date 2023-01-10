package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	models "spkj/Models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	db := models.DB
	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("username = ?", input.Username).First(&input).Error; err != nil {
		user := models.Users{
			Username: input.Username,
			Password: (HashPassowrd(input.Password)),
		}
		db.Create(&user)
		c.JSON(http.StatusOK, gin.H{
			"data":     user,
			"message":  "Registration Succes",
			"response": "200"})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Username Sudah Ada",
			"message":  "Gagal Registrasi",
			"response": "409"})
		return
	}
}

func HashPassowrd(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])

}

func Login(c *gin.Context) {
	db := models.DB
	var input models.Users
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := db.Where("username = ?", input.Username).Where("password = ?", HashPassowrd(input.Password)).First(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error":    "Username / Password Invalid",
			"message":  "Gagal login",
			"response": "409"})
		return
	} else {

		token, _ := createJWT(input.ID)

		c.JSON(http.StatusOK, gin.H{
			"message":  "Anda Berhasil Login",
			"token":    token,
			"response": "200"})
		return
	}
}

func GetUsers(c *gin.Context) {
	var users []models.Users
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"data":     users,
		"response": "200"})
}
