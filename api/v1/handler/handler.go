package handler

import (
	_ "Server/cmd/docs"
	"github.com/gin-gonic/gin"
)

// GetAdminHandle
// @Summary Get admin
// @name GetAdminHandle
// @Accept json
// @Produce json
// @Description Getting admin
// @Router /api/v1/admin [get]
// @Success 200 string
func GetAdminHandle(c *gin.Context) {
	c.String(200, "GET!!!!@#")
}
