package users

import "github.com/gin-gonic/gin"

// RegisterRoutes mendaftarkan semua endpoint user.
func RegisterRoutes(r *gin.Engine, h *Handler) {
	g := r.Group("/users")
	{
		g.GET("", h.GetUsers)
		g.GET(":id", h.GetUserByID)
		g.POST("", h.CreateUser)
		g.PUT(":id", h.UpdateUser)
		g.DELETE(":id", h.DeleteUser)
	}
}
