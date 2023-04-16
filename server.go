package main

import (
	"net/http"
	"server/db"
	"server/pj"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	IndicesData pj.IndicesData
)

func Init() {
	db.InitDB()
	IndicesData, _ = db.FINameGetAll()
}

func GetChangeMat(sd string, ed string, symbols []string, RawData map[string][]db.Record) map[string][]float32 {
	st, _ := time.Parse("2006-01-02", sd)
	et, _ := time.Parse("2006-01-02", ed)
	dur := int(et.Sub(st).Hours()/24) + 1
	changeMat := make(map[string][]float32)
	for _, symbol := range symbols {
		changeMat[symbol] = make([]float32, dur)
		for i, Record := range RawData[symbol] {
			changeMat[symbol][i] = Record.Percent
		}
	}
	return changeMat
}

func CalVar(X []float32, Y []float32) float32 {
	if len(X) != len(Y) {
		return 1000
	}
	var sx, sy, sxy float32
	n := float32(len(X))
	for i := range X {
		sx += X[i]
		sy += Y[i]
		sxy += X[i] * Y[i]
	}
	return sxy/n - sx/n*sy/n
}

func GetCorData(sd string, ed string, sym string) pj.CorData {
	var CorData pj.CorData
	CorData.Symbols = strings.Split(sym, ",")
	CorData.CorMat = make([][]float32, len(CorData.Symbols))
	for i := range CorData.CorMat {
		CorData.CorMat[i] = make([]float32, len(CorData.Symbols))
	}

	RawData, _ := db.CHRecordQuery(sd, ed, CorData.Symbols)
	changeMat := GetChangeMat(sd, ed, CorData.Symbols, RawData)

	for i, s1 := range CorData.Symbols {
		for j, s2 := range CorData.Symbols {
			CorData.CorMat[i][j] = CalVar(changeMat[s1], changeMat[s2])
			CorData.CorMat[j][i] = CorData.CorMat[i][j]
		}
	}

	return CorData
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
			c.JSON(http.StatusOK, GetCorData(sd, ed, sym))
		}
	})

	r.Run(":80")
}
