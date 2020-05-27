package v1

import "github.com/gin-gonic/gin"

// ApplyRoutes is func
func ApplyRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST("/footprint", footPrint)
		v1.GET("/profile_oath", profileOath)
		v1.GET("/cat_trid", catTrid)
		v1.GET("hc", func(ctx *gin.Context) {
			name := ctx.Param("name")
			ctx.JSON(200, gin.H{
				"message": "hello" + name,
			})
		})
	}
}
