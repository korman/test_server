package main

import (
	"comic_server/router"
	"comic_server/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	println("papapapap")

	gin.SetMode(gin.DebugMode)

	utils.Instance().LoadComicFiles("E:\\hanman_jpg")

	router.Init()
}
