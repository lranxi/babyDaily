package record

import (
	"baby-daily-api/configs"
	"baby-daily-api/internal/model"
	"baby-daily-api/internal/response"
	"baby-daily-api/internal/service/records"
	"baby-daily-api/pkg/timeutil"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type recordHandler struct {
	logger        *zap.Logger
	recordService records.RecordService
}

type RecordHandler interface {
	Details(c *gin.Context)
	List(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Create(c *gin.Context)
}

func NewRecordHandler(log *zap.Logger, recordService records.RecordService) RecordHandler {
	return &recordHandler{
		logger:        log,
		recordService: recordService,
	}
}

// Create 创建baby信息
func (r *recordHandler) Create(c *gin.Context) {
	createRecord := new(model.CreateRecord)
	if err := c.BindJSON(createRecord); err != nil {
		response.Fail(c, http.StatusBadRequest, "缺少必要参数")
		return
	}

	val, _ := c.Get(configs.UserContextKey)
	user, ok := val.(*model.User)
	if !ok {
		response.Fail(c, http.StatusUnauthorized, "登录失效，请重新登录")
		return
	}
	record := createRecord.GetRecord()
	record.UserId = user.ID
	record, err := r.recordService.Create(record)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "创建记录失败")
		return
	}
	response.Success(c, record)
}

// Update 更新记录
func (r *recordHandler) Update(c *gin.Context) {
	record := new(model.Record)
	if err := c.BindJSON(record); err != nil {
		response.Fail(c, http.StatusBadRequest, "缺少必要参数")
		return
	}
	val, _ := c.Get(configs.UserContextKey)
	user, ok := val.(*model.User)
	if !ok {
		response.Fail(c, http.StatusUnauthorized, "登录失效，请重新登录")
		return
	}
	record.UserId = user.ID
	record, err := r.recordService.Update(record)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "操作失败")
		return
	}
	response.Success(c, record)
}

// Delete 删除记录
func (r *recordHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		response.Fail(c, http.StatusBadRequest, "缺少必要参数")
		return
	}
	val, _ := c.Get(configs.UserContextKey)
	user, ok := val.(*model.User)
	if !ok {
		response.Fail(c, http.StatusUnauthorized, "登录失效，请重新登录")
		return
	}
	id, _ := strconv.Atoi(idStr)
	record := &model.Record{ID: id, UserId: user.ID}
	err := r.recordService.Delete(record)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "操作失败")
		return
	}
	response.Success(c, "")
}

// List 记录列表
func (r *recordHandler) List(c *gin.Context) {
	val, _ := c.Get(configs.UserContextKey)
	user, ok := val.(*model.User)
	if !ok {
		response.Fail(c, http.StatusUnauthorized, "登录失效，请重新登录")
		return
	}
	babyIdStr := c.Param("babyId")
	if babyIdStr == "" {
		response.Fail(c, http.StatusBadRequest, "缺少必要参数")
		return
	}
	day := c.Param("day")
	if day == "" {
		day = timeutil.CurDay()
	}
	babyId, _ := strconv.Atoi(babyIdStr)
	records, err := r.recordService.List(user.ID, babyId, day)
	if err != nil {
		r.logger.Error("查询记录数据失败", zap.Error(err))
		response.Fail(c, http.StatusInternalServerError, "获取数据失败")
		return
	}
	response.Success(c, records)
}

// Details 记录详情
func (r *recordHandler) Details(c *gin.Context) {
	val, _ := c.Get(configs.UserContextKey)
	user, ok := val.(*model.User)
	if !ok {
		response.Fail(c, http.StatusUnauthorized, "登录失效，请重新登录")
		return
	}
	babyIdStr := c.Param("babyId")
	if babyIdStr == "" {
		response.Fail(c, http.StatusBadRequest, "缺少必要参数")
		return
	}
	babyId, _ := strconv.Atoi(babyIdStr)
	rTypeStr := c.Param("type")
	if rTypeStr == "" {
		response.Fail(c, http.StatusBadRequest, "缺少必要参数")
		return
	}
	rType, _ := strconv.Atoi(rTypeStr)

	day := c.Param("day")
	if day == "" {
		day = timeutil.CurDay()
	}
	details, err := r.recordService.Details(user.ID, babyId, rType, day)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, "获取数据失败")
		return
	}
	response.Success(c, details)
}
