package mysql_model

// Comments 对应数据库表 comments
type Comments struct {
	Id        int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Content   string `gorm:"column:content;not null;default:''" json:"content"`
	UserId    int64  `gorm:"column:user_id;not null;default:0" json:"user_id"`
	PostId    int64  `gorm:"column:post_id;not null;default:0" json:"post_id"`
	CreatedAt int64  `gorm:"column:created_at;not null;default:0" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;not null;default:0" json:"updated_at"`
	DeletedAt int64  `gorm:"column:deleted_at;not null;default:0" json:"deleted_at"`
}

func (Comments) TableName() string {
	return "comments"
}
