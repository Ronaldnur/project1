package handler

import (
	"Project1/docs"
	"Project1/infra/database"
	"Project1/repository/todo_repository/todo_pg"
	"Project1/service"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	// swagger embed files
)

func StartApp() {

	database.InitiazeDatabase()

	db := database.GetDatabaseInstance()

	TodoRepo := todo_pg.NewTodoPG(db)

	TodoService := service.NewTodoService(TodoRepo)

	TodoHandler := NewTodoHandler(TodoService)

	route := gin.Default()

	docs.SwaggerInfo.Title = "Project 1"
	docs.SwaggerInfo.Description = "Ini adalah Project ke 1 dari kelas Kampus Merdeka"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	route.POST("/todos", TodoHandler.CreateData)

	route.GET("/todos", TodoHandler.ReadData)

	route.GET("/todos/:todoId", TodoHandler.ReadDataById)

	route.PUT("/todos/:todoId", TodoHandler.UpdateDataById)

	route.DELETE("/todos/:todoId", TodoHandler.DeleteTodoData)

	route.Run(":8080")
}
