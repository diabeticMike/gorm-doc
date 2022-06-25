package one_to_one

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

// User has one CreditCard, CreditCardID is the foreign key
type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}

func (u User) String() string {
	body, _ := json.Marshal(&u)
	return fmt.Sprintf(string(body))
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}
