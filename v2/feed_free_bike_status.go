package gbfs

type (
	// FeedFreeBikeStatus ...
	FeedFreeBikeStatus struct {
		FeedCommon
		Data *FeedFreeBikeStatusData `json:"data"`
	}
	// FeedFreeBikeStatusData ...
	FeedFreeBikeStatusData struct {
		Bikes []*FeedFreeBikeStatusBike `json:"bikes"`
	}
	// FeedFreeBikeStatusBike ...
	FeedFreeBikeStatusBike struct {
		BikeID             *ID         `json:"bike_id"`
		SystemID           *ID         `json:"system_id,omitempty"` // (v3.0-RC)
		Lat                *Coordinate `json:"lat"`
		Lon                *Coordinate `json:"lon"`
		IsReserved         *Boolean    `json:"is_reserved"`
		IsDisabled         *Boolean    `json:"is_disabled"`
		RentalURIs         *RentalURIs `json:"rental_uris,omitempty"`          // (v1.1)
		VehicleTypeID      *ID         `json:"vehicle_type_id,omitempty"`      // (v2.1-RC)
		LastReported       *Timestamp  `json:"last_reported,omitempty"`        // (v2.1-RC)
		CurrentRangeMeters *float64    `json:"current_range_meters,omitempty"` // (v2.1-RC)
	}
)

// Name ...
func (f *FeedFreeBikeStatus) Name() string {
	return FeedNameFreeBikeStatus
}
