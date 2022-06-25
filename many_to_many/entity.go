package many_to_many

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

// User has one CreditCard, CreditCardID is the foreign key
type User struct {
	gorm.Model
	Name  string
	Teams []Team `gorm:"many2many:users_teams" json:"teams,omitempty"`
}

func (u User) String() string {
	body, _ := json.Marshal(&u)
	return fmt.Sprintf(string(body))
}

type Team struct {
	gorm.Model
	Name  string
	Users []User `gorm:"many2many:users_teams"  json:"users,omitempty"`
}

func (t Team) String() string {
	body, _ := json.Marshal(&t)
	return fmt.Sprintf(string(body))
}

type UsersTeams struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	TeamID uint
}
