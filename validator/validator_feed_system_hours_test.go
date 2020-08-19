package validator

import (
	"testing"
	"time"

	"github.com/petoc/gbfs"
)

func TestValidateFeedSystemHours(t *testing.T) {
	f := &gbfs.FeedSystemHours{}
	f.SetLastUpdated(gbfs.Timestamp(time.Now().Unix()))
	f.SetTTL(5)
	r := ValidateFeedSystemHours(f, "")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data = &gbfs.FeedSystemHoursData{}
	r = ValidateFeedSystemHours(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data.RentalHours = []*gbfs.FeedSystemHoursRentalHour{}
	r = ValidateFeedSystemHours(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	rentalHour := &gbfs.FeedSystemHoursRentalHour{}
	f.Data.RentalHours = append(f.Data.RentalHours, rentalHour)
	r = ValidateFeedSystemHours(f, "")
	if !hasError(r.Errors, ErrInvalidValue) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	rentalHour.UserTypes = []string{}
	r = ValidateFeedSystemHours(f, "")
	if errorCount(r.Errors, ErrRequired) != 4 {
		t.Errorf("expected 4 errors of [%s], got %v", ErrRequired, r.Errors)
		return
	}
	rentalHour.UserTypes = []string{"abc"}
	r = ValidateFeedSystemHours(f, "")
	if !hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	rentalHour.UserTypes = []string{gbfs.UserTypeMember}
	r = ValidateFeedSystemHours(f, "")
	if hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	rentalHour.Days = []string{}
	r = ValidateFeedSystemHours(f, "")
	if errorCount(r.Errors, ErrRequired) != 3 {
		t.Errorf("expected 3 errors of [%s], got %v", ErrRequired, r.Errors)
		return
	}
	rentalHour.Days = []string{"abc"}
	r = ValidateFeedSystemHours(f, "")
	if !hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	rentalHour.Days = []string{gbfs.DayMon}
	r = ValidateFeedSystemHours(f, "")
	if hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	rentalHour.StartTime = gbfs.NewString("00:00:00")
	rentalHour.EndTime = gbfs.NewString("23:59:59")
	r = ValidateFeedSystemHours(f, "")
	if hasError(r.Errors, ErrRequired) {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	rentalHour.StartTime = gbfs.NewString("00:00:60")
	rentalHour.EndTime = gbfs.NewString("24:59:60")
	r = ValidateFeedSystemHours(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 2 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
}
