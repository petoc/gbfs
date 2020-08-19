package validator

import "github.com/petoc/gbfs"

// ValidateFeedFreeBikeStatus ...
func ValidateFeedFreeBikeStatus(f *gbfs.FeedFreeBikeStatus, version string) *Result {
	r := &Result{
		Feed: f,
	}
	if nilOrEmpty(f.Data) {
		r.ErrorW("data", ErrRequired)
		return r
	}
	if f.Data.Bikes == nil {
		r.ErrorW("data.bikes", ErrRequired)
		return r
	}
	if len(f.Data.Bikes) == 0 {
		return r
	}
	for i, s := range f.Data.Bikes {
		sliceIndexName := sliceIndexN("data.bikes", i)
		if nilOrEmpty(s) {
			r.ErrorW(sliceIndexName, ErrInvalidValue)
			continue
		}
		if s.BikeID == nil {
			r.ErrorW(sliceIndexName+".bike_id", ErrRequired)
		} else if *s.BikeID == "" {
			r.ErrorW(sliceIndexName+".bike_id", ErrInvalidValue)
		}
		if verLT(version, gbfs.V30) {
			if s.SystemID != nil {
				r.InfoWV(sliceIndexName+".system_id", ErrAvailableFromVersion, gbfs.V30)
			}
		} else {
			if s.SystemID == nil {
				r.ErrorW(sliceIndexName+".system_id", ErrRequired)
			} else if *s.SystemID == "" {
				r.ErrorW(sliceIndexName+".system_id", ErrInvalidValue)
			}
		}
		if s.Lat != nil && s.Lon != nil && s.Lat.OldType == "" && s.Lon.OldType == "" && s.Lat.Float64 == 0 && s.Lon.Float64 == 0 {
			r.WarningW(sliceIndexName+".lat, "+sliceIndexName+".lon", ErrZeroCoordinates)
		} else {
			if s.Lat == nil {
				r.ErrorW(sliceIndexName+".lat", ErrRequired)
			} else if s.Lat.OldType != "" {
				r.ErrorWSP(sliceIndexName+".lat", ErrInvalidType, "decimal degrees")
			} else if !validateLatitude(&s.Lat.Float64) {
				r.ErrorW(sliceIndexName+".lat", ErrOutOfRange)
			}
			if s.Lon == nil {
				r.ErrorW(sliceIndexName+".lon", ErrRequired)
			} else if s.Lon.OldType != "" {
				r.ErrorWSP(sliceIndexName+".lon", ErrInvalidType, "decimal degrees")
			} else if !validateLongitude(&s.Lon.Float64) {
				r.ErrorW(sliceIndexName+".lon", ErrOutOfRange)
			}
		}
		if s.IsReserved == nil {
			r.ErrorW(sliceIndexName+".is_reserved", ErrRequired)
		}
		if s.IsDisabled == nil {
			r.ErrorW(sliceIndexName+".is_disabled", ErrRequired)
		}
		if verLT(version, gbfs.V11) {
			if s.RentalURIs != nil {
				r.InfoWV(sliceIndexName+".rental_uris", ErrAvailableFromVersion, gbfs.V11)
			}
		} else {
			if !nilOrEmpty(s.RentalURIs) {
				if s.RentalURIs.Android != nil && !validateURI(s.RentalURIs.Android) {
					r.ErrorW(sliceIndexName+".android", ErrInvalidValue)
				}
				if s.RentalURIs.IOS != nil && !validateURI(s.RentalURIs.IOS) {
					r.ErrorW(sliceIndexName+".ios", ErrInvalidValue)
				}
				if s.RentalURIs.Web != nil && !validateURI(s.RentalURIs.Web) {
					r.ErrorW(sliceIndexName+".web", ErrInvalidValue)
				}
			}
		}
		if verLT(version, gbfs.V21) {
			if s.VehicleTypeID != nil {
				r.InfoWV(sliceIndexName+".vehicle_type_id", ErrAvailableFromVersion, gbfs.V21)
			}
			if s.LastReported != nil {
				r.InfoWV(sliceIndexName+".last_reported", ErrAvailableFromVersion, gbfs.V21)
			}
			if s.CurrentRangeMeters != nil {
				r.InfoWV(sliceIndexName+".current_range_meters", ErrAvailableFromVersion, gbfs.V21)
			}
		} else {
			if s.VehicleTypeID != nil && *s.VehicleTypeID == "" {
				r.ErrorW(sliceIndexName+".vehicle_type_id", ErrInvalidValue)
			}
			if s.LastReported != nil && *s.LastReported == 0 {
				r.ErrorWSP(sliceIndexName+".last_reported", ErrInvalidValue, "POSIX time")
			}
			if s.CurrentRangeMeters != nil && *s.LastReported < 0 {
				r.ErrorW(sliceIndexName+".current_range_meters", ErrInvalidValue)
			}
		}
	}
	return r
}
