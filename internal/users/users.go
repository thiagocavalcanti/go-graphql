package users

import (
	"database/sql"
	"log"

	"github.com/thiagocavalcanti/gqlgen-handson/graph/model"
	database "github.com/thiagocavalcanti/gqlgen-handson/internal/pkg/db/migrations/mysql"
	"golang.org/x/crypto/bcrypt"
)

// User - overrides user model
type User struct {
	model.User
}

// Create - Creates a new user into the database
func (user *User) Create() {
	statement, err := database.Db.Prepare("INSERT INTO Users(Name,Password) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}

	hashedPassword, err := HashPassword(user.Password)
	_, err = statement.Exec(user.Name, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

// HashPassword - Creates a hashed password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash - Check the password hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetUserIDByUsername - Get the userID by its username
func GetUserIDByUsername(username string) (int, error) {
	statement, err := database.Db.Prepare("SELECT ID FROM Users WHERE Name = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(username)

	var ID int
	err = row.Scan(&ID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}

	return ID, nil
}
