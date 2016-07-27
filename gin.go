package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	r.GET("/", testGet)
	r.POST("/", testPost)
	r.Run(":8888")
}

func testGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"Code": 0,
	})
}
func testPost(c *gin.Context) {
	c.String(200, "")
}
