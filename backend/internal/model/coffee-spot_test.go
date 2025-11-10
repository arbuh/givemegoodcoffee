package model_test

import (
	"givemegoodcoffee/internal/model"
	"testing"
)

func TestCoffeeSpotTypeFromString(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s       string
		want    model.CoffeeSpotType
		wantErr bool
	}{
		{
			name: "invalidate unknown coffee spot types",
			s:    "eatery",
			want: "",
			wantErr: true,
		},
		{
			name: "validate cafe",
			s:    "cafe",
			want: model.Cafe,
			wantErr: false,
		},
		{
			name: "validate cafe",
			s:    "coffeeShop",
			want: model.CoffeeShop,
			wantErr: false,
		},
		{
			name: "validate cafe",
			s:    "restaurant",
			want: model.Restraurant,
			wantErr: false,
		},
		{
			name: "validate cafe",
			s:    "bakery",
			want: model.Bakery,
			wantErr: false,
		},
		{
			name: "validate cafe",
			s:    "gasStation",
			want: model.GasStation,
			wantErr: false,
		},
		{
			name: "validate cafe",
			s:    "other",
			want: model.Other,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := model.CoffeeSpotTypeFromString(tt.s)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CoffeeSpotTypeFromString() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CoffeeSpotTypeFromString() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("CoffeeSpotTypeFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

