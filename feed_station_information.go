package gbfs

type (
	// FeedStationInformation ...
	FeedStationInformation struct {
		FeedCommon
		Data *FeedStationInformationData `json:"data"`
	}
	// FeedStationInformationData ...
	FeedStationInformationData struct {
		Stations []*FeedStationInformationStation `json:"stations"`
	}
	// FeedStationInformationStation ...
	FeedStationInformationStation struct {
		StationID           string           `json:"station_id"`
		Name                string           `json:"name"`
		ShortName           string           `json:"short_name,omitempty"`
		Lat                 float64          `json:"lat"`
		Lon                 float64          `json:"lon"`
		Address             string           `json:"address,omitempty"`
		CrossStreet         string           `json:"cross_street,omitempty"`
		RegionID            string           `json:"region_id,omitempty"`
		PostCode            string           `json:"post_code,omitempty"`
		RentalMethods       []string         `json:"rental_methods,omitempty"`
		IsVirtualStation    Boolean          `json:"is_virtual_station,omitempty"` // (v2.1-RC)
		StationArea         *GeoJSONGeometry `json:"station_area,omitempty"`       // (v2.1-RC)
		Capacity            int64            `json:"capacity,omitempty"`
		VehicleCapacity     map[string]int64 `json:"vehicle_capacity,omitempty"`      // (v2.1-RC)
		IsValetStation      Boolean          `json:"is_valet_station,omitempty"`      // (v2.1-RC)
		RentalURIs          *RentalURIs      `json:"rental_uris,omitempty"`           // (v1.1)
		VehicleTypeCapacity map[string]int64 `json:"vehicle_type_capacity,omitempty"` // (v2.1-RC)
	}
)

// Name ...
func (f *FeedStationInformation) Name() string {
	return FeedNameStationInformation
}
