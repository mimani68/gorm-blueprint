package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// manager one to many user
// user many to many friend
type User struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Address   []Address
	ManagerID *int
	Manager   []User  `gorm:"foreignkey:ManagerID"`
	Friends   []*User `gorm:"many2many:user_friends;"`
}
type Address struct {
	ID     int `gorm:"primaryKey"`
	Title  string
	UserID int
}

func main() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:1234@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	errDb := db.AutoMigrate(
		&User{},
		&Address{},
	)
	if errDb != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	home := Address{
		ID:    1,
		Title: "home",
	}
	office := Address{
		ID:    2,
		Title: "office",
	}
	louie := User{
		Name: "Louie",
	}
	db.Create(&louie)
	ali := User{
		Name: "Ali",
	}
	db.Create(&ali)
	me := User{
		Name:    "Mahdi",
		Address: []Address{home, office},
		Friends: []*User{&ali},
		Manager: []User{louie},
	}
	db.Create(&me)

	userTmp := User{}
	db.Where("name = ?", "Mahdi").Preload("Address").Preload("Friends").Preload("Manager").Find(&userTmp)
	fmt.Println(&userTmp)
	fmt.Println(&userTmp.Manager)
	fmt.Println(userTmp.Friends[0])

}
