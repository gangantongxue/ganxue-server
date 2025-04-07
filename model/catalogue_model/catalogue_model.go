package catalogue_model

type Subject map[string]string

type CatalogueModel struct {
	Golang Subject `yaml:"golang" json:"golang" bson:"golang"`
	C      Subject `yaml:"c" json:"c" bson:"c"`
	Cpp    Subject `yaml:"cpp" json:"cpp" bson:"cpp"`
}
