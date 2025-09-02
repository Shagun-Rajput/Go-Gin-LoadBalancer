package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleRequest(c *gin.Context) {
	// Your request handling logic here
	c.JSON(http.StatusOK, gin.H{"message": "Request handled successfully"})
}
