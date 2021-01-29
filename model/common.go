package model

const (
	voteOutOf = 10
)

// Rating is comment struct which represents a rating
// Count = total rate received
// Value is actual rating out 10
type Rating struct {
	Count int64   `json:"count" bson:"count"`
	Value float64 `json:"value" bson:"value"`
}
