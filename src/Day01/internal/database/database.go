package database

type Ingredient struct {
	Name  *string `json:"ingredient_name"           xml:"itemname"`
	Count *string `json:"ingredient_count"          xml:"itemcount"`
	Unit  *string `json:"ingredient_unit,omitempty" xml:"itemunit"`
}

type Cake struct {
	Name        *string      `json:"name"        xml:"name"`
	Time        *string      `json:"time"        xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Recipes struct {
	Cakes []Cake `json:"cake" xml:"cake"`
}

func (c Cake) getName() string {
	name := "unknown"
	if c.Name != nil {
		name = *c.Name
	}
	return name
}
