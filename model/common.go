package model

// Rating :
// Count is number rating received
// Value = (currentValue*currentCount+ratingReceived)/(currentCount+1)
type Rating struct {
	Count int64   `json:"count" bson:"count"`
	Value float64 `json:"value" bson:"value"`
}
