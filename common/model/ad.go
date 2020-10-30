package model

import (
	"github.com/valyala/fastjson"
)

// Ad struct holds information of single ad.
type Ad struct {
	ID            int    `json:"ad_id"`
	CategoryID    int    `json:"category_id"`
	GroupID       int    `json:"group_id"`
	LocationID    int    `json:"location_id"`
	Price         int    `json:"price"`
	ViewCount     int    `json:"view_count"`
	FavoriteCount int    `json:"favorite_count"`
	UserID        int    `json:"user_id"`
	URL           string `json:"ad_url"`
	Name          string `json:"name"`
	CategoryName  string `json:"category_name"`
	GroupName     string `json:"group_name"`
	Description   string `json:"description"`
	Condition     string `json:"condition"`
	Kind          string `json:"ad_kind"`
	Type          string `json:"ad_type"`
	Owner         string `json:"owner"`
	Phone         string `json:"phone"`
	Posted        string `json:"posted"`
	Currency      string `json:"currency"`
	LocationName  string `json:"location_name"`
}

// AdID type is used for parsing Ad ID's when crawling through categories
type AdID = int

// AdFromJSON extracts Ad data from JSON
// Returns *Ad and error
func AdFromJSON(data []byte, parser *fastjson.Parser) (*Ad, error) {
	v, err := parser.ParseBytes(data)
	if err != nil {
		return nil, err
	}

	v = v.Get("ad")

	return newAd(v), nil
}

// IDsFromJSON extracts Ads from JSON
// Returns pointer to slice of ad ID's
func IDsFromJSON(v *fastjson.Value) []*AdID {
	var ads []*AdID

	for _, ad := range v.GetArray("ads") {
		id := ad.GetInt("ad_id")
		ads = append(ads, &id)
	}

	return ads
}

// newAd extracts Ad data from JSON
// Returns *Ad
func newAd(v *fastjson.Value) *Ad {
	return &Ad{
		ID:            v.GetInt("ad_id"),
		CategoryID:    v.GetInt("category_id"),
		GroupID:       v.GetInt("group_id"),
		LocationID:    v.GetInt("location_id"),
		UserID:        v.GetInt("user_id"),
		Price:         v.GetInt("price"),
		ViewCount:     v.GetInt("view_count"),
		FavoriteCount: v.GetInt("favorite_count"),
		Name:          string(v.GetStringBytes("name")),
		CategoryName:  string(v.GetStringBytes("category_name")),
		GroupName:     string(v.GetStringBytes("group_name")),
		LocationName:  string(v.GetStringBytes("location_name")),
		Condition:     string(v.GetStringBytes("condition")),
		Currency:      string(v.GetStringBytes("currency")),
		Posted:        string(v.GetStringBytes("posted")),
		Kind:          string(v.GetStringBytes("ad_kind")),
		Type:          string(v.GetStringBytes("ad_type")),
		Owner:         string(v.GetStringBytes("owner")),
		Phone:         string(v.GetStringBytes("phone")),
	}
}
