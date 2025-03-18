package gbfs

type (
	FeedStationInformation struct {
		FeedCommon
		Data *FeedStationInformationData `json:"data"`
	}
	FeedStationInformationData struct {
		Stations []*FeedStationInformationStation `json:"stations"`
	}
	FeedStationInformationStation struct {
		StationID            *ID                     `json:"station_id"`
		Name                 []*LocalizedString      `json:"name"`
		ShortName            []*LocalizedString      `json:"short_name,omitempty"`
		Lat                  *Coordinate             `json:"lat"`
		Lon                  *Coordinate             `json:"lon"`
		Address              *string                 `json:"address,omitempty"`
		CrossStreet          *string                 `json:"cross_street,omitempty"`
		RegionID             *ID                     `json:"region_id,omitempty"`
		PostCode             *string                 `json:"post_code,omitempty"`
		StationOpeningHours  *string                 `json:"station_opening_hours ,omitempty"`
		RentalMethods        []string                `json:"rental_methods,omitempty"`
		IsVirtualStation     *Boolean                `json:"is_virtual_station,omitempty"`
		StationArea          *GeoJSONGeometry        `json:"station_area,omitempty"`
		ParkingType          *string                 `json:"parking_type,omitempty"`
		ParkingHoop          *Boolean                `json:"parking_hoop,omitempty"`
		ContactPhone         *string                 `json:"contact_phone,omitempty"`
		Capacity             *int64                  `json:"capacity,omitempty"`
		VehicleTypesCapacity []*VehicleTypesCapacity `json:"vehicle_types_capacity,omitempty"`
		VehicleDocksCapacity []*VehicleTypesCapacity `json:"vehicle_docks_capacity,omitempty"`
		IsValetStation       *Boolean                `json:"is_valet_station,omitempty"`
		IsChargingStation    *Boolean                `json:"is_charging_station,omitempty"`
		RentalURIs           *RentalURIs             `json:"rental_uris,omitempty"`
	}
)

func (f *FeedStationInformation) Name() string {
	return FeedNameStationInformation
}
