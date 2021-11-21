package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"net/http"
	"path/filepath"
)

type File struct{}

func (h File) SaveFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	extension := filepath.Ext(file.Filename)
	fileId, _ := uuid.New()

	newFileName := fmt.Sprintf("%x", fileId) + extension

	if err := c.SaveUploadedFile(file, "./static/"+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
	return
}
