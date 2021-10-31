package handler

import "github.com/gin-gonic/gin"

type Handler struct {
	
}

// Инициализация маршрутов сервера
func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Заводим контроллер /auth
	auth := router.Group("/auth")
	{
		// Задаем методы контроллера /{relativePath}/{processName}
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-ip", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllItems)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_i", h.deleteItem)
			}
		}
	}

	return router
}