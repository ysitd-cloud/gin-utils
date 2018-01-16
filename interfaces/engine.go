package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Engine interface {
	gin.IRouter
	http.Handler
	Run(addr ...string) error
	HandleContext(c *gin.Context)
}
