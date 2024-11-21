package model

type AdminUser struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;comment:主键ID" json:"id"`
	Username string `gorm:"unique;comment:用户名" json:"username"`
	Password string `gorm:"comment:密码" json:"password"`
}
