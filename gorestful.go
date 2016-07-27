package main

import (
	// "fmt"
	"github.com/emicklei/go-restful"
	"net/http"
)

var (
	Result string
)

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

func main() {
	// setup service
	ws := new(restful.WebService)
	ws.Consumes("*/*")
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(testGet))
	ws.Route(ws.POST("/").To(testPost))
	restful.Add(ws)

	// setup request + writer

	// run
	// restful.DefaultContainer.ServeHTTP(httpWriter, httpRequest)
	http.ListenAndServe(":8888", nil)
}

func testGet(rq *restful.Request, rp *restful.Response) {
	// fmt.Println("get here")
	rp.WriteHeaderAndEntity(http.StatusOK, responseData)
}

func testPost(rq *restful.Request, rp *restful.Response) {
	// fmt.Println("post here")
	rp.WriteHeaderAndEntity(http.StatusOK, "")
}
