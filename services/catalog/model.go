package catalogs

import "time"

//Response will recieve the CatalogID when a new Catalog was created
type Response struct {
	ID int64
}

// Catalog gets the basic data of a Catalog
type Catalog struct {
	ID                 int64
	ConfigurationId    int64
	ProjectId          int64
	RedeemTypeId       int64
	ShippingTypeId     int64
	OrdenationTypeId   int64
	AverageFreight     float32
	ConversionRate     float32
	Name               string
	FormatPoints       string
	ReadOnlyAddress    bool
	UseStoreOrdenation bool
	Active             bool
	InsertDate         time.Time
	UpdateDate         time.Time
	HotsiteLoginUrl    string
	RoundDecimalPlaces int8
}
