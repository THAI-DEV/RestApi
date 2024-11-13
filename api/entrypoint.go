package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"dechdev/pkg/config"
	"dechdev/pkg/handler"
)

var (
	app *gin.Engine
)

func registerRouter(r *gin.RouterGroup) {
	r.GET("/ping", handler.Ping)
	r.GET("/info", handler.Info)
	r.GET("/test", handler.Test)
}

func init() {
	fmt.Println("------------------------------------------------")
	log.Println("--- Start API Server ---")
	fmt.Println("------------------------------------------------")

	app = gin.New()

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

	if config.Mode == "dev" {
		app.Run(":4000")
	}
}

// entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
