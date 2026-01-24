package models

import (
	"time"

	"github.com/vlahanam/company-management/common"
)

type SQLModel struct {
	ID        uint64       `json:"-" gorm:"column:id"`
	FakeId    *common.UID `json:"id" gorm:"-"`
	CreatedAt *time.Time  `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time  `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func NewSQLModel() SQLModel {
	now := time.Now().UTC()

	return SQLModel{
		ID:        0,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}

func (sqlModel *SQLModel) Mask(objectId int64) {
	uid := common.NewUID(uint32(sqlModel.ID), objectId, 1)
	sqlModel.FakeId = &uid
}
