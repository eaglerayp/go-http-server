package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	r.GET("/:version/apps", testJSON)
	r.POST("/:version/apps", testJSON)
	r.DELETE("/:version/apps/:AppId", testJSON)
	// r.DELETE("/:version/apps/clean/", testJSON)
	r.DELETE("/:version/apps/:AppId/recover", testJSON)
	r.POST("/:version/apps/clean/", testJSON)
	r.GET("/:version/apps/:AppId/reset", testJSON)
	r.POST("/:version/apps/production/", testJSON)
	r.PUT("/:version/apps/:AppId/key", testJSON)

	r.GET("/:version/schemas/", testJSON)
	r.GET("/:version/schemas/:collection", testJSON)
	r.POST("/:version/schemas/:collection", testJSON)
	r.DELETE("/:version/schemas/:collection", testJSON)
	r.PUT("/:version/schemas/:collection", testJSON)

	r.POST("/:version/batch", testJSON)

	r.POST("/:version/jobs", testJSON)
	r.GET("/:version/jobs", testJSON)
	r.GET("/:version/jobs/:AppId", testJSON)
	r.PUT("/:version/jobs/:id", testJSON)
	r.DELETE("/:version/jobs/:id", testJSON)
	r.PATCH("/:version/jobs/:id", testJSON)
	r.PATCH("/:version/jobs/", testJSON)

	r.POST("/:version/classes/:collection", testJSON)
	r.GET("/:version/classes/:collection/:id", testJSON)
	r.GET("/:version/classes/:collection", testJSON)
	r.PUT("/:version/classes/:collection/:id", testJSON)
	r.PATCH("/:version/classes/:collection/:id", testJSON)
	r.DELETE("/:version/classes/:collection/:id", testJSON)

	r.PATCH("/:version/classes/:collection", testJSON)
	r.DELETE("/:version/classes/:collection", testJSON)

	r.GET("/:version/passwords/:user", testJSON)
	r.PUT("/:version/passwords/:user", testJSON)
	r.DELETE("/:version/passwords/:user", testJSON)

	r.GET("/:version/groups/:group", testJSON)
	r.PUT("/:version/groups/:group", testJSON)
	r.DELETE("/:version/groups/:group", testJSON)

	r.POST("/:version/archives/classes/:collection", testJSON)
	r.GET("/:version/archives/classes/:collection", testJSON)
	r.DELETE("/:version/archives/classes/:collection/", testJSON)

	// echo
	r.POST("/:version/echo/:collection", testJSON)
	r.GET("/:version/echo/:collection", testJSON)
	r.GET("/:version/echo/:collection/:id", testJSON)
	r.PUT("/:version/echo/:collection/:id", testJSON)
	r.DELETE("/:version/echo/:collection/:id", testJSON)
	r.GET("/:version/ping", testJSON)
	r.GET("/:version/checkmongos", testJSON)
	r.GET("/:version/version", testJSON)
	r.GET("/:version/config", testJSON)
	r.GET("/:version/clean", testJSON)
	r.GET("/:version/init", testJSON)
	r.GET("/:version/cache", testJSON)
	// r.GET("/stats", testJSON)
	r.GET("/:version/goinfo", testJSON)
	r.GET("/:version/kill", testJSON)
	r.GET("/:version/pfinsert", testJSON)
	r.GET("/:version/pfread", testJSON)
	// r.GET("/pwmigration", testJSON)
	r.GET("/:version/checkindex", testJSON)
	r.Run(":8888")
}

func testJSON(c *gin.Context) {
	c.JSON(200, gin.H{
		"Code": 0,
	})
}
