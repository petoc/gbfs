package validator

import (
	"testing"
	"time"

	"github.com/petoc/gbfs"
)

func TestValidateFeedStationInformation(t *testing.T) {
	f := &gbfs.FeedStationInformation{}
	f.SetLastUpdated(gbfs.Timestamp(time.Now().Unix()))
	f.SetTTL(5)
	r := ValidateFeedStationInformation(f, "")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data = &gbfs.FeedStationInformationData{}
	r = ValidateFeedStationInformation(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data.Stations = []*gbfs.FeedStationInformationStation{}
	r = ValidateFeedStationInformation(f, "")
	if hasError(r.Errors, ErrRequired) || len(r.Errors) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	station := &gbfs.FeedStationInformationStation{}
	f.Data.Stations = []*gbfs.FeedStationInformationStation{station}
	r = ValidateFeedStationInformation(f, "")
	if !hasError(r.Errors, ErrInvalidValue) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	station.StationID = gbfs.NewID("123")
	r = ValidateFeedStationInformation(f, "")
	if errorCount(r.Errors, ErrRequired) != 3 {
		t.Errorf("expected 3 errors of [%s], got %v", ErrRequired, r.Errors)
		return
	}
	station.Name = gbfs.NewString("Station Name")
	station.Lat = gbfs.NewCoordinate(0)
	station.Lon = gbfs.NewCoordinate(0)
	r = ValidateFeedStationInformation(f, "")
	if errorCount(r.Warnings, ErrZeroCoordinates) != 1 {
		t.Errorf("expected 1 warning of [%s], got %v", ErrZeroCoordinates, r.Warnings)
		return
	}
	station.Lat = gbfs.NewCoordinate(91)
	station.Lon = gbfs.NewCoordinate(181)
	r = ValidateFeedStationInformation(f, "")
	if errorCount(r.Errors, ErrOutOfRange) != 2 {
		t.Errorf("expected 2 errors of [%s], got %v", ErrOutOfRange, r.Errors)
		return
	}
	station.Lat = gbfs.NewCoordinate(48.7162)
	station.Lon = gbfs.NewCoordinate(21.2613)
	r = ValidateFeedStationInformation(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	station.ShortName = gbfs.NewString("")
	station.Address = gbfs.NewString("")
	station.CrossStreet = gbfs.NewString("")
	station.RegionID = gbfs.NewID("")
	station.PostCode = gbfs.NewString("")
	station.Capacity = gbfs.NewInt64(-1)
	r = ValidateFeedStationInformation(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 6 {
		t.Errorf("expected 6 errors of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	station.ShortName = gbfs.NewString("Station")
	station.Address = gbfs.NewString("Address")
	station.CrossStreet = gbfs.NewString("CrossStreet")
	station.RegionID = gbfs.NewID("region1")
	station.PostCode = gbfs.NewString("12345")
	station.Capacity = gbfs.NewInt64(1)
	r = ValidateFeedStationInformation(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	station.RentalMethods = []string{"invalid"}
	r = ValidateFeedStationInformation(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	station.RentalMethods = []string{gbfs.RentalMethodCreditCard}
	station.RentalURIs = &gbfs.RentalURIs{}
	r = ValidateFeedStationInformation(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	station.RentalURIs = &gbfs.RentalURIs{
		Android: nil,
		IOS:     nil,
		Web:     nil,
	}
	r = ValidateFeedStationInformation(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	if errorCount(r.Infos, ErrAvailableFromVersion) != 1 {
		t.Errorf("expected 1 info of [%s], got %v", ErrAvailableFromVersion, r.Infos)
		return
	}
	f.SetVersion("1.1")
	station.RentalURIs.Android = gbfs.NewString("http://")
	station.RentalURIs.IOS = gbfs.NewString("http://")
	station.RentalURIs.Web = gbfs.NewString("http://")
	r = ValidateFeedStationInformation(f, "1.1")
	if errorCount(r.Errors, ErrInvalidValue) != 3 {
		t.Errorf("expected 3 errors of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	if errorCount(r.Infos, ErrAvailableFromVersion) != 0 {
		t.Errorf("unexpected infos %v", r.Infos)
		return
	}
	station.RentalURIs.Android = gbfs.NewString("http://localhost")
	station.RentalURIs.IOS = gbfs.NewString("http://localhost")
	station.RentalURIs.Web = gbfs.NewString("http://localhost")
	r = ValidateFeedStationInformation(f, "1.1")
	if errorCount(r.Errors, ErrInvalidValue) != 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	station.IsVirtualStation = gbfs.NewBoolean(false)
	station.StationArea = gbfs.NewGeoJSONGeometryMultiPolygon(nil, nil)
	station.VehicleCapacity = map[*gbfs.ID]int64{}
	station.IsValetStation = gbfs.NewBoolean(false)
	station.VehicleTypeCapacity = map[*gbfs.ID]int64{}
	r = ValidateFeedStationInformation(f, "1.1")
	if errorCount(r.Infos, ErrAvailableFromVersion) != 5 {
		t.Errorf("expected 5 infos of [%s], got %v", ErrAvailableFromVersion, r.Infos)
		return
	}
	r = ValidateFeedStationInformation(f, "2.1")
	if errorCount(r.Errors, ErrInvalidValue) != 3 {
		t.Errorf("expected 3 errors of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	station.StationArea.Type = "Polygon"
	r = ValidateFeedStationInformation(f, "2.1")
	if errorCount(r.Errors, ErrInvalidType) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrInvalidType, r.Errors)
		return
	}
	station.StationArea = gbfs.NewGeoJSONGeometryMultiPolygon(
		[][][][]float64{
			[][][]float64{
				[][]float64{
					[]float64{16.8331891, 47.7314286},
					[]float64{22.56571, 47.7314286},
					[]float64{22.56571, 49.6138162},
					[]float64{16.8331891, 49.6138162},
					[]float64{16.8331891, 47.7314286},
				},
			},
		},
		nil,
	)
	station.VehicleCapacity = map[*gbfs.ID]int64{
		gbfs.NewID(""): 0,
	}
	station.VehicleTypeCapacity = map[*gbfs.ID]int64{
		gbfs.NewID(""): 0,
	}
	r = ValidateFeedStationInformation(f, "2.1")
	if errorCount(r.Errors, ErrInvalidValue) != 2 {
		t.Errorf("expected 2 errors of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	station.VehicleCapacity = map[*gbfs.ID]int64{
		gbfs.NewID("vehicleType1"): -1,
	}
	station.VehicleTypeCapacity = map[*gbfs.ID]int64{
		gbfs.NewID("vehicleType1"): -1,
	}
	r = ValidateFeedStationInformation(f, "2.1")
	if errorCount(r.Errors, ErrInvalidValue) != 2 {
		t.Errorf("expected 2 errors of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	station.VehicleCapacity = map[*gbfs.ID]int64{
		gbfs.NewID("vehicleType1"): 0,
	}
	station.VehicleTypeCapacity = map[*gbfs.ID]int64{
		gbfs.NewID("vehicleType1"): 0,
	}
	r = ValidateFeedStationInformation(f, "2.1")
	if errorCount(r.Errors, ErrInvalidValue) != 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
}
