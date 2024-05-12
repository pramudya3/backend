package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:3000"},
			AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
			AllowHeaders: append([]string{"content-type"},
				supertokens.GetAllCORSHeaders()...),
			AllowCredentials: true,
		})
	}
}
