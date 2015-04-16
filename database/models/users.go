package models

// User represents a user in the system
type User struct {
	ID        int    `json:"id"`
	UID       string `gorm:"column:uid" json:"uid"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
