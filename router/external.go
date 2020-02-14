package router

import (
	"comic_server/handler"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"time"
)

func Init() {
	router := gin.New()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, X-Token",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	getTitleList := router.Group("/")
	{
		getTitleList.GET("", handler.Index)
	}

	ginpprof.Wrap(router)

	err := router.Run(":8081")

	if nil != err {
		panic(err)
	}
}