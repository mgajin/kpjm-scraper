package model

import (
	"github.com/valyala/fastjson"
)

// Category struct holds base information about categories.
type Category struct {
	ID      int    `json:"category_id"`
	AdCount int    `json:"ad_count"`
	Parent  int    `json:"parent"`
	Name    string `json:"name"`
	AdKind  string `json:"ad_kind"`
	Ads     *[]AdID
}

// CategoryFilter struct holds information for filtered ads.
// It is used for crawling through category pages.
type CategoryFilter struct {
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Total int `json:"total"`
	Ads   *[]AdID
}

// CategoriesFromJSON parses categories from JSON.
// Returns pointer to slice of categories and error.
func CategoriesFromJSON(data []byte, parser *fastjson.Parser) (*[]Category, error) {
	v, err := parser.ParseBytes(data)
	if err != nil {
		return nil, err
	}

	var categories []Category

	for _, category := range v.GetArray("categories") {
		categories = append(categories, *newCategory(category))
	}

	categories = *filterCategories(&categories)

	return &categories, nil
}

// CategoryFilterFromJSON parses filtered ads from JSON.
// Returns *CategoryFilter and error.
func CategoryFilterFromJSON(data []byte, parser *fastjson.Parser) (*CategoryFilter, error) {
	v, err := parser.ParseBytes(data)
	if err != nil {
		return nil, err
	}

	v = v.Get("search")

	return newCategoryFilter(v), nil
}

// filterCategories returns only categories from given slice.
// Returns pointer to slice of categories.
func filterCategories(all *[]Category) *[]Category {
	var categories []Category

	for _, category := range *all {
		if category.Parent == 0 {
			categories = append(categories, category)
		}
	}

	return &categories
}

// newCategory crates new Category from JSON.
// Returns *Category.
func newCategory(v *fastjson.Value) *Category {
	return &Category{
		ID:      v.GetInt("category_id"),
		AdCount: v.GetInt("ad_count"),
		Parent:  v.GetInt("parent"),
		Name:    string(v.GetStringBytes("name")),
		AdKind:  string(v.GetStringBytes("ad_kind")),
	}
}

// newCategoryFilter creates new Category Filter from JSON.
// Returns *CategoryFilter.
func newCategoryFilter(v *fastjson.Value) *CategoryFilter {
	return &CategoryFilter{
		Page:  v.GetInt("page"),
		Pages: v.GetInt("pages"),
		Total: v.GetInt("total"),
		Ads:   IDsFromJSON(v),
	}
}
