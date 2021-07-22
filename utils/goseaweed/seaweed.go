package goseaweed

import (
	"blog/utils/errmsg"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

type seaweedFS struct {
	serverURL string
	timeout   time.Duration
}

type Res struct {
	Fid string `json:"fid"`
}


func NewSeaweedFs(serverURL string, timeout time.Duration) Seaweed {
	if timeout <= 100 {
		timeout = time.Second * 10
	}
	return &seaweedFS{serverURL: serverURL, timeout: timeout}
}


func (s *seaweedFS) UploadFile(objectName string, content []byte) (string, error) {
	// 拼接url
	url := fmt.Sprintf("%s/%s", s.serverURL, objectName)
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	err := writer.WriteField("file", bytes.NewBuffer(content).String())
	if err != nil {
		return "",errors.WithStack(err)
	}
	err = writer.Close()
	if err != nil {
		return "",errors.WithStack(err)
	}
	client := &http.Client{Timeout: s.timeout}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return "",errors.WithStack(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	// 返回值
	res, _ := client.Do(req)
	// 将数据转换为json
	body, err := ioutil.ReadAll(res.Body)
	result := string(body)
	var fid Res
	_ = json.Unmarshal([]byte(result),&fid)
	if err != nil {
		return "",errors.WithStack(err)
	}
	if res.StatusCode != http.StatusCreated {
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "",errors.WithStack(err)
		}
		return "",fmt.Errorf("create failed, response:%s", string(body))
	}
	return fid.Fid,nil
}

func (s *seaweedFS) GetFile(objectName string) ([]byte, error) {
	body, err := http.Get(fmt.Sprintf("%s/%s", s.serverURL, objectName))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer body.Body.Close()
	return ioutil.ReadAll(body.Body)
}

func (s *seaweedFS) RemoveFile(objectName string) (error,int) {
	url := fmt.Sprintf("%s/%s", s.serverURL, objectName)
	method := "DELETE"
	client := &http.Client{Timeout: s.timeout}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return errors.WithStack(err),errmsg.ERROR
	}
	res, err := client.Do(req)
	if err != nil {
		return errors.WithStack(err),errmsg.ERROR
	}
	if res.StatusCode != http.StatusNoContent {
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return errors.WithStack(err),errmsg.ERROR
		}
		size := strings.SplitN(string(body), ":", 2)[1]

		if size[:len(size)-1] != "0" {
			return nil, errmsg.SUCCSE
		}
		return fmt.Errorf("delete failed, response:%s", string(body)),errmsg.ERROR_FILE_EXIST
	}
	return nil,errmsg.SUCCSE
}
