package model

import "github.com/valyala/fastjson"

// Page struct holds information about pages
type Page struct {
	Category int     `redis:"category_id"`
	Group    int     `redis:"group_id"`
	Page     int     `json:"page" redis:"page"`
	Pages    int     `json:"pages" redis:"pages"`
	Total    int     `json:"total" redis:"total"`
	Ads      []*AdID `json:"ads" redis:"ads"`
}

// PageQuery struct holds information about pages when crawling through pagination.
type PageQuery struct {
	Page     int    `json:"page" redis:"page"`
	Category int    `json:"category" redis:"category"`
	Group    int    `json:"group" redis:"group"`
	Order    string `json:"order" redis:"order"`
	AdType   string `json:"ad_type" redis:"ad_type"`
}

// NewPageQuery creates new PageQuery for given arguments.
func NewPageQuery(page, category, group int) *PageQuery {
	return &PageQuery{
		Page:     page,
		Category: category,
		Group:    group,
		Order:    "newest",
		AdType:   "all",
	}
}

// PageFromJSON parses response from JSON to page struct
// Returns *Page and error
func PageFromJSON(data []byte, parser *fastjson.Parser) (*Page, error) {

	v, err := parser.ParseBytes(data)
	if err != nil {
		return nil, err
	}

	pv := v.Get("search")
	page := pageFromJSON(pv)

	return page, nil
}

// PageQueryFromJSON parses response to page struct.
// Returns *PageQuery and error.
func PageQueryFromJSON(data []byte, parser *fastjson.Parser) (*PageQuery, error) {

	v, err := parser.ParseBytes(data)
	if err != nil {
		return nil, err
	}

	return &PageQuery{
		Page:     v.GetInt("page"),
		Category: v.GetInt("category"),
		Group:    v.GetInt("group"),
		Order:    string(v.GetStringBytes("order")),
		AdType:   string(v.GetStringBytes("ad_type")),
	}, nil
}

// pageFromJSON extracts page's data from JSON.
// Returns *Page.
func pageFromJSON(pv *fastjson.Value) *Page {

	ads := adIDsFromJSON(pv)

	page := &Page{
		Page:  pv.GetInt("page"),
		Pages: pv.GetInt("pages"),
		Total: pv.GetInt("total"),
		Ads:   ads,
	}

	return page
}
