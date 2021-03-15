package validator

import (
	"github.com/petoc/gbfs"
)

// ValidateFeedStationInformation ...
func ValidateFeedStationInformation(f *gbfs.FeedStationInformation, version string) *Result {
	r := &Result{
		Feed: f,
	}
	if nilOrEmpty(f.Data) {
		r.ErrorW("data", ErrRequired)
		return r
	}
	if f.Data.Stations == nil {
		r.ErrorW("data.stations", ErrRequired)
		return r
	}
	if len(f.Data.Stations) == 0 {
		return r
	}
	for i, s := range f.Data.Stations {
		sliceIndexName := sliceIndexN("data.stations", i)
		if nilOrEmpty(s) {
			r.ErrorW(sliceIndexName, ErrInvalidValue)
			continue
		}
		if s.StationID == nil {
			r.ErrorW(sliceIndexName+".station_id", ErrRequired)
		} else if *s.StationID == "" {
			r.ErrorW(sliceIndexName+".station_id", ErrInvalidValue)
		}
		if s.Name == nil || *s.Name == "" {
			r.ErrorW(sliceIndexName+".name", ErrRequired)
		}
		if s.ShortName != nil && *s.ShortName == "" {
			r.ErrorW(sliceIndexName+".short_name", ErrInvalidValue)
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
		if s.Address != nil && *s.Address == "" {
			r.ErrorW(sliceIndexName+".address", ErrInvalidValue)
		}
		if s.CrossStreet != nil && *s.CrossStreet == "" {
			r.ErrorW(sliceIndexName+".cross_street", ErrInvalidValue)
		}
		if s.RegionID != nil && *s.RegionID == "" {
			r.ErrorW(sliceIndexName+".region_id", ErrInvalidValue)
		}
		if s.PostCode != nil && *s.PostCode == "" {
			r.ErrorW(sliceIndexName+".post_code", ErrInvalidValue)
		}
		if s.RentalMethods != nil {
			if len(s.RentalMethods) > 0 {
				for j, m := range s.RentalMethods {
					if !ValidateRentalMethod(m) {
						r.ErrorW(sliceIndexName+sliceIndexN(".rental_methods", j), ErrInvalidValue)
					}
				}
			}
		}
		if s.Capacity != nil && *s.Capacity < 0 {
			r.ErrorW(sliceIndexName+".capacity", ErrInvalidValue)
		}
		if verLT(version, gbfs.V21) {
			if s.IsVirtualStation != nil {
				r.InfoWV(sliceIndexName+".is_virtual_station", ErrAvailableFromVersion, gbfs.V21)
			}
			if s.StationArea != nil {
				r.InfoWV(sliceIndexName+".station_area", ErrAvailableFromVersion, gbfs.V21)
			}
			if s.VehicleCapacity != nil {
				r.InfoWV(sliceIndexName+".vehicle_capacity", ErrAvailableFromVersion, gbfs.V21)
			}
			if s.IsValetStation != nil {
				r.InfoWV(sliceIndexName+".is_valet_station", ErrAvailableFromVersion, gbfs.V21)
			}
			if s.VehicleTypeCapacity != nil {
				r.InfoWV(sliceIndexName+".vehicle_type_capacity", ErrAvailableFromVersion, gbfs.V21)
			}
		} else {
			if s.StationArea != nil {
				if s.StationArea.Type != "MultiPolygon" {
					r.ErrorWSP(sliceIndexName+".station_area", ErrInvalidType, "MultiPolygon")
				} else if s.StationArea.Coordinates == nil {
					r.ErrorW(sliceIndexName+".station_area.coordinates", ErrRequired)
				} else {
					coords, ok := s.StationArea.Coordinates.([][][][]float64)
					if !ok || len(coords) == 0 {
						r.ErrorW(sliceIndexName+".station_area.coordinates", ErrInvalidValue)
					}
				}
			}
			if s.VehicleCapacity != nil {
				if len(s.VehicleCapacity) == 0 {
					r.ErrorW(sliceIndexName+".vehicle_capacity", ErrInvalidValue)
				} else {
					for vehicleTypeID, c := range s.VehicleCapacity {
						if vehicleTypeID == nil || *vehicleTypeID == "" {
							r.ErrorW(sliceIndexName+".vehicle_capacity: vehicleTypeID key", ErrInvalidValue)
							continue
						}
						if c < 0 {
							r.ErrorW(sliceIndexName+".vehicle_capacity["+string(*vehicleTypeID)+"]", ErrInvalidValue)
						}
					}
				}
			}
			if s.VehicleTypeCapacity != nil {
				if len(s.VehicleTypeCapacity) == 0 {
					r.ErrorW(sliceIndexName+".vehicle_type_capacity", ErrInvalidValue)
				} else {
					for vehicleTypeID, c := range s.VehicleTypeCapacity {
						if vehicleTypeID == nil || *vehicleTypeID == "" {
							r.ErrorW(sliceIndexName+".vehicle_type_capacity: vehicleTypeID key", ErrInvalidValue)
							continue
						}
						if c < 0 {
							r.ErrorW(sliceIndexName+".vehicle_type_capacity["+string(*vehicleTypeID)+"]", ErrInvalidValue)
						}
					}
				}
			}
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
	}
	return r
}
