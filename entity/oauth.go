package entity

import "time"

type OauthClient struct {
	ID           int64  `gorm:"column:id;primary_key"`
	ClientID     string `gorm:"column:client_id"`
	ClientSecret string `gorm:"column:client_secret"`
	Name         string `gorm:"column:name"`
	Redirect     string `gorm:"column:redirect"`
	Description  string `gorm:"column:description"`
	Scope        string `gorm:"column:scope"`
}

type OauthAccessToken struct {
	ID        int64        `gorm:"column:id;primary_key"`
	Client    *OauthClient `gorm:"foreignKey:ClientID;references:ID"`
	ClientID  *int64       `gorm:"column:oauth_client_id"`
	UserID    int64        `gorm:"column:user_id"`
	Token     string       `gorm:"column:token"`
	Scope     string       `gorm:"column:scope"`
	ExpiredAt *time.Time   `gorm:"column:expired_at"`
}
