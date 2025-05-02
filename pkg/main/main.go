package main

import (
	// _ "dechdev/api"
	"dechdev/pkg/config"
	"dechdev/pkg/handler"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func main() {
	exec()
}

func registerRouter(r *gin.RouterGroup) {
	r.GET("/ping", handler.Ping)
	r.GET("/info", handler.Info)
	r.GET("/test", handler.Test)

	r.GET("/data", handler.ReadData)
	r.POST("/data", handler.WriteData)

	r.POST("/buy-product", handler.BuyProduct)
}

func exec() {
	fmt.Println("------------------------------------------------")
	log.Println("--- Start API Server ---")
	fmt.Println("------------------------------------------------")

	app = gin.Default()

	// Handling routing errors
	app.NoRoute(func(c *gin.Context) {
		sb := &strings.Builder{}
		sb.WriteString("routing err: no route, try this:\n")
		for _, v := range app.Routes() {
			sb.WriteString(fmt.Sprintf("%s %s\n", v.Method, v.Path))
		}
		c.String(http.StatusBadRequest, sb.String())
	})

	r := app.Group("/api")

	registerRouter(r)

	srv := &http.Server{
		Addr:         ":" + "4000",
		Handler:      app,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	if config.Mode == "dev" {
		srv.ListenAndServe()
	}

}
