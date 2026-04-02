package middleware

import (
	"bytes"
	"context"
	"gin-generate-framework/core/global"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader("X-Trace-ID")
		if traceID == "" {
			traceID = uuid.New().String()
		}
		c.Set("trace_id", traceID)

		// 关键：将 trace_id 注入到 request.Context
		ctx := context.WithValue(c.Request.Context(), "trace_id", traceID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func AccessLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		method := c.Request.Method

		// 读取请求 Body（用于记录，注意需要重新写入）
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 自定义响应写入器，捕获响应 Body
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 处理请求
		c.Next()

		// 记录访问日志
		latency := time.Since(start)
		statusCode := c.Writer.Status()

		// 从 Context 获取 trace_id（由其他中间件注入）
		traceID, _ := c.Get("trace_id")

		global.AccessLog.WithFields(map[string]interface{}{
			"trace_id":      traceID,
			"status":        statusCode,
			"method":        method,
			"path":          path,
			"query":         query,
			"ip":            c.ClientIP(),
			"latency_ms":    latency.Milliseconds(),
			"user_agent":    c.Request.UserAgent(),
			"request_body":  string(requestBody),
			"response_body": blw.body.String(),
			"errors":        c.Errors.String(),
		}).Info("access_log")
	}
}

// 自定义 ResponseWriter，用于捕获响应 Body
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
