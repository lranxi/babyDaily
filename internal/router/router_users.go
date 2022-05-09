package router

import (
	"baby-daily-api/internal/api/v1/users"
)

// api 路由
func setUsersRouter(s *Server) {

	userHandler := users.New(s.Logger, s.Service.User, s.Service.MiniApp, s.Service.Baby, s.Service.Jwt)

	api := s.Engine.Group("/api/v1")
	{
		api.POST("/login", userHandler.Login)
	}

}
