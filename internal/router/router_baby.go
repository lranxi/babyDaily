package router

import (
	"baby-daily-api/internal/api/v1/babies"
	"baby-daily-api/internal/router/interceptor/auth"
)

// api 路由
func setBabyRouter(s *Server) {

	babyHandler := babies.NewBabyHandler(s.Logger, s.Service.Baby)

	api := s.Engine.Group("/api/v1")
	api.Use(auth.AuthMiddleWare(s.Service.Jwt))
	{
		api.PUT("/baby", babyHandler.Create)
		api.GET("/babies", babyHandler.List)
		api.POST("/baby", babyHandler.Update)
		api.DELETE("/baby/:id", babyHandler.Delete)
	}

}
