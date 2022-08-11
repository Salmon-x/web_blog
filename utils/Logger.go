package utils

import (
	"fmt"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func Logger(error error) {
	// 定义文件路径
	filePath := LogPath + "error.log"
	logger := logrus.New()
	// 定位到log文件夹(路径，os新建文件，赋值权限)
	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0775)
	if err != nil {
		fmt.Println(err)
	}
	logger.Out = src
	logWriter, _ := retalog.New(
		"log/%Y%m%d.error",
		retalog.WithMaxAge(7*24*time.Hour),
		retalog.WithRotationTime(24*time.Hour),
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
	entry := logger.WithFields(logrus.Fields{
		"error": error,
	})
	entry.Error()
}
