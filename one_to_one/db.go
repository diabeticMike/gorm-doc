package one_to_one

import "gorm.io/gorm"

func New(db *gorm.DB) *OneToOne {
	db.AutoMigrate(&User{}, &CreditCard{})
	return &OneToOne{db: db}
}

type OneToOne struct {
	db *gorm.DB
}

func (o *OneToOne) Create(user User) error {
	return o.db.Create(&user).Error
}

func (o *OneToOne) GetByName(name string) (User, error) {
	var user User
	err := o.db.Table("users").Preload("CreditCard").First(&user, "name = ?", name).Error
	return user, err
}

func (o *OneToOne) GetByNameWithoutRelations(name string) (User, error) {
	var user User
	err := o.db.First(&user, "name = ?", name).Error
	return user, err
}
