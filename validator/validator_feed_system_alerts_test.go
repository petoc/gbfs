package validator

import (
	"testing"
	"time"

	"github.com/petoc/gbfs"
)

func TestValidateFeedSystemAlerts(t *testing.T) {
	f := &gbfs.FeedSystemAlerts{}
	f.SetLastUpdated(gbfs.Timestamp(time.Now().Unix()))
	f.SetTTL(5)
	r := ValidateFeedSystemAlerts(f, "")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data = &gbfs.FeedSystemAlertsData{}
	r = ValidateFeedSystemAlerts(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data.Alerts = []*gbfs.FeedSystemAlertsAlert{}
	r = ValidateFeedSystemAlerts(f, "")
	if hasError(r.Errors, ErrRequired) || len(r.Errors) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	alert := &gbfs.FeedSystemAlertsAlert{}
	f.Data.Alerts = []*gbfs.FeedSystemAlertsAlert{alert}
	r = ValidateFeedSystemAlerts(f, "")
	if !hasError(r.Errors, ErrInvalidValue) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	alert.AlertID = gbfs.NewID("123")
	r = ValidateFeedSystemAlerts(f, "")
	if errorCount(r.Errors, ErrRequired) != 2 {
		t.Errorf("expected 2 errors of [%s], got %v", ErrRequired, r.Errors)
		return
	}
	alert.Type = gbfs.NewString("")
	alert.Times = []*gbfs.FeedSystemAlertsAlertTime{}
	alert.StationIDs = []*gbfs.ID{}
	alert.RegionIDs = []*gbfs.ID{}
	alert.URL = gbfs.NewString("")
	alert.Summary = gbfs.NewString("")
	alert.Description = gbfs.NewString("")
	r = ValidateFeedSystemAlerts(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 3 {
		t.Errorf("expected 3 errors of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	alert.Type = gbfs.NewString(gbfs.AlertTypeSystemClosure)
	alertTime := &gbfs.FeedSystemAlertsAlertTime{}
	alert.Times = []*gbfs.FeedSystemAlertsAlertTime{alertTime}
	alert.StationIDs = []*gbfs.ID{gbfs.NewID("station1")}
	alert.RegionIDs = []*gbfs.ID{gbfs.NewID("region1")}
	alert.URL = gbfs.NewString("http://localhost/alert")
	alert.Summary = gbfs.NewString("Summary")
	alert.Description = gbfs.NewString("Description")
	r = ValidateFeedSystemAlerts(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	alertTime.End = gbfs.NewTimestamp(123456789)
	r = ValidateFeedSystemAlerts(f, "")
	if errorCount(r.Errors, ErrRequired) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrRequired, r.Errors)
		return
	}
	alertTime.Start = gbfs.NewTimestamp(0)
	r = ValidateFeedSystemAlerts(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 1 {
		t.Errorf("expected 1 error of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	alertTime.Start = gbfs.NewTimestamp(123456789)
	r = ValidateFeedSystemAlerts(f, "")
	if len(r.Errors) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
}
