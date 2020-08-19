package validator

import (
	"testing"
	"time"

	"github.com/petoc/gbfs"
)

func TestValidateFeedGeofencingZones(t *testing.T) {
	f := &gbfs.FeedGeofencingZones{}
	f.SetLastUpdated(gbfs.Timestamp(time.Now().Unix()))
	f.SetTTL(5)
	r := ValidateFeedGeofencingZones(f, "")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data = &gbfs.FeedGeofencingZonesData{}
	r = ValidateFeedGeofencingZones(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data.GeofencingZones = gbfs.NewFeedGeofencingZonesGeoJSONFeatureCollection(nil)
	f.Data.GeofencingZones.Type = ""
	r = ValidateFeedGeofencingZones(f, "")
	if errorCount(r.Errors, ErrInvalidType) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrInvalidType, r.Errors)
		return
	}
	f.Data.GeofencingZones = gbfs.NewFeedGeofencingZonesGeoJSONFeatureCollection(nil)
	r = ValidateFeedGeofencingZones(f, "")
	if errorCount(r.Errors, ErrRequired) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data.GeofencingZones = gbfs.NewFeedGeofencingZonesGeoJSONFeatureCollection([]*gbfs.FeedGeofencingZonesGeoJSONFeature{})
	r = ValidateFeedGeofencingZones(f, "")
	if len(r.Errors) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	feature := gbfs.NewFeedGeofencingZonesGeoJSONFeature(nil, nil)
	feature.Type = ""
	f.Data.GeofencingZones = gbfs.NewFeedGeofencingZonesGeoJSONFeatureCollection([]*gbfs.FeedGeofencingZonesGeoJSONFeature{feature})
	r = ValidateFeedGeofencingZones(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	feature.Type = "invalid"
	r = ValidateFeedGeofencingZones(f, "")
	if errorCount(r.Errors, ErrInvalidType) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrInvalidType, r.Errors)
		return
	}
	feature = gbfs.NewFeedGeofencingZonesGeoJSONFeature(nil, nil)
	f.Data.GeofencingZones = gbfs.NewFeedGeofencingZonesGeoJSONFeatureCollection([]*gbfs.FeedGeofencingZonesGeoJSONFeature{feature})
	r = ValidateFeedGeofencingZones(f, "")
	if errorCount(r.Errors, ErrRequired) != 2 {
		t.Errorf("expected 2 errors of [%s], got %v", ErrRequired, r.Errors)
		return
	}
	feature.Geometry = gbfs.NewGeoJSONGeometryMultiPolygon(nil, nil)
	feature.Properties = &gbfs.FeedGeofencingZonesGeoJSONFeatureProperties{}
	r = ValidateFeedGeofencingZones(f, "")
	if errorCount(r.Errors, ErrRequired) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrRequired, r.Errors)
		return
	}
	feature.Geometry.Coordinates = [][][][]float64{
		[][][]float64{
			[][]float64{
				[]float64{16.8331891, 47.7314286},
				[]float64{22.56571, 47.7314286},
				[]float64{22.56571, 49.6138162},
				[]float64{16.8331891, 49.6138162},
				[]float64{16.8331891, 47.7314286},
			},
		},
	}
	r = ValidateFeedGeofencingZones(f, "")
	if len(r.Errors) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	feature.Properties.Name = gbfs.NewString("")
	feature.Properties.Start = gbfs.NewTimestamp(0)
	feature.Properties.End = gbfs.NewTimestamp(0)
	feature.Properties.Rules = []*gbfs.FeedGeofencingZonesGeoJSONFeaturePropertiesRule{}
	r = ValidateFeedGeofencingZones(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 3 {
		t.Errorf("expected 3 errors of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	feature.Properties.Name = gbfs.NewString("Rule")
	feature.Properties.Start = gbfs.NewTimestamp(123456789)
	feature.Properties.End = gbfs.NewTimestamp(123456789)
	rule := &gbfs.FeedGeofencingZonesGeoJSONFeaturePropertiesRule{}
	feature.Properties.Rules = []*gbfs.FeedGeofencingZonesGeoJSONFeaturePropertiesRule{rule}
	r = ValidateFeedGeofencingZones(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	rule.VehicleTypeIDs = []*gbfs.ID{}
	r = ValidateFeedGeofencingZones(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	if errorCount(r.Errors, ErrRequired) != 2 {
		t.Errorf("expected 2 errors of [%s], got %v", ErrRequired, r.Errors)
		return
	}
	rule.VehicleTypeIDs = []*gbfs.ID{gbfs.NewID("")}
	rule.RideAllowed = gbfs.NewBoolean(false)
	rule.RideThroughAllowed = gbfs.NewBoolean(false)
	rule.MaximumSpeedKph = gbfs.NewInt64(-1)
	r = ValidateFeedGeofencingZones(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 2 {
		t.Errorf("expected 2 errors of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	rule.VehicleTypeIDs = []*gbfs.ID{gbfs.NewID("vehicleType1")}
	rule.MaximumSpeedKph = gbfs.NewInt64(15)
	r = ValidateFeedGeofencingZones(f, "")
	if len(r.Errors) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
}
