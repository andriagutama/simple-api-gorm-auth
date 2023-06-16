package auth

import (
	"net/http"
	"simple-api-gorm-auth/models"
	"time"

	"github.com/gin-gonic/gin"

	jwt "github.com/golang-jwt/jwt/v4"
)

const (
	USER     = "admin"
	PASSWORD = "password123"
	SECRET   = "secret"
)

func LoginHandler(c *gin.Context) {
	var user models.Credential
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	if user.Username != USER {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized. invalid user",
		})
		return
	} else {
		if user.Password != PASSWORD {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "anauthorized. invalid password",
			})
			return
		}
	}

	claim := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		Issuer:    "test",
		IssuedAt:  time.Now().Unix(),
	}

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := sign.SignedString([]byte(SECRET))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"token":   token,
	})
}
