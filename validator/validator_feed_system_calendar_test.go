package validator

import (
	"testing"
	"time"

	"github.com/petoc/gbfs"
)

func TestValidateFeedSystemCalendar(t *testing.T) {
	f := &gbfs.FeedSystemCalendar{}
	f.SetLastUpdated(gbfs.Timestamp(time.Now().Unix()))
	f.SetTTL(5)
	r := ValidateFeedSystemCalendar(f, "")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data = &gbfs.FeedSystemCalendarData{}
	r = ValidateFeedSystemCalendar(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data.Calendars = []*gbfs.FeedSystemCalendarCalendar{}
	r = ValidateFeedSystemCalendar(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	calendar := &gbfs.FeedSystemCalendarCalendar{}
	f.Data.Calendars = []*gbfs.FeedSystemCalendarCalendar{calendar}
	r = ValidateFeedSystemCalendar(f, "")
	if !hasError(r.Errors, ErrInvalidValue) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	calendar.StartMonth = gbfs.NewInt64(1)
	calendar.StartDay = gbfs.NewInt64(1)
	calendar.StartYear = gbfs.NewInt64(2020)
	calendar.EndMonth = gbfs.NewInt64(12)
	calendar.EndDay = gbfs.NewInt64(31)
	calendar.EndYear = gbfs.NewInt64(2020)
	r = ValidateFeedSystemCalendar(f, "")
	if hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	calendar.StartMonth = gbfs.NewInt64(-1)
	calendar.StartDay = gbfs.NewInt64(-1)
	calendar.StartYear = gbfs.NewInt64(-1)
	calendar.EndMonth = gbfs.NewInt64(-1)
	calendar.EndDay = gbfs.NewInt64(-1)
	calendar.EndYear = gbfs.NewInt64(12020)
	r = ValidateFeedSystemCalendar(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 6 {
		t.Errorf("expected 6 errors of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
}
