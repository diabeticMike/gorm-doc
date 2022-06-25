package one_to_many

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

// User has one CreditCard, CreditCardID is the foreign key
type User struct {
	gorm.Model
	Name        string
	CreditCards []CreditCard
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
