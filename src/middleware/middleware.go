package middleware

import (
	"net/http"

	"github.com/WilhelmWeber/iiif_annotation_server_v2/src/libs"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.JSON((http.StatusUnauthorized), gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}

	_, notOk := libs.ParseToken(tokenString)
	if notOk != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		c.Abort()
		return
	}

	c.Next()
}
