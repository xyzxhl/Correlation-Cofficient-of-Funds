package main

import (
	"net/http"
	"server/db"
	"server/pj"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	IndicesData pj.IndicesData
)

func Init() {
	db.InitDB()
	IndicesData, _ = db.FINameGetAll()
}

func GetChanges(sd string, ed string, sym string) [][]interface{} {
	symbols := strings.Split(sym, ",")
	RawData, _ := db.CHRecordQuery(sd, ed, symbols)
	return RawData
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

	r.GET("/QueryChanges", func(c *gin.Context) {
		sd, ok1 := c.GetQuery("sd")
		ed, ok2 := c.GetQuery("ed")
		sym, ok3 := c.GetQuery("sym")
		if ok1 && ok2 && ok3 {
			c.JSON(http.StatusOK, GetChanges(sd, ed, sym))
		}
	})

	r.Run(":80")
}
