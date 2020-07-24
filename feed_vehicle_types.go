package gbfs

type (
	// FeedVehicleTypes (v2.1-RC)
	FeedVehicleTypes struct {
		FeedCommon
		Data *FeedVehicleTypesData `json:"data"`
	}
	// FeedVehicleTypesData ...
	FeedVehicleTypesData struct {
		VehicleTypes []*FeedVehicleTypesType `json:"vehicle_types"`
	}
	// FeedVehicleTypesType ...
	FeedVehicleTypesType struct {
		VehicleTypeID  string         `json:"vehicle_type_id"`
		FormFactor     FormFactor     `json:"form_factor"`
		PropulsionType PropulsionType `json:"propulsion_type"`
		MaxRangeMeters float64        `json:"max_range_meters,omitempty"`
		Name           string         `json:"name,omitempty"`
	}
)

// Name ...
func (f *FeedVehicleTypes) Name() string {
	return FeedNameVehicleTypes
}
