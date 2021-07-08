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


func GetFile(c *gin.Context){
	fid := c.Query("fid")
	fs := goseaweed.NewSeaweedFs("http://127.0.0.1:8080", time.Second * 10)
	content,err := fs.GetFile(fid)
	if err != nil {
		log.Fatalln(err)
	}
	//_ = ioutil.WriteFile("xxx.png",content,000666)
	code = errmsg.SUCCSE
	//c.Header("content-disposition", `attachment; filename=` + "xxx.png")
	c.Data(200, "image/png", content)
}


/*
/data/mainboard/weed server -master.port=9334 -dir=/data/2 -master.peers=172.27.156.14:93
33,172.27.156.14:9334,172.27.156.15:9333 -volume.port=9344 -ip.bind=172.27.156.14
*/