package md_model

type Markdown struct {
	ID      string `json:"id" bson:"id"`
	Content string `json:"content" bson:"content"`
	Code    string `json:"code" bson:"code"`
}
