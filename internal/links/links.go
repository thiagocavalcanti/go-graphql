package links

import (
	"log"

	"github.com/thiagocavalcanti/gqlgen-handson/graph/model"
	database "github.com/thiagocavalcanti/gqlgen-handson/internal/pkg/db/migrations/mysql"
)

// Link - overrides the link model
type Link struct {
	model.Link
}

// Save - Save new link into database
func (link Link) Save() int64 {
	statement, err := database.Db.Prepare("INSERT INTO Links(Title,Address, UserID) VALUES(?,?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := statement.Exec(link.Title, link.Address, link.User.ID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}
