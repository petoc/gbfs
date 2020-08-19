package validator

import (
	"testing"
	"time"

	"github.com/petoc/gbfs"
)

func TestValidateFeedCommon(t *testing.T) {
	f := &gbfs.FeedGbfs{}
	r := ValidateFeedCommon(f, "")
	if !hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	f.SetLastUpdated(gbfs.Timestamp(time.Now().Unix()))
	r = ValidateFeedCommon(f, "")
	if hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("unexpected error [%s]", ErrInvalidValue)
		return
	}
	f.SetTTL(5)
	r = ValidateFeedCommon(f, "")
	if hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("unexpected error [%s]", ErrInvalidValue)
		return
	}
	f.SetTTL(-1)
	r = ValidateFeedCommon(f, "")
	if !hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	f.SetTTL(5)
	r = ValidateFeedCommon(f, "1.1")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.SetVersion("1.0")
	r = ValidateFeedCommon(f, "")
	if hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	f.SetVersion("1.0")
	r = ValidateFeedCommon(f, "1.1")
	if !hasError(r.Errors, ErrInconsistent) {
		t.Errorf("expected error [%s], got %v", ErrInconsistent, r.Errors)
		return
	}
	f.SetVersion("1.0x")
	r = ValidateFeedCommon(f, "")
	if !hasError(r.Errors, ErrInvalidValue) {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	f.SetTTL(-1)
	f.SetLastUpdated(0)
	r = ValidateFeedCommon(f, "")
	ee := []error{ErrInvalidValue, ErrInvalidValue, ErrInvalidValue}
	if !hasErrors(r.Errors, ee) || len(r.Errors) != 3 {
		t.Errorf("expected errors %v, got %v", ee, r.Errors)
		return
	}
}
