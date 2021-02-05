package model

// User :
type User struct {
	Username  string        `json:"username" bson:"username"`
	Details   UserDetails   `json:'details" bson:'details"`
	Contacts  []UserContact `json:"contacts" bson:"contacts"`
	Rating    UserRating    `json:"rating" bson:"rating"`
	CreatedOn int64         `json:"-" bson:"created_on"`
	Password  string        `json:"-" bson:"password"`
}

// UserDetails :
type UserDetails struct {
	Name   string     `json:"name" bson:"name"`
	DOB    int64      `json:"dob" bson:"dob"`
	Gender GenderType `json:"gender" bson:"gender"`
	Img    UserImages `json:"img" bson:"img"`
}

type UserImages struct {
	ProfileID string `json:"profile_id" bson:"profile_id"`
	CoverID   string `json:"cover_id" bson:"cover_id"`
}

type GenderType uint8

const (
	GenderTypeFemale GenderType = iota + 1
	GenderTypeMale
)

type UserContact struct {
	Type  string `json:"type" bson:"type"`
	Value string `json:"value" bson:"value"`
}

type UserContactType uint8

const (
	UserContactTypeEmail UserContactType = iota + 1
)

type UserRating struct {
	AsSeller   Rating `json:"as_seller" bson:"as_seller"`
	AsBorrower Rating `json:"as_borrower" bson:"as_borrower"`
}
