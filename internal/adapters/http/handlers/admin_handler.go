package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/input/services"
)

type AdminHandler struct {
	adminService services.AdminService
}

func NewAdminHandler(adminService services.AdminService) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
	}
}

// GetActiveOrders godoc
// @Summary      Get active orders
// @Description  Retrieve all active orders
// @Tags         admin
// @Accept       json
// @Produce      json
// @Success      200  {array}  entities.Order
// @Router       /admin/orders/active [get]
func (h *AdminHandler) GetActiveOrders(c *gin.Context) {
	orders, err := h.adminService.GetActiveOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve active orders",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, orders)
}
