// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dbmodel

import (
	"time"
)

const TableNameTag = "tags"

// Tag mapped from table <tags>
type Tag struct {
	ID         int64     `gorm:"column:id;primaryKey;comment:ID" json:"id"` // ID
	TagTypeID  int64     `gorm:"column:tag_type_id;not null" json:"tag_type_id"`
	Value      string    `gorm:"column:value;not null;comment:标签名称" json:"value"` // 标签名称
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	TagType    TagType   `gorm:"foreignKey:tag_type_id" json:"tag_type"`
}

// TableName Tag's table name
func (*Tag) TableName() string {
	return TableNameTag
}
