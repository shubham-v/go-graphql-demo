package users

import (
	"database/sql"
	_ "database/sql"
	database "go-graphql-demo/internal/pkg/db/migrations/postgres"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (user *User) Create() {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
	}
	query := "INSERT INTO Users(Username,Password) VALUES('" + user.Username + "','" + hashedPassword + "" +"');"
	log.Print("Executing Query: ", query)
	statement, err := database.Db.Prepare(query)
	defer statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Print(statement)
	_, err = statement.Exec()
	if err != nil {
		log.Fatal()
	}
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserIdByUsername(username string) (int, error) {
	query := "select ID from Users WHERE Username = '" + username + "';"
	statement, err := database.Db.Prepare(query)
	log.Print(statement)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(statement)
	row := statement.QueryRow()

	var Id int
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}

	return Id, nil
}

func (user *User) Authenticate() bool {
	query := "SELECT Password FROM Users WHERE username = '" + user.Username + "';"
	statement, err := database.Db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(statement)
	row := statement.QueryRow()
	log.Print("Row: ", row)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}

