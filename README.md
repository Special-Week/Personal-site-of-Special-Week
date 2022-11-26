# 使用Golang的[Gin](https://gin-gonic.com/zh-cn/)框架搭建一个简易的个人主页



## 使用的第三方库
  
    gin         go get -u github.com/gin-gonic/gin
    sqlite3     go get -u github.com/mattn/go-sqlite3


## [成品](http://homepage.51246352.xyz/)



## 设置路由部分、端口
```golang
func main() {
	r := gin.Default()
	r.Static("/static", "./html/static")
	r.LoadHTMLGlob("html/*.html")
	r.GET("/dragon", dragonGet)               // 随机龙图
	r.GET("/dirtyLanguage", dirtyLanguage)    // 随机脏话
	r.GET("/yydz", yydzGet)                   // 随机丁真
	r.GET("/setuApi", SetuApi)                // 随机setuApi	
	r.GET("/homepage", htmlGet)               // 首页	
	r.Run(":23856")                           // run在23856端口
}

```


## 数据库及其他文件在database文件夹里面
注释已补