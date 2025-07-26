package routes

import (
	"mini-library-backend/controllers"
	"mini-library-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Auth
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	// Book
	book := router.Group("/books")
	{
		book.GET("", controllers.GetBooks)
		book.GET("/:id", controllers.GetBookDetail)

		book.Use(middlewares.AuthMiddleware("admin"))
		book.POST("", controllers.CreateBook)
		book.PUT("/:id", controllers.UpdateBook)
		book.DELETE("/:id", controllers.DeleteBook)
	}

	// Borrow
	borrow := router.Group("/borrow", middlewares.AuthMiddleware("user", "admin"))
	{
		borrow.POST("/:id", controllers.BorrowBook)
		borrow.POST("/return/:id", controllers.ReturnBook)
		borrow.GET("/history", controllers.GetUserHistory)
	}

	admin := router.Group("/admin", middlewares.AuthMiddleware("admin"))
	{
		admin.GET("/history", controllers.GetAllHistory)
	}
}
