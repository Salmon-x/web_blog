package middleware

import (
	"blog/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {

	filePath := utils.LogPath + "blog.log"
	// 软连接路径(可以根据运维来进行调整)
	// linkName := "latest_log.log"
	// 初始化
	logger := logrus.New()
	// 定位到log文件夹(路径，os新建文件，赋值权限)
	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	}
	// 输出到日志文件中
	logger.Out = src

	logger.SetLevel(logrus.DebugLevel)
	logWriter, _ := retalog.New(
		utils.LogPath+"%Y%m%d.log",
		retalog.WithMaxAge(7*24*time.Hour),
		retalog.WithRotationTime(24*time.Hour),
		// retalog.WithLinkName(linkName),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(Hook)

	return func(c *gin.Context) {
		// 输出开始时间
		startTime := time.Now()
		c.Next()
		// 结束时间运行时间
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		// 主机名
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unKnown"
		}
		// 状态码
		statusCode := c.Writer.Status()
		// 客户端ip
		clientIP := c.ClientIP()
		// 客户端信息
		userAgent := c.Request.UserAgent()
		// 请求大小
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		// 请求方法和路径
		method := c.Request.Method
		path := c.Request.RequestURI
		// 组合数据
		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"Status":    statusCode,
			"SpendTime": spendTime,
			"IP":        clientIP,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		// 分级别
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}

}
