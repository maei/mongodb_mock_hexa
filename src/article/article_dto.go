package article

type Article struct {
	Name     string `json:"name" bson:"name"`
	Quantity int64  `json:"quantity" bson:"quantity"`
}
