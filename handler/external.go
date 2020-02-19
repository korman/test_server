package handler

import (
	"comic_server/utils"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"log"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	fmt.Print(c.Request.Header)

	c.JSON(http.StatusOK, gin.H{
		"msg": "Hello World",
	})
}

func BookList(c *gin.Context) {
	msgId,err := strconv.Atoi(c.Query("msgId"))

	if nil != err {
		log.Fatal(err.Error())
		return
	}

	pbList := utils.Instance().GetBookList()

	data, err := proto.Marshal(pbList)

	if nil != err {
		log.Fatal("Marshal booklist data error: \n",err)
		return
	}

	base64Data := base64.StdEncoding.EncodeToString(data)

	c.JSON(http.StatusOK, gin.H{
		"msg_id":msgId,
		"pb_data": base64Data,
	})
}

func ChapterList(c *gin.Context) {
	msgId,err := strconv.Atoi(c.Query("msgId"))

	if nil != err {
		log.Fatal(err.Error())
		return
	}

	bookId,err := strconv.Atoi(c.Query("id"))

	if nil != err {
		log.Fatal(err.Error())
		return
	}

	book := utils.Instance().GetBook(bookId)

	log.Print("ChapterListCount is ",len(book.GetChapterList().Chapters),"\n")

	data, err := proto.Marshal(book.GetChapterList())

	if nil != err {
		log.Fatal("Marshal booklist data error: \n",err)
		return
	}

	base64Data := base64.StdEncoding.EncodeToString(data)

	log.Print("data :",base64Data)

	c.JSON(http.StatusOK, gin.H{
		"msg_id":msgId,
		"pb_data": base64Data,
	})
}

func PageList(c *gin.Context) {
	msgId,err := strconv.Atoi(c.Query("msgId"))

	if nil != err {
		log.Fatal(err.Error())
		return
	}

	bookId,err := strconv.Atoi(c.Query("book_id"))

	if nil != err {
		log.Fatal(err.Error())
		return
	}

	chapterId,err := strconv.Atoi(c.Query("chapter_id"))

	if nil != err {
		log.Fatal(err.Error())
		return
	}

	book := utils.Instance().GetBook(bookId)
	chapter := book.GetChapter(chapterId)

	data, err := proto.Marshal(chapter.GetPageList())

	if nil != err {
		log.Fatal("Marshal booklist data error: \n",err)
		return
	}

	base64Data := base64.StdEncoding.EncodeToString(data)

	log.Print("Send Pages List,len is :",len(base64Data))

	c.JSON(http.StatusOK, gin.H{
		"msg_id":msgId,
		"pb_data": base64Data,
	})
}