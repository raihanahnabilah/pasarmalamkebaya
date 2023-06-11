package entity

// Entities are the results we'd expect from the database.
type User struct {
	UserID       int    `gorm:"column:user_id;primary_key"`
	Name         string `gorm:"column:name"`
	Email        string `gorm:"column:email"`
	CodeVerified string `gorm:"column:code_verified"`
}
