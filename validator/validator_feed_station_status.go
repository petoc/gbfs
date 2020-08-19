package validator

import "github.com/petoc/gbfs"

// ValidateFeedStationStatus ...
func ValidateFeedStationStatus(f *gbfs.FeedStationStatus, version string) *Result {
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
		if s.NumBikesAvailable == nil {
			r.ErrorW(sliceIndexName+".num_bikes_available", ErrRequired)
		} else if *s.NumBikesAvailable < 0 {
			r.ErrorW(sliceIndexName+".num_bikes_available", ErrInvalidValue)
		}
		if s.NumBikesDisabled != nil && *s.NumBikesDisabled < 0 {
			r.ErrorW(sliceIndexName+".num_bikes_disabled", ErrInvalidValue)
		}
		if s.NumDocksAvailable != nil && *s.NumDocksAvailable < 0 {
			r.ErrorW(sliceIndexName+".num_docks_available", ErrInvalidValue)
		}
		if s.NumDocksDisabled != nil && *s.NumDocksDisabled < 0 {
			r.ErrorW(sliceIndexName+".num_docks_disabled", ErrInvalidValue)
		}
		if s.IsInstalled == nil {
			r.ErrorW(sliceIndexName+".is_installed", ErrRequired)
		}
		if s.IsRenting == nil {
			r.ErrorW(sliceIndexName+".is_renting", ErrRequired)
		}
		if s.IsReturning == nil {
			r.ErrorW(sliceIndexName+".is_returning", ErrRequired)
		}
		if s.LastReported == nil {
			r.ErrorW(sliceIndexName+".last_reported", ErrRequired)
		} else if *s.LastReported == 0 {
			r.ErrorWSP(sliceIndexName+".last_reported", ErrInvalidValue, "POSIX time")
		}
		if verLT(version, gbfs.V21) {
			if s.VehicleDocksAvailable != nil {
				r.InfoWV(sliceIndexName+".vehicle_docks_available", ErrAvailableFromVersion, gbfs.V21)
			}
			if s.Vehicles != nil {
				r.InfoWV(sliceIndexName+".vehicles", ErrAvailableFromVersion, gbfs.V21)
			}
		} else {
			if s.VehicleDocksAvailable != nil && len(s.VehicleDocksAvailable) > 0 {
				for j, dock := range s.VehicleDocksAvailable {
					dockIndexName := sliceIndexN(sliceIndexName+".vehicle_docks_available", j)
					if nilOrEmpty(dock) {
						r.ErrorW(dockIndexName, ErrInvalidValue)
						continue
					}
					if dock.VehicleTypeIDs == nil {
						r.ErrorW(dockIndexName+".vehicle_type_ids", ErrRequired)
					} else if len(dock.VehicleTypeIDs) == 0 {
						r.ErrorW(dockIndexName+".vehicle_type_ids", ErrInvalidValue)
					} else {
						for k, vehicleTypeID := range dock.VehicleTypeIDs {
							if vehicleTypeID == nil || *vehicleTypeID == "" {
								r.ErrorW(sliceIndexN(dockIndexName+".vehicle_type_ids", k), ErrInvalidValue)
							}
						}
					}
					if dock.Count == nil {
						r.ErrorW(dockIndexName+".count", ErrRequired)
					} else if *dock.Count < 0 {
						r.ErrorW(dockIndexName+".count", ErrInvalidValue)
					}
				}
			}
			if s.Vehicles != nil && len(s.Vehicles) > 0 {
				for j, vehicle := range s.Vehicles {
					vehicleIndexName := sliceIndexN(sliceIndexName+".vehicles", j)
					if nilOrEmpty(vehicle) {
						r.ErrorW(vehicleIndexName, ErrInvalidValue)
						continue
					}
					if vehicle.BikeID == nil {
						r.ErrorW(vehicleIndexName+".bike_id", ErrRequired)
					} else if *vehicle.BikeID == "" {
						r.ErrorW(vehicleIndexName+".bike_id", ErrInvalidValue)
					}
					if vehicle.IsReserved == nil {
						r.ErrorW(vehicleIndexName+".is_reserved", ErrRequired)
					}
					if vehicle.IsDisabled == nil {
						r.ErrorW(vehicleIndexName+".is_disabled", ErrRequired)
					}
					if vehicle.VehicleTypeID == nil {
						r.ErrorW(vehicleIndexName+".vehicle_type_id", ErrRequired)
					} else if *vehicle.VehicleTypeID == "" {
						r.ErrorW(vehicleIndexName+".vehicle_type_id", ErrInvalidValue)
					}
					if vehicle.CurrentRangeMeters != nil && *vehicle.CurrentRangeMeters < 0 {
						r.ErrorW(vehicleIndexName+".current_range_meters", ErrInvalidValue)
					}
				}
			}
		}
	}
	return r
}
