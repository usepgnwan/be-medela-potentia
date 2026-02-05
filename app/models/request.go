package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Request struct {
	ID          string         `json:"id" gorm:"type:uuid;primaryKey" swaggerignore:"true"`
	WorkflowId  string         `gorm:"size:255" json:"workflow_id" validate:"required"`
	Workflow    *Workflow      `json:"workflow" gorm:"foreignKey:WorkflowId" swaggerignore:"true"`
	CurrentStep int            `json:"current_step" validate:"required" swaggerignore:"true"`
	Status      string         `json:"status" gorm:"size:255" validate:"required" swaggerignore:"true"`
	UserID      string         `json:"user_id" gorm:"size:50;not null" swaggerignore:"true"`
	Amount      int64          `json:"amount" gorm:"not null" validate:"required"`
	ApproveBy   *User          `json:"approve_by" gorm:"foreignKey:UserID;references:ID" swaggerignore:"true"`
	CreatedAt   time.Time      `json:"created_at" swaggerignore:"true"`
	UpdatedAt   time.Time      `json:"updated_at" swaggerignore:"true"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" swaggerignore:"true"`
}

func (Request) TableName() string {
	return "tbl_request"
}

func (w *Request) BeforeCreate(tx *gorm.DB) (err error) {
	if w.ID == "" {
		w.ID = uuid.NewString()
	}
	return
}
