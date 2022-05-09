package feedback

import (
	"baby-daily-api/configs"
	"baby-daily-api/internal/model"
	"baby-daily-api/internal/response"
	"baby-daily-api/internal/service/feedbacks"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"unicode/utf8"
)

type feedbackHandler struct {
	logger          *zap.Logger
	feedbackService feedbacks.FeedbackService
}

type FeedbackHandler interface {
	// Create 提交反馈数据
	Create(c *gin.Context)
}

func NewFeedbackHandler(log *zap.Logger, feedbackService feedbacks.FeedbackService) FeedbackHandler {
	return &feedbackHandler{
		logger:          log,
		feedbackService: feedbackService,
	}
}

func (f *feedbackHandler) Create(c *gin.Context) {
	feedback := &model.Feedback{}
	c.ShouldBindJSON(feedback)

	if utf8.RuneCountInString(feedback.Content) <= 0 || utf8.RuneCountInString(feedback.Content) > 200 {
		response.Fail(c, http.StatusUnauthorized, "反馈信息字数限制为1～200字")
		return
	}

	// 从cookie中读取用户ID
	val, _ := c.Get(configs.UserContextKey)
	user, ok := val.(*model.User)
	if !ok {
		response.Fail(c, http.StatusUnauthorized, "登录失效，请重新登录")
		return
	}
	feedback.UserId = user.ID
	back, err := f.feedbackService.Create(feedback)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, "提交失败")
		return
	}
	response.Success(c, back)
}
