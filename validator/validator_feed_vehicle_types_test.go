package validator

import (
	"testing"
	"time"

	"github.com/petoc/gbfs"
)

func TestValidateFeedVehicleTypes(t *testing.T) {
	f := &gbfs.FeedVehicleTypes{}
	f.SetLastUpdated(gbfs.Timestamp(time.Now().Unix()))
	f.SetTTL(5)
	r := ValidateFeedVehicleTypes(f, "")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data = &gbfs.FeedVehicleTypesData{}
	r = ValidateFeedVehicleTypes(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data.VehicleTypes = []*gbfs.FeedVehicleTypesVehicleType{}
	r = ValidateFeedVehicleTypes(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	vehicleType := &gbfs.FeedVehicleTypesVehicleType{}
	f.Data.VehicleTypes = append(f.Data.VehicleTypes, vehicleType)
	r = ValidateFeedVehicleTypes(f, "")
	if !hasError(r.Errors, ErrInvalidValue) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	vehicleType.FormFactor = gbfs.NewString("invalid")
	vehicleType.PropulsionType = gbfs.NewString("invalid")
	r = ValidateFeedVehicleTypes(f, "")
	if !hasError(r.Errors, ErrInvalidValue) || len(r.Errors) != 2 {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	vehicleType.FormFactor = gbfs.NewString(gbfs.FormFactorBicycle)
	vehicleType.PropulsionType = gbfs.NewString(gbfs.PropulsionTypeHuman)
	r = ValidateFeedVehicleTypes(f, "")
	if hasError(r.Errors, ErrInvalidValue) || len(r.Errors) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	vehicleType.PropulsionType = gbfs.NewString(gbfs.PropulsionTypeElectricAssist)
	r = ValidateFeedVehicleTypes(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	vehicleType.MaxRangeMeters = gbfs.NewFloat64(20000)
	r = ValidateFeedVehicleTypes(f, "")
	if hasError(r.Errors, ErrRequired) || len(r.Errors) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
}
