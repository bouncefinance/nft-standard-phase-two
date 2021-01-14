package util

import (
	"github.com/gin-gonic/gin"
	"nft_standard/msgs"
)

type Context struct {
	C *gin.Context
}

func (c *Context) Response(httpCode int, code int, data interface{}) {
	c.C.Header("Access-Control-Allow-Origin","*")
	c.C.Header("Access-Control-Allow-Methods","*")
	c.C.Header("Access-Control-Allow-Headers","*")

	c.C.JSON(httpCode, gin.H{
		"code":code,
		"error_msg":  msgs.MsgReturn[code],
		"data": data,
	})
	return
}

