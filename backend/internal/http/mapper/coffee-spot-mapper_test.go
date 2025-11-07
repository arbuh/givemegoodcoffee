package mapper_test

import (
	"givemegoodcoffee/internal/http/mapper"
	"givemegoodcoffee/internal/http/request"
	"givemegoodcoffee/internal/http/response"
	"givemegoodcoffee/internal/model"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

var coffeeSpot *model.CoffeeSpot = &model.CoffeeSpot{
	ID:   uuid.MustParse("06217c20-8782-4ed3-801f-b582e91fb8a7"),
	Name: "Koefie Coffee",
	Type: model.CoffeeShop,
	Location: model.Location{
		GeoPoint: model.GeoPoint{
			Lat: "0.0",
			Lon: "1.0",
		},
		Address: model.Address{
			CountryCode:      "nl",
			FormattedAddress: "Jan de Vriesstraat 7, Uden",
		},
	},
}

func TestCoffeeSpotMapper_ToResponse(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		spot *model.CoffeeSpot
		want *response.CoffeeSpotResponse
	}{
		{
			name: "map coffee spot model to response",
			spot: coffeeSpot,
			want: &response.CoffeeSpotResponse{
				Name: "Koefie Coffee",
				Type: "coffeeShop",
				Location: response.LocationResponse{
					CountryCode: "nl",
					Address:     "Jan de Vriesstraat 7, Uden",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mapper.NewCoffeeSpotMapper()
			got := m.ToResponse(tt.spot)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoffeeSpotMapper_FromRequest(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		request *request.CoffeeSpotRequest
		want    *model.CoffeeSpot
		wantErr bool
	}{
		{
			name: "map request to coffee spot",
			request: &request.CoffeeSpotRequest{
				Name: "Koefie Coffee",
				Type: "coffeeShop",
				Location: request.LocationRequest{
					CountryCode: "nl",
					Address:     "Jan de Vriesstraat 7, Uden",
				},
			},
			want:    coffeeSpot,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mapper.NewCoffeeSpotMapper()
			got, gotErr := m.FromRequest(tt.request)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("FromRequest() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("FromRequest() succeeded unexpectedly")
			}
			if !areCoffeeSpotsEqual(got, tt.want) {
				t.Errorf("FromRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func areCoffeeSpotsEqual(cs1, cs2 *model.CoffeeSpot) bool {
	// The fields we exclude from comparison
	cs2.ID = cs1.ID
	cs2.Location.GeoPoint = cs1.Location.GeoPoint

	return reflect.DeepEqual(cs1, cs2)
}
