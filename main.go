package main

import (
	"github.com/gin-gonic/gin"
	"lmsmodule/handlers"
)

func main() {
	r := gin.Default()

	r.POST("/login", handlers.LoginHandler)
	r.GET("/courses", handlers.GetCourses)
	r.GET("/courses/:id", handlers.GetCourseByID)
	r.GET("/progress/:user_id", handlers.GetUserProgress)
	r.POST("/progress/:user_id/:assignment_id/complete", handlers.CompleteAssignment)

	r.Run(":8080") // Запуск сервиса на порту 8080
}
