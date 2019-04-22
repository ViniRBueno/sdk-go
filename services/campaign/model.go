package campaign

import "time"

//Response will recieve the CampaignID when a new Campaign was created
type Response struct {
	ID int64
}

// Campaign gets the basic data of a campaign
type Campaign struct {
	ID                          int64
	ClientIDDefault             int64
	ProjectCellID               int64
	ContactUsConfigurationID    int64
	ProjectErpInfoID            int64
	LogonTypeID                 int64
	ProjectTypeID               int64
	FinancialDataTypeID         int64
	FirstAccessTypeID           int64
	ExternalPoints              bool
	FixedTax                    float32
	MinimalMonthlyBilling       float32
	MonthlyAwardForecast        float32
	TotalAwardForecast          float32
	Name                        string
	LoginURL                    string
	SalesForceCode              string
	CypherType                  int64
	IsAuthenticatedFromHotsite  bool
	StartDate                   time.Time
	EndDate                     time.Time
	Active                      bool
	InsertDate                  time.Time
	UpdateDate                  time.Time
	IsHybridAccount             bool
	RateRuleID                  int64
	Additional                  int
	CampaignTypeID              int8
	Invoice                     bool
	SalesForceIntegrationTypeID int
	UsingMobileApp              bool
	EnableCommunPopUp           bool
	ProjectIconURL              string
	CallCenterTypeID            int8
}
