package links

import (
	database "go-graphql-demo/internal/pkg/db/migrations/postgres"
	"go-graphql-demo/internal/users"
	"log"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save() int64 {
	query := "INSERT INTO Links(Title, Address, UserID) VALUES('" + link.Title + "', '" + link.Address + "', '" + link.User.ID + "') RETURNING ID"
	log.Print("Executing Query: ", query)
	stmt, err := database.Db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Statement: ", stmt)

	var id int64
	err = stmt.QueryRow().Scan(&id	)
	log.Print("Executed Query: ", query)
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

func GetAll(userId string) []Link {
	stmt, err := database.Db.Prepare("select L.id, L.title, L.address, L.UserID, U.Username from Links L inner join Users U on L.UserID = U.ID where U.ID = '" + userId + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	var username string
	var id string
	var i = 1;
	for rows.Next() {
		i++
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &id, &username)
		link.User = &users.User{
			ID:   id,
			Username: username,
		}
		log.Print("Link: ", link)
		if err != nil{
			log.Fatal(err)
		}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	log.Print("Returning Links", links)
	return links
}

