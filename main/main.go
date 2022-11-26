package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var yydzPath = "./database/yydz"
var dargonPath = "./database/Dragon"
var yydzAllFileName []interface{}
var dargonAllFileName []interface{}

// 开始前获取所有文件名
func init() {
	filepath.Walk(yydzPath, func(_ string, info os.FileInfo, _ error) error {
		if info.IsDir() {
			return nil
		}
		yydzAllFileName = append(yydzAllFileName, info.Name())
		return nil
	})
	filepath.Walk(dargonPath, func(_ string, info os.FileInfo, _ error) error {
		if info.IsDir() {
			return nil
		}
		dargonAllFileName = append(dargonAllFileName, info.Name())
		return nil
	})
}

// main函数
func main() {
	// 路由
	r := gin.Default()
	r.Static("/static", "./html/static")
	r.LoadHTMLGlob("html/*.html")
	// 随机龙图
	r.GET("/dragon", dragonGet)
	// 随机脏话
	r.GET("/dirtyLanguage", dirtyLanguage)
	// 随机丁真
	r.GET("/yydz", yydzGet)
	// 随机setuApi
	r.GET("/setuApi", SetuApi)
	// 首页
	r.GET("/homepage", htmlGet)
	// run在23856端口
	r.Run(":23856")

}

func dirtyLanguage(c *gin.Context) {
	c.PureJSON(http.StatusOK, selectDirtyLanguage())
}

func yydzGet(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.File(yydzPath + "\\" + fmt.Sprintf("%v", randomChoiceItem(yydzAllFileName)))
}

func dragonGet(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.File(dargonPath + "\\" + fmt.Sprintf("%v", randomChoiceItem(dargonAllFileName)))
}

func htmlGet(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func SetuApi(c *gin.Context) {
	keyword := c.DefaultQuery("tag", "")
	argum := c.DefaultQuery("num", "1")
	argr18 := c.DefaultQuery("r18", "false")
	argumInt, err := strconv.Atoi(argum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "num参数错误",
		})
		return
	}
	// num最大1000
	if argumInt > 1000 {
		argumInt = 1000
	}
	argr18Bool, err := strconv.ParseBool(argr18)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "r18参数错误",
		})
		return
	}
	c.PureJSON(http.StatusOK, selectSetu(keyword, argumInt, argr18Bool))
}
