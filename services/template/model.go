package template

import "time"

//Response will recieve the CampaignID when a new Campaign was created
type Response struct {
	ID int64
}

// Template gets the basic data of a template
type Template struct {
	ID                       int64
	catalogID                int64
	campaignID               int64
	body                     string
	Subject                  string
	NotificationTypeID       int
	TemplateTypeID           float32
	TemplateVendorProviderID float32
	FromAddress              string
	FromName                 string
	Active                   bool
	InsertDate               time.Time
	UpdateDate               time.Time
}
