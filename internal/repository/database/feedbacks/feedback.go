package feedbacks

import (
	"baby-daily-api/internal/model"
	"baby-daily-api/internal/repository/database"
	"gorm.io/gorm"
)

type feedbackRepository struct {
	db *gorm.DB
}

type FeedbackRepository interface {
	Migrate() error
	// Insert 新增反馈
	Insert(feedback *model.Feedback) (*model.Feedback, error)
}

func NewFeedbackRepository() FeedbackRepository {
	return &feedbackRepository{db: database.Get()}
}

func (f *feedbackRepository) Insert(feedback *model.Feedback) (*model.Feedback, error) {
	return feedback, f.db.Create(feedback).Error
}

func (r *feedbackRepository) Migrate() error {
	return r.db.AutoMigrate(&model.Feedback{})
}
