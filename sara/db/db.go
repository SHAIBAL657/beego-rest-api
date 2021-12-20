package db

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

type User struct {
	Id        int    `orm:"auto";primary_key`
	Firstname string `json:"firstname" orm:"size(40)`
	Lastname  string `json:"lastname" orm:"size(40)`
	Phone     string `json:"phone" orm:"size(50)`
	Email     string `json:"email" orm:"size(64);unique"`
	Password  string `json:"password"`
	Dob       string `json:"dob"`
}

type Orm struct {
	CreatedORM orm.Ormer
}

func InitDB() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=newadmin host=127.0.0.1 port=5432 dbname=Object sslmode=disable")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}
