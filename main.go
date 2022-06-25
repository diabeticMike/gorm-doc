package main

import (
	"fmt"
	"github.com/gorm-doc/many_to_many"
	"github.com/gorm-doc/one_to_many"
	"github.com/gorm-doc/one_to_one"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := NewConn()
	if err != nil {
		panic(err)
	}

	//DoOneToOne(db)
	//fmt.Println()
	//DoOneToMany(db)
	DoManyToMany(db)
}

func NewConn() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("postgres://postgres:secret@localhost:5432/test?sslmode=disable"), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	return db, err
}

func DoOneToOne(db *gorm.DB) {
	name := "user 1"
	oto := one_to_one.New(db)
	err := oto.Create(one_to_one.User{
		Name:       name,
		CreditCard: one_to_one.CreditCard{Number: "some number"},
	})
	if err != nil {
		panic(err)
	}

	user, err := oto.GetByName(name)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}

func DoOneToMany(db *gorm.DB) {
	name := "user 2"
	oto := one_to_many.New(db)
	err := oto.Create(one_to_many.User{
		Name:        name,
		CreditCards: []one_to_many.CreditCard{{Number: "some number 1"}, {Number: "some number 2"}},
	})
	if err != nil {
		panic(err)
	}

	user, err := oto.GetByName(name)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}

func DoManyToMany(db *gorm.DB) {
	mtm := many_to_many.New(db)
	userName1 := "MTMUser1"
	userName2 := "MTMUser2"
	teamName1 := "MTMTeam1"
	teamName2 := "MTMTeam2"
	err := mtm.CreateUser(many_to_many.User{
		Name: userName1,
		Teams: []many_to_many.Team{
			{
				Model: gorm.Model{ID: 10},
				Name:  teamName1,
			},
			{
				Model: gorm.Model{ID: 11},
				Name:  teamName2,
			},
		},
	})

	err = mtm.CreateUser(many_to_many.User{
		Name: userName2,
		Teams: []many_to_many.Team{
			{
				Model: gorm.Model{ID: 10},
				Name:  teamName1,
			},
			{
				Model: gorm.Model{ID: 11},
				Name:  teamName2,
			},
		},
	})

	if err != nil {
		panic(err)
	}

	user, err := mtm.GetUserByName(userName1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	user, err = mtm.GetUserByName(userName2)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	team, err := mtm.GetTeamByName(teamName1)
	if err != nil {
		panic(err)
	}

	fmt.Println(team)
}
