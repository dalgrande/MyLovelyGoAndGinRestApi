package models //models/UserModel.go

// User type
type User struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// TableName s
func (b *User) TableName() string {
	return "user"
}
