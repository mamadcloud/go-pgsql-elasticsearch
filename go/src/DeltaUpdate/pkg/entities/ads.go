package entities

import (
	"time"
)

type Ads struct {
	ID int `json:"adsId"`
	ListID *int `json:"listId"`
	OriginalListTime *time.Time `json:"originalListTime"`
	ListTime *time.Time `json:"listTime"`
	Status string `json:"status"`
	Type string `json:"type"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Region *int `json:"region"`
	City *int `json:"city"`
	Category *int `json:"category"` 
	UserID *int `json:"userId"`
	Password string `json:"-"`
	PhoneHidden *bool `json:"phoneHidden"`
	NoSalesmen *bool `json:"noSalesmen"`
	CompanyAd *bool `json:"companyAd"`
	Subject string `json:"subject"`
	Body string `json:"body"`
	Price *float64 `json:"price"`
	OldPrice *float64 `json:"oldPrice"`
	Image *string `json:"image"`
	InfoPage string `json:"infoPage"`
	InfoPageTitle string `json:"infoPageTitle`
	StoreID *int `json:"storeId"`
	SaltedPassword string `json:"-"`
	ModifiedAt *time.Time `json:"modifiedAt"`
	AdType string `json:"adType"`
	AdParams []AdParam `json:"adParams"`
}