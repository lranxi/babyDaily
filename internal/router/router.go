package router

import (
	"baby-daily-api/configs"
	cache "baby-daily-api/internal/repository/cache"
	"baby-daily-api/internal/repository/database"
	babiesRepository "baby-daily-api/internal/repository/database/babies"
	feedbacksRepository "baby-daily-api/internal/repository/database/feedbacks"
	recordsRepository "baby-daily-api/internal/repository/database/records"
	usersRepository "baby-daily-api/internal/repository/database/users"
	"baby-daily-api/internal/router/interceptor/cors"
	"baby-daily-api/internal/router/interceptor/logger"
	"baby-daily-api/internal/service/babies"
	"baby-daily-api/internal/service/feedbacks"
	"baby-daily-api/internal/service/jwt"
	"baby-daily-api/internal/service/miniapp"
	"baby-daily-api/internal/service/records"
	"baby-daily-api/internal/service/users"
	"errors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Server struct {
	Engine     *gin.Engine
	Logger     *zap.Logger
	Cache      cache.Cache
	DB         *gorm.DB
	Service    *Service
	Repository *Repository
}

// Service service层实例
type Service struct {
	User     users.UserService
	Baby     babies.BabyService
	Record   records.RecordService
	Feedback feedbacks.FeedbackService
	Jwt      *jwt.JwtService
	MiniApp  miniapp.MiniAppService
}

// Repository 持久战实例
type Repository struct {
	User     usersRepository.UserRepository
	Baby     babiesRepository.BabyRepository
	Record   recordsRepository.RecordRepository
	Feedback feedbacksRepository.FeedbackRepository
}

// New 创建server
func New(accessLogger *zap.Logger) (*Server, error) {
	cfg := configs.Get().Server
	// 设置启动模式
	gin.SetMode(cfg.Env)
	engine := gin.New()
	engine.Use(
		logger.LogMiddleware(accessLogger),
		cors.CORSMiddleware(),
		gin.Recovery(),
	)

	// 创建redis
	redis, err := cache.NewCache()
	if err != nil {
		return nil, err
	}

	db := database.Get()
	if db == nil {
		return nil, errors.New("failed to connect database")
	}

	server := &Server{
		Logger: accessLogger,
		Cache:  redis,
		Engine: engine,
		DB:     db,
	}

	server.newRepository()
	server.newService()

	// debug环境开启pprof
	if configs.Get().Server.Env == "debug" {
		pprof.Register(server.Engine)
	}

	// 用户路由
	setUsersRouter(server)

	setRecordRouter(server)
	setBabyRouter(server)
	setFeedbackRouter(server)

	return server, nil

}

// 创建持久层实例
func (s *Server) newRepository() {
	user := usersRepository.New()
	record := recordsRepository.NewRecordRepository()
	baby := babiesRepository.NewBabyRepository()
	feedback := feedbacksRepository.NewFeedbackRepository()

	s.Repository = &Repository{
		User:     user,
		Baby:     baby,
		Record:   record,
		Feedback: feedback,
	}
}

// 创建业务层实例
func (s *Server) newService() {
	user := users.NewUserService(s.Repository.User, s.Cache)
	record := records.NewRecordService(s.Repository.Record)
	baby := babies.NewBabyService(s.Repository.Baby)
	feedback := feedbacks.NewFeedbackService(s.Repository.Feedback)
	jwt := jwt.NewJwtService()
	miniApp := miniapp.NewMiniAppService()

	s.Service = &Service{
		User:     user,
		Baby:     baby,
		Record:   record,
		Feedback: feedback,
		Jwt:      jwt,
		MiniApp:  miniApp,
	}

}
