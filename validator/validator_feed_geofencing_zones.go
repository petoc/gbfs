package validator

import "github.com/petoc/gbfs"

// ValidateFeedGeofencingZones ...
func ValidateFeedGeofencingZones(f *gbfs.FeedGeofencingZones, version string) *Result {
	r := &Result{
		Feed: f,
	}
	if nilOrEmpty(f.Data) {
		r.ErrorW("data", ErrRequired)
		return r
	}
	if f.Data.GeofencingZones == nil {
		r.ErrorW("data.geofencing_zones", ErrRequired)
		return r
	}
	if f.Data.GeofencingZones.Type != "FeatureCollection" {
		r.ErrorWSP("data.geofencing_zones.type", ErrInvalidType, "FeatureCollection")
		return r
	}
	if f.Data.GeofencingZones.Features == nil {
		r.ErrorW("data.geofencing_zones.features", ErrRequired)
		return r
	}
	if len(f.Data.GeofencingZones.Features) == 0 {
		return r
	}
	for i, s := range f.Data.GeofencingZones.Features {
		sliceIndexName := sliceIndexN("data.geofencing_zones.features", i)
		if nilOrEmpty(s) {
			r.ErrorW(sliceIndexName, ErrInvalidValue)
			continue
		}
		if s.Type != "Feature" {
			r.ErrorWSP(sliceIndexName+".type", ErrInvalidType, "Feature")
			return r
		}
		if s.Geometry == nil {
			r.ErrorW(sliceIndexName+".geometry", ErrRequired)
		} else {
			if s.Geometry.Type != "MultiPolygon" {
				r.ErrorWSP(sliceIndexName+".geometry.type", ErrInvalidType, "MultiPolygon")
			} else if s.Geometry.Coordinates == nil {
				r.ErrorW(sliceIndexName+".geometry.coordinates", ErrRequired)
			} else {
				coords, ok := s.Geometry.Coordinates.([]interface{})
				if !ok || len(coords) == 0 {
					r.ErrorW(sliceIndexName+".geometry.coordinates", ErrInvalidValue)
				}
			}
		}
		if s.Properties == nil {
			r.ErrorW(sliceIndexName+".properties", ErrRequired)
		} else {
			if s.Properties.Name != nil && *s.Properties.Name == "" {
				r.ErrorW(sliceIndexName+".properties.name", ErrInvalidValue)
			}
			if s.Properties.Start != nil && !validateTimestamp(*s.Properties.Start) {
				r.ErrorW(sliceIndexName+".properties.start", ErrInvalidValue)
			}
			if s.Properties.End != nil && !validateTimestamp(*s.Properties.End) {
				r.ErrorW(sliceIndexName+".properties.end", ErrInvalidValue)
			}
			if s.Properties.Rules != nil {
				if len(s.Properties.Rules) > 0 {
					for j, rule := range s.Properties.Rules {
						ruleIndexName := sliceIndexN(sliceIndexName+".rules", j)
						if nilOrEmpty(rule) {
							r.ErrorW(ruleIndexName, ErrInvalidValue)
							continue
						}
						if rule.VehicleTypeIDs != nil {
							if len(rule.VehicleTypeIDs) == 0 {
								r.ErrorW(ruleIndexName+".vehicle_type_ids", ErrInvalidValue)
							} else {
								for k, vehicleTypeID := range rule.VehicleTypeIDs {
									if vehicleTypeID == nil || *vehicleTypeID == "" {
										r.ErrorW(sliceIndexN(ruleIndexName+".vehicle_type_ids", k), ErrInvalidValue)
									}
								}
							}
						}
						if rule.RideAllowed == nil {
							r.ErrorW(ruleIndexName+".ride_allowed", ErrRequired)
						}
						if rule.RideThroughAllowed == nil {
							r.ErrorW(ruleIndexName+".ride_through_allowed", ErrRequired)
						}
						if rule.MaximumSpeedKph != nil && *rule.MaximumSpeedKph < 0 {
							r.ErrorW(ruleIndexName+".maximum_speed_kph", ErrInvalidValue)
						}
					}
				}
			}
		}
	}
	return r
}
