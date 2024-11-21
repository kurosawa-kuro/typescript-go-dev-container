package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PingDBHandler struct {
	db *gorm.DB
}

func NewPingDBHandler(db *gorm.DB) *PingDBHandler {
	return &PingDBHandler{db: db}
}

// PingDB godoc
// @Summary      Ping DB
// @Description  ping DB
// @Tags         ping-db
// @Accept       json
// @Produce      json
// @Success      200        {object}  string
// @Router       /ping-db [get]
func (h *PingDBHandler) PingDB(c *gin.Context) {
	sqlDB, err := h.db.DB()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Just ping the existing connection
	err = sqlDB.Ping()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, "DB connected")
}
