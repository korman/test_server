package handler

import (
	"comic_server/utils"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"log"
	"net/http"
)

func Index(c *gin.Context) {
	fmt.Print(c.Request.Header)

	c.JSON(http.StatusOK, gin.H{
		"msg": "Hello World",
	})
}

func BookList(c *gin.Context) {
	pbList := utils.Instance().GetBookList()

	data, err := proto.Marshal(pbList)

	if nil != err {
		log.Fatal("Marshal booklist data error: \n",err)
		return
	}

	base64Data := base64.StdEncoding.EncodeToString(data)

	c.JSON(http.StatusOK, gin.H{
		"pb_data": base64Data,
	})
}
