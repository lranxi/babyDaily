package babies

import (
	"baby-daily-api/configs"
	"baby-daily-api/internal/model"
	"baby-daily-api/internal/response"
	"baby-daily-api/internal/service/babies"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type babyHandler struct {
	babyService babies.BabyService
	logger      *zap.Logger
}

type BabyHandler interface {
	// Create 创建宝宝信息
	Create(c *gin.Context)
	// Update 更新宝宝信息
	Update(c *gin.Context)
	// Delete 删除宝宝信息
	Delete(c *gin.Context)
	// List 宝贝列表
	List(c *gin.Context)
}

func NewBabyHandler(log *zap.Logger, babyService babies.BabyService) BabyHandler {
	return &babyHandler{
		logger:      log,
		babyService: babyService,
	}
}

// Create 创建baby信息
func (b *babyHandler) Create(c *gin.Context) {
	createBaby := new(model.CreateBaby)
	if err := c.BindJSON(createBaby); err != nil {
		response.Fail(c, http.StatusBadRequest, "缺少必要参数")
		return
	}

	val, _ := c.Get(configs.UserContextKey)
	user, ok := val.(*model.User)
	if !ok {
		response.Fail(c, http.StatusBadRequest, "登录过期，请重新登录")
		return
	}
	baby := createBaby.GetBaby()
	baby.UserId = user.ID
	baby, err := b.babyService.Create(baby)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "创建用户失败")
		return
	}
	response.Success(c, baby)
}

func (b *babyHandler) Update(c *gin.Context) {
	baby := new(model.Baby)
	if err := c.BindJSON(baby); err != nil {
		response.Fail(c, http.StatusBadRequest, "缺少必要参数")
		return
	}
	baby, err := b.babyService.Update(baby)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "更新数据失败")
		return
	}
	response.Success(c, baby)
}

func (b *babyHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		response.Fail(c, http.StatusBadRequest, "缺少必要参数")
		return
	}
	val, _ := c.Get(configs.UserContextKey)
	user, ok := val.(*model.User)
	if !ok {
		response.Fail(c, http.StatusBadRequest, "登录过期,请重新登录")
		return
	}
	id, _ := strconv.Atoi(idStr)
	baby := &model.Baby{ID: id, UserId: user.ID}
	err := b.babyService.Delete(baby)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "删除失败")
		return
	}
	response.Success(c, "")
}

func (b *babyHandler) List(c *gin.Context) {
	val, _ := c.Get(configs.UserContextKey)
	user, ok := val.(*model.User)
	if !ok {
		response.Fail(c, http.StatusBadRequest, "登录过期,请重新登录")
		return
	}
	babies, err := b.babyService.List(user.ID)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "获取数据失败")
		return
	}
	response.Success(c, babies)
}
