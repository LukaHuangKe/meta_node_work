package mysql_model

// Posts 对应数据库表 posts
type Posts struct {
	Id        int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Title     string `gorm:"column:title;not null;default:''" json:"title"`
	Content   string `gorm:"column:content;not null;default:''" json:"content"`
	UserId    int64  `gorm:"column:user_id;not null;default:0" json:"user_id"`
	CreatedAt int64  `gorm:"column:created_at;not null;default:0" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;not null;default:0" json:"updated_at"`
	DeletedAt int64  `gorm:"column:deleted_at;not null;default:0" json:"deleted_at"`
}

func (Posts) TableName() string {
	return "posts"
}
