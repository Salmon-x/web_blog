package v1

import (
	"blog/utils"
	"blog/utils/errmsg"
	"blog/utils/goseaweed"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"time"
)

// 上传文件
func UploadFile(c *gin.Context){
	file, _, _ := c.Request.FormFile("file")
	// 初始化
	fs := goseaweed.NewSeaweedFs(utils.SeaweedAddress, time.Second * 10)
	code = errmsg.SUCCSE
	// 转换数据类型为bytes
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		code = errmsg.ERROR
	}
	fid,err := fs.UploadFile("submit",buf.Bytes())
	// 上传文件
	if  err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg": errmsg.GetErrorMsg(code),
		"data": fid,
	})
}

