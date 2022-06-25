package one_to_many

import "gorm.io/gorm"

func New(db *gorm.DB) *OneToMany {
	db.AutoMigrate(&User{}, &CreditCard{})
	return &OneToMany{db: db}
}

type OneToMany struct {
	db *gorm.DB
}

func (o *OneToMany) Create(user User) error {
	return o.db.Create(&user).Error
}

func (o *OneToMany) GetByName(name string) (User, error) {
	var user User
	err := o.db.Table("users").Preload("CreditCards").First(&user, "name = ?", name).Error
	return user, err
}

func (o *OneToMany) GetByNameWithoutRelations(name string) (User, error) {
	var user User
	err := o.db.First(&user, "name = ?", name).Error
	return user, err
}
