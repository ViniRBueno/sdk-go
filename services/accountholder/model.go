package accountholder

import "time"

// AccountHolder gets the basic data of an accountholder
type AccountHolder []struct {
	ID                           int64
	ClientID                     int64
	CampaignID                   int64
	OriginTest                   bool
	Name                         string
	Code                         string
	AccountHolderType            int
	TransactionsExpirationDate   time.Time
	TransactionsExpirationPeriod int
	DateInterval                 bool
	FixedTax                     time.Time
	BillingCode                  string
	Status                       int
	InsertDate                   time.Time
	UpdateDate                   time.Time
	LegacyID                     int64
	AccountHolderBreakageID      int64
	SendToProtheus               bool
	ApprovalByPass               bool
	FriendlyName                 string
}
