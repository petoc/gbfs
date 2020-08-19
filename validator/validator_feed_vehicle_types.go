package validator

import "github.com/petoc/gbfs"

// ValidateFeedVehicleTypes ...
func ValidateFeedVehicleTypes(f *gbfs.FeedVehicleTypes, version string) *Result {
	r := &Result{
		Feed: f,
	}
	if nilOrEmpty(f.Data) {
		r.ErrorW("data", ErrRequired)
		return r
	}
	if nilOrZero(f.Data.VehicleTypes) {
		r.ErrorW("data.vehicle_types", ErrRequired)
		return r
	}
	for i, s := range f.Data.VehicleTypes {
		sliceIndexName := sliceIndexN("data.vehicle_types", i)
		if nilOrEmpty(s) {
			r.ErrorW(sliceIndexName, ErrInvalidValue)
			continue
		}
		if !ValidateFormFactor(s.FormFactor) {
			r.ErrorW(sliceIndexName+".form_factor", ErrInvalidValue)
		}
		if !ValidatePropulsionType(s.PropulsionType) {
			r.ErrorW(sliceIndexName+".propulsion_type", ErrInvalidValue)
		} else {
			if *s.PropulsionType != gbfs.PropulsionTypeHuman {
				if s.MaxRangeMeters == nil {
					r.ErrorWS(sliceIndexName+".max_range_meters", ErrRequired, " for non-human propulsion type")
				}
			}
		}
	}
	return r
}
