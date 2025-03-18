package gbfs

type (
	FeedStationStatus struct {
		FeedCommon
		Data *FeedStationStatusData `json:"data"`
	}
	FeedStationStatusData struct {
		Stations []*FeedStationStatusStation `json:"stations"`
	}
	FeedStationStatusStation struct {
		StationID             *ID                     `json:"station_id"`
		NumVehiclesAvailable  *int64                  `json:"num_vehicles_available"`
		VehicleTypesAvailable []*VehicleTypeCapacity  `json:"vehicle_types_available,omitempty"`
		NumVehiclesDisabled   *int64                  `json:"num_vehicles_disabled,omitempty"`
		NumDocksAvailable     *int64                  `json:"num_docks_available,omitempty"`
		VehicleDocksAvailable []*VehicleTypesCapacity `json:"vehicle_docks_available,omitempty"`
		NumDocksDisabled      *int64                  `json:"num_docks_disabled,omitempty"`
		IsInstalled           *Boolean                `json:"is_installed"`
		IsRenting             *Boolean                `json:"is_renting"`
		IsReturning           *Boolean                `json:"is_returning"`
		LastReported          *Timestamp              `json:"last_reported"`
	}
)

func (f *FeedStationStatus) Name() string {
	return FeedNameStationStatus
}
