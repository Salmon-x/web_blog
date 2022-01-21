package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type ResponseAll struct {
	Response
	Total int64 `json:"total"`
}

func Result(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func ResultAll(code int, msg string, data interface{}, total int64, c *gin.Context)  {
	result := ResponseAll{}
	result.Code = code
	result.Msg = msg
	result.Data = data
	result.Total = total
	c.JSON(http.StatusOK, result)
}
