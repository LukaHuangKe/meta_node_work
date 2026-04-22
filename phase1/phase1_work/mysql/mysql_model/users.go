package mysql_model

// Users 对应数据库表 users
type Users struct {
	Id        int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Username  string `gorm:"column:username;not null;default:''" json:"username"`
	Password  string `gorm:"column:password;not null;default:''" json:"password"`
	Email     string `gorm:"column:email;not null;default:''" json:"email"`
	Status    int8   `gorm:"column:status;not null;default:1" json:"status"`
	CreatedAt int64  `gorm:"column:created_at;not null;default:0" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;not null;default:0" json:"updated_at"`
	DeletedAt int64  `gorm:"column:deleted_at;not null;default:0" json:"deleted_at"`
}

func (Users) TableName() string {
	return "users"
}
