package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func apiGetaway() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("------------------------ start ------------------------")
		for k, v := range c.Request.Header {
			fmt.Println(k, ":", v)
		}
		fmt.Println("------------------------  end  ------------------------")
		c.Next()

		return
	}
}


func main()  {
	r := gin.Default()

	r.Use(apiGetaway())  // global middleware

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "index for golang.")
	})

	_ = r.Run(":6005")

}
