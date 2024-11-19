package handler

import (
	"net/http"

	"backend/src/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MicropostHandler struct {
	db *gorm.DB
}

func NewMicropostHandler(db *gorm.DB) *MicropostHandler {
	return &MicropostHandler{db: db}
}

// CreateMicropost godoc
// @Summary      Create new micropost
// @Description  creates a new micropost
// @Tags         microposts
// @Accept       json
// @Produce      json
// @Param        micropost  body      model.Micropost  true  "Micropost Info"
// @Success      201        {object}  model.Micropost
// @Router       /microposts [post]
func (h *MicropostHandler) Create(c *gin.Context) {
	var micropost model.Micropost
	if err := c.ShouldBindJSON(&micropost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.Create(&micropost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, micropost)
}

// FindAll godoc
// @Summary      List microposts
// @Description  get all microposts
// @Tags         microposts
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.Micropost
// @Router       /microposts [get]
func (h *MicropostHandler) FindAll(c *gin.Context) {
	var microposts []model.Micropost
	if err := h.db.Find(&microposts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, microposts)
}
