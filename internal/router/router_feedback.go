package router

import (
	"baby-daily-api/internal/api/v1/feedback"
	"baby-daily-api/internal/router/interceptor/auth"
)

func setFeedbackRouter(s *Server) {
	feedbackHandler := feedback.NewFeedbackHandler(s.Logger, s.Service.Feedback)

	api := s.Engine.Group("/api/v1")
	api.Use(auth.AuthMiddleWare(s.Service.Jwt))
	{
		api.POST("/feedback", feedbackHandler.Create)
	}
}
