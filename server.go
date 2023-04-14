package main

import (
	"net/http"
	"server/db"
	"server/pj"

	"github.com/gin-gonic/gin"
)

var (
	IndicesData pj.IndicesData
)

func Init() {
	db.InitDB()
	IndicesData, _ = db.FINameGetAll()
}

func main() {
	Init()

	r := gin.Default()
	r.LoadHTMLFiles("html/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/QueryIndices", func(c *gin.Context) {
		c.JSON(http.StatusOK, IndicesData)
	})
	r.Run(":80")
}
