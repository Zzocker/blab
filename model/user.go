package model

type User struct {
	Username  string        `json:"username" bson:"username"`
	Details   UserDetails   `json:"details" bson:"details"`
	Contacts  []UserContact `json:"contacts" bson:"contacts"`
	Rating    UserRating    `json:"rating" bson:"rating"`
	CreatedOn int64         `json:"created_on" bson:"created_on"`
	Password  string        `json:"-" bson:"password"`
}

const (
	UserContactEmail = "email"
)

type UserContact struct {
	Type  string `json:"type" bson:"type"`
	Value string `json:"value" bson:"value"`
}

type UserRating struct {
	AsSeller   Rating `json:"as_seller" bson:"as_seller"`
	AsBorrower Rating `json:"as_borrower" bson:"as_borrower"`
}

type UserDetails struct {
	Name   string `json:"name" bson:"name"`
	Age    int    `json:"age" bson:"age"`
	Gender string `json:"gender"`
}
