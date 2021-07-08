package test
//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"io/ioutil"
//	"log"
//	"net/http"
//)
//
//func UtilGetFile(objectName string) ([]byte, error) {
//	body, err := http.Get(fmt.Sprintf("%s/%s", "http://127.0.0.1:8080", objectName))
//	if err != nil {
//		return nil, nil
//	}
//	defer body.Body.Close()
//	return ioutil.ReadAll(body.Body)
//}
//
//func GetFile(c *gin.Context){
//	fid,_ := c.Params.Get("fid")
//	content,err := UtilGetFile(fid)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	c.Data(200, "image/png", content)
//}
//
//func main() {
//	// 1.创建路由
//	r := gin.Default()
//	gin.SetMode("release")
//	r.GET("/:fid/", GetFile)
//	// Run("里面不指定端口号默认为8080")
//	_ = r.Run(":8020")
//}

