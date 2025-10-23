package model

type Location struct {
	GeoPoint GeoPoint `json:"geoPoint"`
	Address  Address  `json:"address"`
}

type GeoPoint struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type Address struct {
	CountryCode      string `json:"countryCode"`
	FormattedAddress string `json:"address"`
}
