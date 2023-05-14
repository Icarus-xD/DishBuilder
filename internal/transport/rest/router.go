package rest

import (
	"github.com/gin-gonic/gin"
)

type Dish interface {
	GetByIngredients(ctx *gin.Context)
}

type Handler struct {
	dishService Dish
}

func NewHandler(d Dish) *Handler {
	return &Handler{
		dishService: d,
	}
}

func (h *Handler) InitRouter(r *gin.Engine) {
	routes := r.Group("dish")
	{
		routes.GET("/", h.dishService.GetByIngredients)
	}
}