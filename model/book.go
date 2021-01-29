package model

import "time"

type Book struct {
	ISBN      string        `json:"isbn" bson:"isbn"`
	Details   BookDetails   `json:"details" bson:"details"`
	Ownership BookOwnership `json:"ownership" bson:"ownership"`
	Rating    Rating        `json:"rating" bson:"rating"`
	CreatedOn time.Time     `json:"created_on" bson:"created_on"`
}

type bookGenre string

var (
	GenreList = map[bookGenre]string{
		// bookGenre("fun"): "fun",
	}
)

type BookDetails struct {
	Author string      `json:"author" bson:"author"`
	Genre  []bookGenre `json:"genre" bson:"genre"`
}

type BookOwnership struct {
	Owner   string `json:"owner" bson:"owner"`
	Current string `json:"current" bson:"current"`
}
