package model

import "github.com/valyala/fastjson"

// Category struct holds information about categories / groups.
type Category struct {
	ID      int    `json:"category_id" redis:"id"`
	AdCount int    `json:"ad_count" redis:"ad_count"`
	Parent  int    `json:"parent" redis:"parent"`
	Name    string `json:"name" redis:"name"`
	AdKind  string `json:"ad_kind" redis:"ad_kind"`
}

// CategoriesFromJSON parses response from JSON to slice of categories.
// Returns []*Category and error.
func CategoriesFromJSON(data []byte, parser *fastjson.Parser) ([]*Category, error) {

	v, err := parser.ParseBytes(data)
	if err != nil {
		return nil, err
	}

	var categories []*Category

	for _, cv := range v.GetArray("categories") {
		category := categoryFromJSON(cv)
		categories = append(categories, category)
	}

	return categories, nil
}

// categoryFromJSON extracts category's data from JSON.
// Returns *Category.
func categoryFromJSON(cv *fastjson.Value) *Category {

	category := &Category{
		ID:      cv.GetInt("category_id"),
		AdCount: cv.GetInt("ad_count"),
		Parent:  cv.GetInt("parent"),
		Name:    string(cv.GetStringBytes("name")),
		AdKind:  string(cv.GetStringBytes("ad_kind")),
	}

	return category
}
