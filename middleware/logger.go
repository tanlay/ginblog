package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

//Logger 自定义日志格式，输出到文件
func Logger() gin.HandlerFunc {
	filePath := "log/ginblog"
	//linkName := "latest_log.log"
	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE,0755)
	if err != nil {
		fmt.Println("err：", err.Error())
	}
	logger := logrus.New()

	//不记录在控制台
	logger.Out = src

	logger.SetLevel(logrus.DebugLevel)
	logWriter, _ := rotalog.New(
		filePath+"%Y%m%d.log",
		rotalog.WithMaxAge(7*24*time.Hour),
		rotalog.WithRotationTime(24*time.Hour),
		//rotalog.WithLinkName(linkName),
		)

		writeMap := lfshook.WriterMap{
			logrus.InfoLevel: logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel: logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}

		hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
		logger.AddHook(hook)
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		costTime := fmt.Sprintf("%d", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unKnown"
		}
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName": hostName,
			"status": statusCode,
			"CostTime": costTime,
			"IP": clientIP,
			"Method": method,
			"Path": path,
			"DataSize": dataSize,
			"UserAgent": userAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500{
			entry.Error()
		} else if  statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}

}
