package v1

import (
	"blog/model/response"
	"blog/utils/errmsg"
	"blog/utils/goseaweed"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type FileApi struct {

}


// 上传文件
//func (f *FileApi)UploadFile(c *gin.Context){
//	file, _, _ := c.Request.FormFile("file")
//	// 初始化
//	fs := goseaweed.NewSeaweedFs(utils.SeaweedAddress, time.Second * 10)
//	code = errmsg.SUCCSE
//	// 转换数据类型为bytes
//	buf := bytes.NewBuffer(nil)
//	if _, err := io.Copy(buf, file); err != nil {
//		code = errmsg.ERROR
//	}
//	fid,err := fs.UploadFile("submit",buf.Bytes())
//	// 上传文件
//	if  err != nil {
//		log.Fatalln(err)
//	}
//	c.JSON(http.StatusOK,gin.H{
//		"code":code,
//		"msg": errmsg.GetErrorMsg(code),
//		"data": fid,
//	})
//}

func (f *FileApi)UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "static/"+file.Filename)
	if err != nil{
		code = errmsg.ERROR
	}
	// 上传文件
	response.Result(code, errmsg.GetErrorMsg(code), file.Filename, c)
}


func (f *FileApi)GetFile(c *gin.Context){
	fid := c.Query("fid")
	fs := goseaweed.NewSeaweedFs("http://127.0.0.1:8081", time.Second * 10)
	content,err := fs.GetFile(fid)
	if err != nil {
		log.Fatalln(err)
	}
	//_ = ioutil.WriteFile("xxx.png",content,000666)
	code = errmsg.SUCCSE
	//c.Header("content-disposition", `attachment; filename=` + "xxx.png")
	c.Data(200, "image/png", content)
}

// 删除文件
func (f *FileApi)DelFile(c *gin.Context){
	fid := c.Query("fid")
	fs := goseaweed.NewSeaweedFs("http://127.0.0.1:8081", time.Second * 10)
	err,code := fs.RemoveFile(fid)
	if err != nil {
		fmt.Println(err)
	}
	response.Result(code, errmsg.GetErrorMsg(code), "", c)
}

/*
/data/mainboard/weed server -master.port=9334 -dir=/data/2 -master.peers=172.27.156.14:93
33,172.27.156.14:9334,172.27.156.15:9333 -volume.port=9344 -ip.bind=172.27.156.14
*/