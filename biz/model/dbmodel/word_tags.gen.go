// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dbmodel

const TableNameWordTag = "word_tags"

// WordTag mapped from table <word_tags>
type WordTag struct {
	WordID int64 `gorm:"column:word_id;not null" json:"word_id"`
	TagID  int64 `gorm:"column:tag_id;not null" json:"tag_id"`
}

// TableName WordTag's table name
func (*WordTag) TableName() string {
	return TableNameWordTag
}
