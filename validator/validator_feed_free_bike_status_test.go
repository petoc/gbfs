package validator

import (
	"testing"
	"time"

	"github.com/petoc/gbfs"
)

func TestValidateFeedFreeBikeStatus(t *testing.T) {
	f := &gbfs.FeedFreeBikeStatus{}
	f.SetLastUpdated(gbfs.Timestamp(time.Now().Unix()))
	f.SetTTL(5)
	r := ValidateFeedFreeBikeStatus(f, "")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data = &gbfs.FeedFreeBikeStatusData{}
	r = ValidateFeedFreeBikeStatus(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data.Bikes = []*gbfs.FeedFreeBikeStatusBike{}
	r = ValidateFeedFreeBikeStatus(f, "")
	if hasError(r.Errors, ErrRequired) || len(r.Errors) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	bike := &gbfs.FeedFreeBikeStatusBike{}
	f.Data.Bikes = []*gbfs.FeedFreeBikeStatusBike{bike}
	r = ValidateFeedFreeBikeStatus(f, "")
	if !hasError(r.Errors, ErrInvalidValue) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	bike.BikeID = gbfs.NewID("123")
	bike.IsReserved = gbfs.NewBoolean(false)
	r = ValidateFeedFreeBikeStatus(f, "")
	if errorCount(r.Errors, ErrRequired) != 3 {
		t.Errorf("expected 3 error of [%s], got %v", ErrRequired, r.Errors)
		return
	}
	bike.IsDisabled = gbfs.NewBoolean(false)
	r = ValidateFeedFreeBikeStatus(f, "")
	if errorCount(r.Errors, ErrRequired) != 2 {
		t.Errorf("expected 2 errors of [%s], got %v", ErrRequired, r.Errors)
		return
	}
	bike.Lat = gbfs.NewCoordinate(0)
	bike.Lon = gbfs.NewCoordinate(0)
	r = ValidateFeedFreeBikeStatus(f, "")
	if errorCount(r.Warnings, ErrZeroCoordinates) != 1 {
		t.Errorf("expected 1 warning of [%s], got %v", ErrZeroCoordinates, r.Warnings)
		return
	}
	bike.Lat = gbfs.NewCoordinate(91)
	bike.Lon = gbfs.NewCoordinate(181)
	r = ValidateFeedFreeBikeStatus(f, "")
	if errorCount(r.Errors, ErrOutOfRange) != 2 {
		t.Errorf("expected s errors of [%s], got %v", ErrOutOfRange, r.Errors)
		return
	}
	bike.Lat = gbfs.NewCoordinate(48.7162)
	bike.Lon = gbfs.NewCoordinate(21.2613)
	r = ValidateFeedFreeBikeStatus(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	bike.RentalURIs = &gbfs.RentalURIs{}
	r = ValidateFeedFreeBikeStatus(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	bike.RentalURIs = &gbfs.RentalURIs{
		Android: nil,
		IOS:     nil,
		Web:     nil,
	}
	r = ValidateFeedFreeBikeStatus(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	if errorCount(r.Infos, ErrAvailableFromVersion) != 1 {
		t.Errorf("expected 1 info of [%s], got %v", ErrAvailableFromVersion, r.Infos)
		return
	}
	bike.RentalURIs.Android = gbfs.NewString("http://")
	bike.RentalURIs.IOS = gbfs.NewString("http://")
	bike.RentalURIs.Web = gbfs.NewString("http://")
	r = ValidateFeedFreeBikeStatus(f, "1.1")
	if errorCount(r.Errors, ErrInvalidValue) != 3 {
		t.Errorf("expected 3 errors of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	if errorCount(r.Infos, ErrAvailableFromVersion) != 0 {
		t.Errorf("unexpected infos %v", r.Infos)
		return
	}
	bike.RentalURIs.Android = gbfs.NewString("http://localhost")
	bike.RentalURIs.IOS = gbfs.NewString("http://localhost")
	bike.RentalURIs.Web = gbfs.NewString("http://localhost")
	r = ValidateFeedFreeBikeStatus(f, "1.1")
	if errorCount(r.Errors, ErrInvalidValue) != 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	bike.VehicleTypeID = gbfs.NewID("")
	bike.LastReported = gbfs.NewTimestamp(0)
	bike.CurrentRangeMeters = gbfs.NewFloat64(0)
	r = ValidateFeedFreeBikeStatus(f, "1.1")
	if errorCount(r.Infos, ErrAvailableFromVersion) != 3 {
		t.Errorf("expected 3 infos of [%s], got %v", ErrAvailableFromVersion, r.Infos)
		return
	}
	bike.VehicleTypeID = gbfs.NewID("vehicleType1")
	bike.LastReported = gbfs.NewTimestamp(123456789)
	bike.CurrentRangeMeters = gbfs.NewFloat64(100000)
	r = ValidateFeedFreeBikeStatus(f, "2.1")
	if errorCount(r.Infos, ErrAvailableFromVersion) != 0 {
		t.Errorf("unexpected infos %v", r.Infos)
		return
	}
}
