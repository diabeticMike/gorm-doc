package many_to_many

import "gorm.io/gorm"

func New(db *gorm.DB) *OneToMany {
	db.AutoMigrate(&UsersTeams{}, &User{}, &Team{})
	return &OneToMany{db: db}
}

type OneToMany struct {
	db *gorm.DB
}

func (o *OneToMany) CreateUser(user User) error {
	return o.db.Create(&user).Error
}

func (o *OneToMany) CreateTeam(team Team) error {
	return o.db.Create(&team).Error
}

func (o *OneToMany) GetUserByName(name string) (User, error) {
	var user User
	err := o.db.Table("users").Preload("Teams").Last(&user, "name = ?", name).Error
	return user, err
}

func (o *OneToMany) GetTeamByName(name string) (Team, error) {
	var team Team
	err := o.db.Table("teams").Preload("Users").Last(&team, "name = ?", name).Error
	return team, err
}
