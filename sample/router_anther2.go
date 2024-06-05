package app

import "github.com/gin-gonic/gin"

func SetRouter() *gin.Engine {
	r := gin.Default()
	hoge := r.Group("/hoge")
    {
        hoge.GET("/fuga")
        hoge.GET("/piyo/*action")
    }
	return r
}
