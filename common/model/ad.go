package model

import "github.com/valyala/fastjson"

// Ad struct holds information about ads.
type Ad struct {
	ID            int    `json:"ad_id" bson:"ad_id" redis:"id"`
	CategoryID    int    `json:"category_id" bson:"category_id" redis:"category_id"`
	GroupID       int    `json:"group_id" bson:"group_id" redis:"group_id"`
	LocationID    int    `json:"location_id" bson:"location_id" redis:"location_id"`
	Price         int    `json:"price" bson:"price" redis:"price"`
	ViewCount     int    `json:"view_count" bson:"view_count" redis:"view_count"`
	FavoriteCount int    `json:"favorite_count" bson:"favorite_count" redis:"favorite_count"`
	UserID        int    `json:"user_id" bson:"user_id" redis:"user_id"`
	URL           string `json:"ad_url" bson:"ad_url" redis:"ad_url"`
	Name          string `json:"name" bson:"name" redis:"name"`
	CategoryName  string `json:"category_name" bson:"category_name" redis:"category_name"`
	GroupName     string `json:"group_name" bson:"group_name" redis:"group_name"`
	Description   string `json:"description" bson:"description" redis:"description"`
	Condition     string `json:"condition" bson:"condition" redis:"condition"`
	Kind          string `json:"ad_kind" bson:"ad_kind" redis:"ad_kind"`
	Type          string `json:"ad_type" bson:"ad_type" redis:"ad_type"`
	Owner         string `json:"owner" bson:"owner" redis:"owner"`
	Phone         string `json:"phone" bson:"phone" redis:"phone"`
	Posted        string `json:"posted" bson:"posted" redis:"posted"`
	Currency      string `json:"currency" bson:"currency" redis:"currency"`
	LocationName  string `json:"location_name" bson:"location_name" redis:"location_name"`
}


// AdID represents ad's ID
type AdID = int

// AdFromJSON parses response from JSON to ad struct
// Returns *Ad and error
func AdFromJSON(data []byte, parser *fastjson.Parser) (*Ad, error) {

	v, err := parser.ParseBytes(data)
	if err != nil {
		return nil, err
	}

	av := v.Get("ad")
	ad := adFromJSON(av)

	return ad, nil
}

// adFromJSON extracts ad's data from JSON.
// Returns *Ad
func adFromJSON(av *fastjson.Value) *Ad {

	ad := &Ad{
		ID:            av.GetInt("ad_id"),
		CategoryID:    av.GetInt("category_id"),
		GroupID:       av.GetInt("group_id"),
		LocationID:    av.GetInt("location_id"),
		UserID:        av.GetInt("user_id"),
		Price:         av.GetInt("price"),
		ViewCount:     av.GetInt("view_count"),
		FavoriteCount: av.GetInt("favorite_count"),
		Name:          string(av.GetStringBytes("name")),
		CategoryName:  string(av.GetStringBytes("category_name")),
		GroupName:     string(av.GetStringBytes("group_name")),
		LocationName:  string(av.GetStringBytes("location_name")),
		Condition:     string(av.GetStringBytes("condition")),
		Currency:      string(av.GetStringBytes("currency")),
		Posted:        string(av.GetStringBytes("posted")),
		Kind:          string(av.GetStringBytes("ad_kind")),
		Type:          string(av.GetStringBytes("ad_type")),
		Owner:         string(av.GetStringBytes("owner")),
		Phone:         string(av.GetStringBytes("phone")),
	}

	return ad
}

// adIDsFromJSON extracts ad ids from JSON
// Returns []*AdID
func adIDsFromJSON(pv *fastjson.Value) []*AdID {

	var ids []*AdID

	for _, av := range pv.GetArray("ads") {
		id := av.GetInt("ad_id")
		ids = append(ids, &id)
	}

	return ids
}
