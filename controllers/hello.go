package controllers

import (
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

func GetHello(c *gin.Context) {
	v, _ := time.Now().UTC().MarshalText()
	c.String(http.StatusOK, "Hello world: "+string(v))
}
