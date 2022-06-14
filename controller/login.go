package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"project.com/model"
)

type login struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func Utama(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "selamat datang di aplikasi heroku",
	})
}

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var login login
	var user model.User
	db.Where("email = ?", login.Email).Where("password = ?", login.Password).Find(&user)
	if user.Email != login.Email {
		c.JSON(400, gin.H{
			"status": "email atau password salah",
		})
	}
	if user.Level == 1 {
		c.JSON(200, gin.H{
			"status": "berhasil",
			"level":  "admin",
			"email":  user.Email,
		})
	}
	if user.Level == 2 {
		c.JSON(200, gin.H{
			"status": "berhasil",
			"level":  "dokter",
			"email":  user.Email,
		})
	}
	if user.Level == 3 {
		c.JSON(200, gin.H{
			"status": "berhasil",
			"level":  "perawat",
			"email":  user.Email,
		})
	}
}
