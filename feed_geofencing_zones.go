package gbfs

type (
	// FeedGeofencingZones ...
	FeedGeofencingZones struct {
		FeedCommon
		Data *FeedGeofencingZonesData `json:"data"`
	}
	// FeedGeofencingZonesData ...
	FeedGeofencingZonesData struct {
		GeofencingZones *FeedGeofencingZonesGeoJSONFeatureCollection `json:"geofencing_zones"`
	}
	// FeedGeofencingZonesGeoJSONFeatureCollection ...
	FeedGeofencingZonesGeoJSONFeatureCollection struct {
		GeoJSONFeatureCollection
		Features []*FeedGeofencingZonesGeoJSONFeature `json:"features"`
	}
	// FeedGeofencingZonesGeoJSONFeature ...
	FeedGeofencingZonesGeoJSONFeature struct {
		GeoJSONFeature
		Properties *FeedGeofencingZonesGeoJSONFeatureProperties `json:"properties"`
	}
	// FeedGeofencingZonesGeoJSONFeatureProperties ...
	FeedGeofencingZonesGeoJSONFeatureProperties struct {
		Name  string                                             `json:"name,omitempty"`
		Start Timestamp                                          `json:"start,omitempty"`
		End   Timestamp                                          `json:"end,omitempty"`
		Rules []*FeedGeofencingZonesGeoJSONFeaturePropertiesRule `json:"rules,omitempty"`
	}
	// FeedGeofencingZonesGeoJSONFeaturePropertiesRule ...
	FeedGeofencingZonesGeoJSONFeaturePropertiesRule struct {
		VehicleTypeIDs     []string `json:"vehicle_type_ids,omitempty"`
		RideAllowed        Boolean  `json:"ride_allowed"`
		RideThroughAllowed Boolean  `json:"ride_through_allowed"`
		MaximumSpeedKph    int64    `json:"maximum_speed_kph,omitempty"`
	}
)

// Name ...
func (f *FeedGeofencingZones) Name() string {
	return FeedNameGeofencingZones
}

// NewFeedGeofencingZonesGeoJSONFeature ...
func NewFeedGeofencingZonesGeoJSONFeature(geometry *GeoJSONGeometry, properties *FeedGeofencingZonesGeoJSONFeatureProperties) *FeedGeofencingZonesGeoJSONFeature {
	return &FeedGeofencingZonesGeoJSONFeature{
		GeoJSONFeature: *NewGeoJSONFeature(geometry, nil),
		Properties:     properties,
	}
}

// NewFeedGeofencingZonesGeoJSONFeatureCollection ...
func NewFeedGeofencingZonesGeoJSONFeatureCollection(features []*FeedGeofencingZonesGeoJSONFeature) *FeedGeofencingZonesGeoJSONFeatureCollection {
	return &FeedGeofencingZonesGeoJSONFeatureCollection{
		GeoJSONFeatureCollection: *NewGeoJSONFeatureCollection(nil),
		Features:                 features,
	}
}
