package accountholder

import "time"

// AccountHolder gets the basic data of an accountholder
type AccountHolder []struct {
	ID                           int       `json:"id"`
	ClientID                     int       `json:"clientId"`
	CampaignID                   int       `json:"campaignId"`
	OriginTest                   bool      `json:"originTest"`
	Name                         string    `json:"name"`
	Code                         string    `json:"code"`
	AccountHolderType            int       `json:"accountHolderType"`
	TransactionsExpirationDate   time.Time `json:"transactionsExpirationDate"`
	TransactionsExpirationPeriod int       `json:"transactionsExpirationPeriod"`
	DateInterval                 int       `json:"dateInterval"`
	BillingCode                  string    `json:"billingCode"`
	Status                       int       `json:"status"`
	InsertDate                   time.Time `json:"insertDate"`
	UpdateDate                   time.Time `json:"updateDate"`
	LegacyID                     int       `json:"legacyId"`
	AccountHolderBreakageID      int       `json:"accountHolderBreakageId"`
	SendToProtheus               bool      `json:"sendToProtheus"`
	ApprovalByPass               bool      `json:"approvalByPass"`
	FriendlyName                 string    `json:"friendlyName"`
}
