package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	AppMode        string
	HttpPort       string
	JwtKey         string
	SeaweedAddress string
	UseMultipoint  bool
	TransferTable  bool
	LogPath        string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	RedisHost     string
	RedisPort     string
	RedisPassword string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查路径", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadRedis(file)

}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8020")
	JwtKey = file.Section("server").Key("JwtKey").MustString("+d0(4=-$hy(cdl$tu^@509r#t$e60-1&v$09kb^tife&gfcfmd")
	SeaweedAddress = file.Section("server").Key("SeaweedAddress").MustString("127.0.0.1:9333")
	UseMultipoint = file.Section("server").Key("UseMultipoint").MustBool(false)
	TransferTable = file.Section("server").Key("TransferTable").MustBool(false)
	LogPath = file.Section("server").Key("LogPath").MustString("log/")
	if !strings.HasSuffix(LogPath, "/") {
		LogPath += "/"
	}
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("")
	DbName = file.Section("database").Key("DbName").MustString("blogdb")
}

// redis配置
func LoadRedis(file *ini.File) {
	RedisHost = file.Section("redis").Key("RedisHost").MustString("localhost")
	RedisPort = file.Section("redis").Key("RedisPort").MustString("6379")
	RedisPassword = file.Section("redis").Key("RedisPassword").MustString("")
}
