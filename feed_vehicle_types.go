package gbfs

type (
	// FeedVehicleTypes (v2.1-RC)
	FeedVehicleTypes struct {
		FeedCommon
		Data *FeedVehicleTypesData `json:"data"`
	}
	// FeedVehicleTypesData ...
	FeedVehicleTypesData struct {
		VehicleTypes []*FeedVehicleTypesVehicleType `json:"vehicle_types"`
	}
	// FeedVehicleTypesVehicleType ...
	FeedVehicleTypesVehicleType struct {
		VehicleTypeID  string  `json:"vehicle_type_id"`
		FormFactor     string  `json:"form_factor"`
		PropulsionType string  `json:"propulsion_type"`
		MaxRangeMeters float64 `json:"max_range_meters,omitempty"`
		Name           string  `json:"name,omitempty"`
	}
)

// Name ...
func (f *FeedVehicleTypes) Name() string {
	return FeedNameVehicleTypes
}
