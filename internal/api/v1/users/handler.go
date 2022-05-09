package users

import (
	"baby-daily-api/configs"
	"baby-daily-api/internal/model"
	"baby-daily-api/internal/response"
	"baby-daily-api/internal/service/babies"
	"baby-daily-api/internal/service/jwt"
	"baby-daily-api/internal/service/miniapp"
	"baby-daily-api/internal/service/users"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type userHandler struct {
	logger         *zap.Logger
	userService    users.UserService
	jwtService     *jwt.JwtService
	miniAppService miniapp.MiniAppService
	babyService    babies.BabyService
}

type UserHandler interface {
	// Login 用户登录注册接口
	Login(c *gin.Context)
}

func New(log *zap.Logger, user users.UserService, miniApp miniapp.MiniAppService, baby babies.BabyService, jwt *jwt.JwtService) UserHandler {
	return &userHandler{
		logger:         log,
		userService:    user,
		jwtService:     jwt,
		miniAppService: miniApp,
		babyService:    baby,
	}
}

// Login 用户登录
func (u *userHandler) Login(c *gin.Context) {
	createUser := new(model.CreatedUser)
	if err := c.BindJSON(createUser); err != nil {
		u.logger.Error("用户登录缺少参数", zap.Error(err))
		response.Fail(c, http.StatusBadRequest, "缺少参数")
		return
	}
	// 获取openId
	session, err := u.miniAppService.Code2Session(createUser.Code)
	if err != nil {
		u.logger.Error("获取用户OpenId失败", zap.Error(err))
		response.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	// 检查用户是否存在，如果存在则生成token，不存在就新增用户
	user, err := u.userService.GetUserByOpenId(session.OpenId)
	if err == nil {
		// 用户存在，生成
		token, _ := u.jwtService.CreateToken(user)
		authUser := &model.AuthUser{
			ID:       user.ID,
			Nickname: user.Nickname,
			OpenId:   user.OpenId,
			Avatar:   user.Avatar,
			Token:    token,
			Expires:  configs.TokenExpires,
		}

		// 更新用户信息
		user.Avatar = createUser.Avatar
		user.Nickname = createUser.Nickname
		user.OpenId = session.OpenId
		u.userService.Update(session.OpenId, user)

		// 新建一条默认的baby数据

		response.Success(c, authUser)
		return
	}
	// 用户不存在
	user = &model.User{
		Nickname:  createUser.Nickname,
		OpenId:    session.OpenId,
		UnionId:   session.UnionId,
		Avatar:    createUser.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// 创建用户
	user, err = u.userService.Create(user)
	u.babyService.GenerateDefault(user.ID)
	if err != nil {
		u.logger.Error("创建用户信息失败", zap.Error(err))
		response.Fail(c, http.StatusInternalServerError, "用户创建失败")
		return
	}
	token, _ := u.jwtService.CreateToken(user)
	authUser := &model.AuthUser{
		ID:       user.ID,
		Nickname: user.Nickname,
		OpenId:   user.OpenId,
		Avatar:   user.Avatar,
		Token:    token,
		Expires:  configs.TokenExpires,
	}
	response.Success(c, authUser)
}
