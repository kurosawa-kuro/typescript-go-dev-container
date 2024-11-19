package handler

import (
	"github.com/gin-gonic/gin"
)

// PingHandler godoc
// @Summary      Ping test endpoint
// @Description  returns pong message
// @Tags         ping
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
