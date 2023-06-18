package entity

// Entities are the results we'd expect from the database.
type User struct {
	UserID       int64  `gorm:"column:id;primary_key"`
	Name         string `gorm:"column:name"`
	Email        string `gorm:"column:email"`
	Password     string `gorm:"column:password"`
	CodeVerified string `gorm:"column:code_verified"`
}
