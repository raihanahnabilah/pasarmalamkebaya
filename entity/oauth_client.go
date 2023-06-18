package entity

type OauthClient struct {
	ID           int64  `gorm:"column:id;primary_key"`
	ClientID     string `gorm:"column:client_id"`
	ClientSecret string `gorm:"column:client_secret"`
	Name         string `gorm:"column:name"`
	Redirect     string `gorm:"column:redirect"`
	Description  string `gorm:"column:description"`
	Scope        string `gorm:"column:scope"`
}
