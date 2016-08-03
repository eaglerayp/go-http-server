package main

import (
	// "fmt"
	"github.com/emicklei/go-restful"
	"net/http"
)

func main() {
	// setup service
	droiID := "{Id}"
	ws := new(restful.WebService)
	ws.Consumes("*/*")
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{version}/apps").To(testJSON))
	ws.Route(ws.POST("/{version}/apps").To(testJSON))
	ws.Route(ws.DELETE("/{version}/apps/{AppId}").To(testJSON))
	ws.Route(ws.DELETE("/{version}/apps/clean/").To(testJSON))
	ws.Route(ws.DELETE("/{version}/apps/{AppId}/recover").To(testJSON))
	ws.Route(ws.POST("/{version}/apps/clean/").To(testJSON))
	ws.Route(ws.GET("/{version}/apps/{AppId}/reset").To(testJSON))
	ws.Route(ws.POST("/{version}/apps/production/").To(testJSON))
	ws.Route(ws.PUT("/{version}/apps/{AppId}/key").To(testJSON))
	ws.Route(ws.GET("/{version}/schemas/").To(testJSON))
	ws.Route(ws.GET("/{version}/schemas/{collection}").To(testJSON))
	ws.Route(ws.POST("/{version}/schemas/{collection}").To(testJSON))
	ws.Route(ws.DELETE("/{version}/schemas/{collection}").To(testJSON))
	ws.Route(ws.PUT("/{version}/schemas/{collection}").To(testJSON))
	ws.Route(ws.POST("/{version}/batch").To(testJSON))
	ws.Route(ws.POST("/{version}/jobs").To(testJSON))
	ws.Route(ws.GET("/{version}/jobs").To(testJSON))
	ws.Route(ws.GET("/{version}/jobs/{AppId}").To(testJSON))
	ws.Route(ws.PUT("/{version}/jobs/" + droiID).To(testJSON))
	ws.Route(ws.DELETE("/{version}/jobs/" + droiID).To(testJSON))
	ws.Route(ws.PATCH("/{version}/jobs/" + droiID).To(testJSON))
	ws.Route(ws.PATCH("/{version}/jobs/").To(testJSON))
	ws.Route(ws.POST("/{version}/classes/{collection}").To(testJSON))
	ws.Route(ws.GET("/{version}/classes/{collection}/" + droiID).To(testJSON))
	ws.Route(ws.GET("/{version}/classes/{collection}").To(testJSON))
	ws.Route(ws.PUT("/{version}/classes/{collection}/" + droiID).To(testJSON))
	ws.Route(ws.PATCH("/{version}/classes/{collection}/" + droiID).To(testJSON))
	ws.Route(ws.DELETE("/{version}/classes/{collection}/" + droiID).To(testJSON))
	ws.Route(ws.PATCH("/{version}/classes/{collection}").To(testJSON))
	ws.Route(ws.DELETE("/{version}/classes/{collection}").To(testJSON))
	ws.Route(ws.GET("/{version}/passwords/{user}").To(testJSON))
	ws.Route(ws.PUT("/{version}/passwords/{user}").To(testJSON))
	ws.Route(ws.DELETE("/{version}/passwords/{user}").To(testJSON))
	ws.Route(ws.GET("/{version}/groups/{group}").To(testJSON))
	ws.Route(ws.PUT("/{version}/groups/{group}").To(testJSON))
	ws.Route(ws.DELETE("/{version}/groups/{group}").To(testJSON))
	ws.Route(ws.POST("/{version}/archives/classes/{collection}").To(testJSON))
	ws.Route(ws.GET("/{version}/archives/classes/{collection}").To(testJSON))
	ws.Route(ws.DELETE("/{version}/archives/classes/{collection}/").To(testJSON))
	ws.Route(ws.POST("/{version}/echo/{collection}").To(testJSON))
	ws.Route(ws.GET("/{version}/echo/{collection}").To(testJSON))
	ws.Route(ws.GET("/{version}/echo/{collection}/" + droiID).To(testJSON))
	ws.Route(ws.PUT("/{version}/echo/{collection}/" + droiID).To(testJSON))
	ws.Route(ws.DELETE("/{version}/echo/{collection}/" + droiID).To(testJSON))
	ws.Route(ws.GET("/{version}/ping").To(testJSON))
	ws.Route(ws.GET("/{version}/checkmongos").To(testJSON))
	ws.Route(ws.GET("/{version}/version").To(testJSON))
	ws.Route(ws.GET("/{version}/config").To(testJSON))
	ws.Route(ws.GET("/{version}/clean").To(testJSON))
	ws.Route(ws.GET("/{version}/init").To(testJSON))
	ws.Route(ws.GET("/{version}/cache").To(testJSON))
	ws.Route(ws.GET("/{version}/goinfo").To(testJSON))
	ws.Route(ws.GET("/{version}/kill").To(testJSON))
	ws.Route(ws.GET("/{version}/pfinsert").To(testJSON))
	ws.Route(ws.GET("/{version}/pfread").To(testJSON))
	ws.Route(ws.GET("/{version}/checkindex").To(testJSON))
	restful.Add(ws)

	// setup request + writer

	// run
	// restful.DefaultContainer.ServeHTTP(httpWriter, httpRequest)
	http.ListenAndServe(":8888", nil)
}

func testJSON(rq *restful.Request, rp *restful.Response) {
	json := map[string]int{"Code": 0}
	rp.WriteHeaderAndEntity(http.StatusOK, json)
}
