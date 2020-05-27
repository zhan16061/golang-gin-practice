package v1

import (
	"github.com/gin-gonic/gin"
)

func catTrid(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"cat_trid": "qaz123-wsx123-edc123",
		"pt_jwt":   "qwertyuiolkjhgfdsazxcvbnm",
	})
}

func profileOath(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"qwe": "qwertyuioplkjhgfdsazxcvbnm",
	})
}

func footPrint(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"page_id": "qwertyuioplkjhgfdsazxcvbnm",
	})
}
