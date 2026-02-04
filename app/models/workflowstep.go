package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkflowStep struct {
	ID         string         `json:"id" gorm:"type:uuid;primaryKey" swaggerignore:"true"`
	Level      int            `json:"level" gorm:"not null" validate:"required"`
	WorkflowId string         `gorm:"size:255" json:"workflow_id" validate:"required"`
	Workflow   *Workflow      `json:"workflow,omitempty" gorm:"foreignKey:WorkflowId" swaggerignore:"true"`
	RoleId     uint           `gorm:"size:255" json:"role_id"`
	Actor      *UserRole      `json:"actor,omitempty" gorm:"foreignKey:RoleId" swaggerignore:"true"`
	MinAmount  int64          `json:"min_amount" gorm:"not null" validate:"required"`
	CreatedAt  time.Time      `json:"created_at" swaggerignore:"true"`
	UpdatedAt  time.Time      `json:"updated_at" swaggerignore:"true"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" swaggerignore:"true"`
}

func (WorkflowStep) TableName() string {
	return "tbl_workflow_step"
}

func (w *WorkflowStep) BeforeCreate(tx *gorm.DB) (err error) {
	if w.ID == "" {
		w.ID = uuid.NewString()
	}
	return
}
