package models

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var (
	Objects map[string]*Object
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "newadmin"
	dbname   = "Object"
)

type Object struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Phone     string `json:"Phone"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	DoB       string `json:"DoB"`
}

func init() {
	Objects = make(map[string]*Object)

}

func AddOne(object Object) (ObjectId string) {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	fmt.Println("Connected!")
	// close database
	defer db.Close()

	rep := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	phone := rep.MatchString(object.Phone)

	red := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
	date := red.MatchString(object.DoB)

	rem := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	email := rem.MatchString(object.Email)

	hash, _ := HashPassword(object.Password)

	// dynamic
	if email && date && phone {
		insertDynStmt := `insert into "user"("firstname", "lastname", "phone", "email", "password", "dob") values($1, $2, $3, $4, $5, $6)`
		_, e := db.Exec(insertDynStmt, object.FirstName, object.LastName, object.Phone, object.Email, hash, object.DoB)
		if e != nil {
			return "Choose another email"
		}

	} else {
		return "Invalid Inputs"
	}
	return "Successfully created user."
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// func CheckPasswordHash(password, hash string) bool {
//     err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//     return err == nil
// }

func GetOne(ObjectId string) (object *Object, err error) {
	if v, ok := Objects[ObjectId]; ok {
		return v, nil
	}
	return nil, errors.New("ObjectId Not Exist")
}

func GetAll() map[string]*Object {
	return Objects
}

func Delete(ObjectId string) {
	delete(Objects, ObjectId)
}
