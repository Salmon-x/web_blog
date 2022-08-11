package v1

import (
	"blog/model/response"
	"blog/utils"
	"blog/utils/errmsg"
	"blog/utils/goseaweed"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"time"
)

type FileApi struct {
}

// 上传文件

func (f *FileApi) UploadFile(c *gin.Context) {
	var code int
	file, _, _ := c.Request.FormFile("file")
	// 初始化
	fs := goseaweed.NewSeaweedFs(utils.SeaweedAddress, time.Second*10)
	// 转换数据类型为bytes
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		code = errmsg.ERROR
	} else {
		code = errmsg.SUCCSE
	}
	fid, err := fs.UploadFile("submit", buf.Bytes())
	// 上传文件
	if err != nil {
		log.Fatalln(err)
	}
	response.Result(code, fid, c)
}

func (f *FileApi) GetFile(c *gin.Context) {
	fid := c.Query("fid")
	fs := goseaweed.NewSeaweedFs("http://127.0.0.1:8080", time.Second*10)
	content, err := fs.GetFile(fid)
	if err != nil {
		log.Fatalln(err)
	}
	c.Data(200, "image/png", content)
}

// 删除文件
func (f *FileApi) DelFile(c *gin.Context) {
	fid := c.Query("fid")
	fs := goseaweed.NewSeaweedFs("http://127.0.0.1:8080", time.Second*10)
	err, code := fs.RemoveFile(fid)
	if err != nil {
		utils.Logger(err)
	}
	response.Result(code, "", c)
}
