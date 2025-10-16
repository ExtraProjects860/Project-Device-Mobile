package middleware

import "github.com/gin-gonic/gin"

func SecurityHeaders(router *gin.Engine) {
	router.Use(
func(c *gin.Context) {
		c.Header("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Referrer-Policy", "no-referrer")
		c.Next()
		},
	)
}
