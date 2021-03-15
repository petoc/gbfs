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
		StationID             *ID                             `json:"station_id"`
		NumBikesAvailable     *int64                          `json:"num_bikes_available"`
		NumBikesDisabled      *int64                          `json:"num_bikes_disabled,omitempty"`
		NumDocksAvailable     *int64                          `json:"num_docks_available,omitempty"` // conditionally required (v2.0)
		NumDocksDisabled      *int64                          `json:"num_docks_disabled,omitempty"`
		IsInstalled           *Boolean                        `json:"is_installed"`
		IsRenting             *Boolean                        `json:"is_renting"`
		IsReturning           *Boolean                        `json:"is_returning"`
		LastReported          *Timestamp                      `json:"last_reported"`
		VehicleTypesAvailable []*FeedStationStatusVehicleType `json:"vehicle_types_available,omitempty"` // conditionally required (v2.1-RC)
		VehicleDocksAvailable []*FeedStationStatusVehicleDock `json:"vehicle_docks_available,omitempty"` // conditionally required (v2.1-RC)
	}
	// FeedStationStatusVehicleType ...
	FeedStationStatusVehicleType struct {
		VehicleTypeID *ID    `json:"vehicle_type_id"` // (v2.1-RC)
		Count         *int64 `json:"count"`           // (v2.1-RC)
	}
	// FeedStationStatusVehicleDock ...
	FeedStationStatusVehicleDock struct {
		VehicleTypeIDs []*ID  `json:"vehicle_type_ids"` // (v2.1-RC)
		Count          *int64 `json:"count"`            // (v2.1-RC)
	}
)

// Name ...
func (f *FeedStationStatus) Name() string {
	return FeedNameStationStatus
}
