package middleware

import (
	"codly/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"strings"
	"time"

	"github.com/valyala/fasttemplate"
)

// LoggerConfig defines config of logging middleware
type (
	LoggerConfig struct {
		Format   string
		Level    uint8
		IsLogger bool
	}
)

// DefaultLoggerConfig is default config of logging middleware
var DefaultLoggerConfig = LoggerConfig{
	Format:   `${status} | ${request_time} | ${remote_ip} | ${method} "${url}"`,
	Level:    2,
	IsLogger: true,
}

// LoggerMiddleware is default implementation of logging middleware
func LoggerMiddleware() gin.HandlerFunc {
	return LoggerWithConfig(DefaultLoggerConfig)
}

// LoggerWithConfig is custom implementation of logging middleware
func LoggerWithConfig(config LoggerConfig) gin.HandlerFunc {
	if config.Format == "" {
		config.Format = DefaultLoggerConfig.Format
	}

	if config.Level == 0 {
		config.Level = DefaultLoggerConfig.Level
	}

	t, err := fasttemplate.NewTemplate(config.Format, "${", "}")

	if err != nil {
		log.Fatalf("unexpected error when parsing template: %s", err)
	}

	loggerInstance := logging.DefaultLogger

	loggerInstance.SetLevel(config.Level)

	loggerInstance.IsLogger(config.IsLogger)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		requestTime := time.Since(startTime)
		s := t.ExecuteFuncString(func(w io.Writer, tag string) (int, error) {
			switch tag {
			case "url":
				p := c.Request.URL.Path
				if p == "" {
					p = "/"
				}
				return w.Write([]byte(p))
			case "method":
				return w.Write([]byte(strings.ToUpper(c.Request.Method)))
			case "status":
				return w.Write([]byte(fmt.Sprintf("%d", c.Writer.Status())))
			case "remote_ip":
				return w.Write([]byte(fmt.Sprintf("%15s", c.RemoteIP())))
			case "host":
				return w.Write([]byte(c.Request.Host))
			case "protocol":
				return w.Write([]byte(c.Request.Proto))
			case "bytes_in":
				cl := c.Request.Header.Get("Content-length")
				if cl == "" {
					cl = "0"
				}
				return w.Write([]byte(cl))
			case "bytes_out":
				return w.Write([]byte(fmt.Sprintf("%d", c.Writer.Size())))
			case "request_time":
				return w.Write([]byte(fmt.Sprintf("%8dms", requestTime.Milliseconds())))
			default:
				return w.Write([]byte(""))
			}
		})
		if c.Writer.Status() < 400 {
			loggerInstance.Info(s)
		} else {
			loggerInstance.Warn(s)
		}
	}
}
