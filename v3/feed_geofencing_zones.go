package gbfs

type (
	FeedGeofencingZones struct {
		FeedCommon
		Data *FeedGeofencingZonesData `json:"data"`
	}
	FeedGeofencingZonesData struct {
		GeofencingZones *FeedGeofencingZonesGeoJSONFeatureCollection `json:"geofencing_zones"`
		GlobalRules     []*FeedGeofencingZonesRule                   `json:"global_rules"`
	}
	FeedGeofencingZonesGeoJSONFeatureCollection struct {
		GeoJSONFeatureCollection
		Features []*FeedGeofencingZonesGeoJSONFeature `json:"features"`
	}
	FeedGeofencingZonesGeoJSONFeature struct {
		GeoJSONFeature
		Properties *FeedGeofencingZonesGeoJSONFeatureProperties `json:"properties"`
	}
	FeedGeofencingZonesGeoJSONFeatureProperties struct {
		Name  []*LocalizedString         `json:"name,omitempty"`
		Start *Timestamp                 `json:"start,omitempty"`
		End   *Timestamp                 `json:"end,omitempty"`
		Rules []*FeedGeofencingZonesRule `json:"rules,omitempty"`
	}
	FeedGeofencingZonesRule struct {
		VehicleTypeIDs     []*ID    `json:"vehicle_type_ids,omitempty"`
		RideStartAllowed   *Boolean `json:"ride_start_allowed"`
		RideEndAllowed     *Boolean `json:"ride_end_allowed"`
		RideThroughAllowed *Boolean `json:"ride_through_allowed"`
		MaximumSpeedKph    *int64   `json:"maximum_speed_kph,omitempty"`
		StationParking     *Boolean `json:"station_parking,omitempty"`
	}
)

func (f *FeedGeofencingZones) Name() string {
	return FeedNameGeofencingZones
}

func NewFeedGeofencingZonesGeoJSONFeature(geometry *GeoJSONGeometry, properties *FeedGeofencingZonesGeoJSONFeatureProperties) *FeedGeofencingZonesGeoJSONFeature {
	return &FeedGeofencingZonesGeoJSONFeature{
		GeoJSONFeature: *NewGeoJSONFeature(geometry, nil),
		Properties:     properties,
	}
}

func NewFeedGeofencingZonesGeoJSONFeatureCollection(features []*FeedGeofencingZonesGeoJSONFeature) *FeedGeofencingZonesGeoJSONFeatureCollection {
	return &FeedGeofencingZonesGeoJSONFeatureCollection{
		GeoJSONFeatureCollection: *NewGeoJSONFeatureCollection(nil),
		Features:                 features,
	}
}
