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
	statement, err := database.Db.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := statement.Exec(link.Title, link.Address)
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

// GetAll - Get all links
func GetAll() []Link {
	statement, err := database.Db.Prepare("SELECT id, title, address FROM Links")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	rows, err := statement.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
