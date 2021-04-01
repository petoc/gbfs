package validator

import (
	"testing"
	"time"

	"github.com/petoc/gbfs"
)

func TestValidateFeedStationStatus(t *testing.T) {
	f := &gbfs.FeedStationStatus{}
	f.SetLastUpdated(gbfs.Timestamp(time.Now().Unix()))
	f.SetTTL(5)
	r := ValidateFeedStationStatus(f, "")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data = &gbfs.FeedStationStatusData{}
	r = ValidateFeedStationStatus(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data.Stations = []*gbfs.FeedStationStatusStation{}
	r = ValidateFeedStationStatus(f, "")
	if hasError(r.Errors, ErrRequired) || len(r.Errors) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	station := &gbfs.FeedStationStatusStation{}
	f.Data.Stations = []*gbfs.FeedStationStatusStation{station}
	r = ValidateFeedStationStatus(f, "")
	if !hasError(r.Errors, ErrInvalidValue) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	station.StationID = gbfs.NewID("")
	r = ValidateFeedStationStatus(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	if errorCount(r.Errors, ErrRequired) != 5 {
		t.Errorf("expected 5 errors of [%s], got %v", ErrRequired, r.Errors)
		return
	}
	station.StationID = gbfs.NewID("123")
	station.NumBikesAvailable = gbfs.NewInt64(-1)
	station.NumBikesDisabled = gbfs.NewInt64(-1)
	station.NumDocksAvailable = gbfs.NewInt64(-1)
	station.NumDocksDisabled = gbfs.NewInt64(-1)
	station.IsInstalled = gbfs.NewBoolean(false)
	station.IsRenting = gbfs.NewBoolean(false)
	station.IsReturning = gbfs.NewBoolean(false)
	station.LastReported = gbfs.NewTimestamp(0)
	r = ValidateFeedStationStatus(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 5 {
		t.Errorf("expected 5 errors of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	station.NumBikesAvailable = gbfs.NewInt64(0)
	station.NumBikesDisabled = gbfs.NewInt64(0)
	station.NumDocksAvailable = gbfs.NewInt64(0)
	station.NumDocksDisabled = gbfs.NewInt64(0)
	station.LastReported = gbfs.NewTimestamp(123456789)
	r = ValidateFeedStationStatus(f, "")
	if errorCount(r.Errors, ErrInvalidValue) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	station.VehicleTypesAvailable = []*gbfs.FeedStationStatusVehicleType{}
	station.VehicleDocksAvailable = []*gbfs.FeedStationStatusVehicleDock{}
	r = ValidateFeedStationStatus(f, "")
	if errorCount(r.Infos, ErrAvailableFromVersion) != 2 {
		t.Errorf("expected 2 infos of [%s], got %v", ErrAvailableFromVersion, r.Infos)
		return
	}
	r = ValidateFeedStationStatus(f, "2.1")
	if errorCount(r.Errors, ErrInvalidValue) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	vehicleType := &gbfs.FeedStationStatusVehicleType{}
	station.VehicleTypesAvailable = []*gbfs.FeedStationStatusVehicleType{vehicleType}
	r = ValidateFeedStationStatus(f, "2.1")
	if errorCount(r.Errors, ErrInvalidValue) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	vehicleType.VehicleTypeID = nil
	r = ValidateFeedStationStatus(f, "2.1")
	if errorCount(r.Errors, ErrInvalidValue) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	vehicleType.Count = gbfs.NewInt64(0)
	r = ValidateFeedStationStatus(f, "2.1")
	if errorCount(r.Errors, ErrRequired) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrRequired, r.Errors)
		return
	}
	vehicleType.VehicleTypeID = gbfs.NewID("")
	r = ValidateFeedStationStatus(f, "2.1")
	if errorCount(r.Errors, ErrRequired) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrRequired, r.Errors)
		return
	}
	vehicleType.VehicleTypeID = gbfs.NewID("vehicleType1")
	r = ValidateFeedStationStatus(f, "2.1")
	if errorCount(r.Errors, ErrRequired) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	vehicleDock := &gbfs.FeedStationStatusVehicleDock{}
	station.VehicleDocksAvailable = []*gbfs.FeedStationStatusVehicleDock{vehicleDock}
	r = ValidateFeedStationStatus(f, "2.1")
	if errorCount(r.Errors, ErrInvalidValue) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	vehicleDock.VehicleTypeIDs = []*gbfs.ID{}
	r = ValidateFeedStationStatus(f, "2.1")
	if errorCount(r.Errors, ErrInvalidValue) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	if errorCount(r.Errors, ErrRequired) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrRequired, r.Errors)
		return
	}
	vehicleDock.VehicleTypeIDs = []*gbfs.ID{gbfs.NewID("")}
	vehicleDock.Count = gbfs.NewInt64(0)
	r = ValidateFeedStationStatus(f, "2.1")
	if errorCount(r.Errors, ErrInvalidValue) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	vehicleDock.VehicleTypeIDs = []*gbfs.ID{gbfs.NewID("vehicleType1")}
	r = ValidateFeedStationStatus(f, "2.1")
	if errorCount(r.Errors, ErrInvalidValue) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
}
