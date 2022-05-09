package feedbacks

import (
	"baby-daily-api/internal/model"
	"baby-daily-api/internal/repository/database/feedbacks"
)

type feedbackService struct {
	repository feedbacks.FeedbackRepository
}

type FeedbackService interface {
	Create(feedback *model.Feedback) (*model.Feedback, error)
}

func NewFeedbackService(rep feedbacks.FeedbackRepository) FeedbackService {
	return &feedbackService{repository: rep}
}

func (f *feedbackService) Create(feedback *model.Feedback) (*model.Feedback, error) {
	return f.repository.Insert(feedback)
}
