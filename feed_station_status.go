package gbfs

type (
	// FeedStationStatus ...
	FeedStationStatus struct {
		FeedCommon
		Data *FeedStationStatusData `json:"data"`
	}
	// FeedStationStatusData ...
	FeedStationStatusData struct {
		Stations []*FeedStationStatusStation `json:"stations"`
	}
	// FeedStationStatusStation ...
	FeedStationStatusStation struct {
		StationID             string                          `json:"station_id"`
		NumBikesAvailable     int64                           `json:"num_bikes_available"`
		NumBikesDisabled      int64                           `json:"num_bikes_disabled,omitempty"`
		NumDocksAvailable     int64                           `json:"num_docks_available,omitempty"` // conditionally required (v2.0)
		NumDocksDisabled      int64                           `json:"num_docks_disabled,omitempty"`
		IsInstalled           Boolean                         `json:"is_installed"`
		IsRenting             Boolean                         `json:"is_renting"`
		IsReturning           Boolean                         `json:"is_returning"`
		LastReported          Timestamp                       `json:"last_reported"`
		VehicleDocksAvailable []*FeedStationStatusVehicleDock `json:"vehicle_docks_available,omitempty"` // (v2.1-RC)
		Vehicles              []*FeedStationStatusVehicle     `json:"vehicles,omitempty"`                // (v2.1-RC)
	}
	// FeedStationStatusVehicleDock ...
	FeedStationStatusVehicleDock struct {
		VehicleTypeIDs []string `json:"rental_methods"` // (v2.1-RC)
		Count          int64    `json:"count"`          // (v2.1-RC)
	}
	// FeedStationStatusVehicle ...
	FeedStationStatusVehicle struct {
		BikeID             string  `json:"bike_id"`                        // (v2.1-RC)
		IsReserved         Boolean `json:"is_reserved"`                    // (v2.1-RC)
		IsDisabled         Boolean `json:"is_disabled"`                    // (v2.1-RC)
		VehicleTypeID      string  `json:"vehicle_type_id"`                // (v2.1-RC)
		CurrentRangeMeters float64 `json:"current_range_meters,omitempty"` // (v2.1-RC)
	}
)

// Name ...
func (f *FeedStationStatus) Name() string {
	return FeedNameStationStatus
}
