package phase1_work

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"phase1/phase1_work/handler"
	"phase1/phase1_work/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func InitGinRouting() {
	// 初始化 Gin
	r := gin.Default()
	r.Use(loggerMiddleware())
	// 初始化GIN路由
	apiV1 := r.Group("/api/v1")
	{
		// 无需登录
		noLogin := apiV1.Group("/no_login")
		{
			// 登录注册
			noLogin.POST("/register", handler.Register)
			noLogin.POST("/login", handler.Login)

			// 文章
			noLogin.POST("/post/list", handler.ListPost)
			noLogin.POST("/post/detail", handler.GetPostDetail)

			// 评论
			noLogin.POST("/comment/list_by_post_id", handler.ListCommentByPostId)
		}

		// 需要登录
		login := apiV1.Group("/login")
		login.Use(authMiddleware())
		{
			// 文章
			login.POST("/post/create", handler.CreatePost)
			login.POST("/post/update", handler.UpdatePost)
			login.POST("/post/delete", handler.DeletePost)

			// 评论
			login.POST("/comment/create", handler.CreateComment)
		}
	}

	// 启动服务
	_ = r.Run(":8080")

}

// responseWriter 用于捕获响应数据
type responseWriter struct {
	gin.ResponseWriter
	body []byte
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.body = append(w.body, b...)
	return w.ResponseWriter.Write(b)
}

func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 前置处理
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// 读取请求参数
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = c.GetRawData()
			// 重新设置请求体，因为 GetRawData 会消费掉请求体
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 捕获响应数据
		rw := &responseWriter{ResponseWriter: c.Writer}
		c.Writer = rw

		// 进入下一个处理函数
		c.Next()

		// 后置处理
		latency := time.Since(start)
		status := c.Writer.Status()

		// 打印请求信息
		log.Printf("[API] %s %s", method, path)
		log.Printf("[Request] %s", string(requestBody))
		log.Printf("[Response] Status: %d, Body: %s", status, string(rw.body))
		log.Printf("[Latency] %v\n", latency)
	}
}

// ========== 认证中间件 ==========
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Header 获取 Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Authorization header required",
			})
			c.Abort()
			return
		}

		// 提取 Token（Bearer <token>）
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Invalid authorization header format",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 验证 Token
		userInfo, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Invalid token",
			})
			c.Abort()
			return
		}

		// 将用户信息存储到 Context
		c.Set(utils.UserId, userInfo.UserID)

		c.Next()
	}
}
