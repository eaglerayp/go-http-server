package main

import "github.com/kataras/iris"

func main() {
	iris.Get("/", testGet)
	iris.Post("/", testPost)
	// iris.Put("/add", testPut)
	// iris.Delete("/remove", testDelete)
	// iris.Head("/testHead", testHead)
	// iris.Patch("/testPatch", testPatch)
	// iris.Options("/testOptions", testOptions)

	iris.Listen(":8888")
}

type Response struct {
	Code int `json:"Code"`
	// Message string      `json:"Message,omitempty"`
	// Result  interface{} `json:"Result,omitempty"`
	// Count   int         `json:"Count,omitempty"`
	// MD5     string      `json:"MD5,omitempty"`
}

var responseData Response

// var responseData = make(map[string]interface{})

func init() {
	responseData = Response{Code: 0}
	// responseData["Code"] = 0
}

func testGet(c *iris.Context) {
	c.JSON(iris.StatusOK, responseData)
}
func testPost(c *iris.Context) {
	c.Write("")
}
