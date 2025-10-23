package model

type Location struct {
	GeoPoint GeoPoint `json:"geoPoint"`
	Address  Address  `json:"address"`
}

type GeoPoint struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

type Address struct {
	CountryCode      string `json:"countryCode"`
	FormattedAddress string `json:"address"`
}
