package model

type Location struct {
	GeoPoint GeoPoint
	Address  Address
}

type GeoPoint struct {
	Lat string
	Lon string
}

type Address struct {
	CountryCode      string
	FormattedAddress string
}
