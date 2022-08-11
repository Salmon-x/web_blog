package response

import (
	"blog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
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

func Result(code int, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		errmsg.GetErrorMsg(code),
	})
}

func ResultAll(code int, data interface{}, total int64, c *gin.Context) {
	result := ResponseAll{}
	result.Code = code
	result.Msg = errmsg.GetErrorMsg(code)
	result.Data = data
	result.Total = total
	c.JSON(http.StatusOK, result)
}
