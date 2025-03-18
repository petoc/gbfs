package gbfs

type (
	FeedVehicleStatus struct {
		FeedCommon
		Data *FeedVehicleStatusData `json:"data"`
	}
	FeedVehicleStatusData struct {
		Vehicles []*FeedVehicleStatusVehicle `json:"vehicles"`
	}
	FeedVehicleStatusVehicle struct {
		VehicleID          *ID         `json:"vehicle_id"`
		Lat                *Coordinate `json:"lat"`
		Lon                *Coordinate `json:"lon"`
		IsReserved         *Boolean    `json:"is_reserved"`
		IsDisabled         *Boolean    `json:"is_disabled"`
		RentalURIs         *RentalURIs `json:"rental_uris,omitempty"`
		VehicleTypeID      *ID         `json:"vehicle_type_id,omitempty"`
		LastReported       *Timestamp  `json:"last_reported,omitempty"`
		CurrentRangeMeters *float64    `json:"current_range_meters,omitempty"`
		CurrentFuelPercent *float64    `json:"current_fuel_percent,omitempty"`
		StationID          *ID         `json:"station_id,omitempty"`
		HomeStationID      *ID         `json:"home_station_id,omitempty"`
		PricingPlanID      *ID         `json:"pricing_plan_id,omitempty"`
		VehicleEquipment   []string    `json:"vehicle_equipment,omitempty"`
		AvailableUntil     *string     `json:"available_until,omitempty"`
	}
)

func (f *FeedVehicleStatus) Name() string {
	return FeedNameVehicleStatus
}
