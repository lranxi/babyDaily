package router

import (
	"baby-daily-api/internal/api/v1/record"
	"baby-daily-api/internal/router/interceptor/auth"
)

func setRecordRouter(s *Server) {
	recordHandler := record.NewRecordHandler(s.Logger, s.Service.Record)

	api := s.Engine.Group("/api/v1")
	api.Use(auth.AuthMiddleWare(s.Service.Jwt))
	{
		api.GET("/record/:babyId/:day", recordHandler.List)
		api.PUT("/record", recordHandler.Create)
		api.POST("/record", recordHandler.Update)
		api.DELETE("/record/:id", recordHandler.Delete)
		api.GET("/record/detail/:babyId/:type/:day", recordHandler.Details)
	}
}
