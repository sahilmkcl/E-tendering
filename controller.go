package main

import (
	"E-Tendering/middleware"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.POST("/excel", func(c *gin.Context) {
		f, err := c.FormFile("file")
		if err != nil {
			log.Println(err)
		}
		extension := filepath.Ext(f.Filename)

		newFileName := uuid.New().String() + extension
		if err := c.SaveUploadedFile(f, "./excelfiles/"+newFileName); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
				"error":   err.Error(),
			})
			return
		}
		xlsx, err := excelize.OpenFile("./excelfiles/" + newFileName)
		if err != nil {
			fmt.Println(err.Error())
			print("[file] failed to read")
			return
		}
		for _, row := range xlsx.GetRows("Prop add") {
			log.Printf("%v", row)
		}
	})
	return r
}
