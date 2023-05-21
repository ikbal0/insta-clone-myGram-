package utils

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func DeleteTempFile(path string, ctx *gin.Context) {
	if errRemove := os.Remove(path); errRemove != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   errRemove.Error(),
			"message": "file has not found!",
		})
		return
	}
}
