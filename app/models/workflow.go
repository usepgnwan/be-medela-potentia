package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Workflow struct {
	ID           string          `json:"id" gorm:"type:uuid;primaryKey" swaggerignore:"true"`
	Name         string          `json:"name" gorm:"size:255" validate:"required"`
	UserID       string          `json:"user_id" gorm:"size:50;not null" swaggerignore:"true"`
	RequestBy    *User           `json:"request_by" gorm:"foreignKey:UserID;references:ID" swaggerignore:"true"`
	WorkflowStep *[]WorkflowStep `json:"workflow_step" gorm:"foreignKey:WorkflowId" swaggerignore:"true"`
	CreatedAt    time.Time       `json:"created_at" swaggerignore:"true"`
	UpdatedAt    time.Time       `json:"updated_at" swaggerignore:"true"`
	DeletedAt    gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty" swaggerignore:"true"`
}

func (Workflow) TableName() string {
	return "tbl_workflow"
}

func (w *Workflow) BeforeCreate(tx *gorm.DB) (err error) {
	if w.ID == "" {
		w.ID = uuid.NewString()
	}
	return
}

// type FormWorkflow struct {
// 	Name   string `json:"name" gorm:"size:255" validate:"required"`
// 	UserID string `json:"user_id" swaggerignore:"true"`
// }
