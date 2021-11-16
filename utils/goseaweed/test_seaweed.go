package goseaweed

import (
	"blog/utils"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func TestSeaweed() {
	log.SetFlags(log.Llongfile | log.LstdFlags)

	fs := NewSeaweedFs(utils.SeaweedAddress, time.Second*10)
	file, err := ioutil.ReadFile("/Users/salmon/demo/myhexo/themes/hexo-theme-matery-develop/source/medias/logo.png")
	if err != nil {
		log.Fatalln(err)
	}
	if fid,err := fs.UploadFile("submit", file); err != nil {
		log.Fatalln(err)
	}else {
		fmt.Println(fid)
	}
	//object, err := fs.GetObject("logo.png")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//ioutil.WriteFile("xxx.png",object,000666)
}