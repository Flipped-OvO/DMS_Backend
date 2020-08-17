package router

import (
	FS "backend/pkg/fs"
	"backend/util/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
)

func Initialization() *gin.Engine {
	gin.SetMode(setting.Config.Runtime.Mode)

	//创建gin访问日志文件
	//修改gin日志输出的writer
	f, _ := os.Create(setting.Config.Runtime.LogPath + "access.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.New()

	//设置gin日志格式
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage)
	}))

	router.Use(gin.Recovery())
	router.Use(corsMiddleware())
	router.Static("/assets/", "./assets/")

	router.GET("/", func(context *gin.Context) {
		context.AbortWithStatus(http.StatusForbidden)
	})
	api := router.Group("/api")
	setupRouters(api)

	return router
}

func setupRouters(r *gin.RouterGroup) {
	FS.Initialization(r)
}

// TODO 建立连接池 在得知页面有更新后，通知其他链接更新内容
//r.GET("/ws", func(ctx *gin.Context) {
//	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
//		return true
//	}}
//	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
//	if err != nil {
//		return
//	}
//	defer ws.Close()
//
//	for {
//		messageType, message, err := ws.ReadMessage()
//		if err != nil {
//			break
//		}
//		fmt.Println(string(message))
//		if string(message) == "upgrade" {
//			err = ws.WriteMessage(messageType, []byte("new data"))
//			if err != nil {
//				fmt.Println(err)
//				break
//			}
//		}
//	}
//})

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}

		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()
	}
}
