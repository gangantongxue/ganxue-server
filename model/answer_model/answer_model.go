package answer_model

type Answer struct {
	ID     string `json:"id" bson:"id"`
	Input  string `json:"input" bson:"input"`
	Output string `json:"output" bson:"output"`
}
