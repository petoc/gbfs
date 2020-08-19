package validator

import (
	"testing"
	"time"

	"github.com/petoc/gbfs"
)

func TestValidateFeedSystemPricingPlans(t *testing.T) {
	f := &gbfs.FeedSystemPricingPlans{}
	f.SetLastUpdated(gbfs.Timestamp(time.Now().Unix()))
	f.SetTTL(5)
	r := ValidateFeedSystemPricingPlans(f, "")
	if !hasError(r.Errors, ErrRequired) {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data = &gbfs.FeedSystemPricingPlansData{}
	r = ValidateFeedSystemPricingPlans(f, "")
	if !hasError(r.Errors, ErrRequired) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrRequired, r.Errors)
		return
	}
	f.Data.Plans = []*gbfs.FeedSystemPricingPlansPricingPlan{}
	r = ValidateFeedSystemPricingPlans(f, "")
	if hasError(r.Errors, ErrRequired) || len(r.Errors) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
	plan := &gbfs.FeedSystemPricingPlansPricingPlan{}
	f.Data.Plans = []*gbfs.FeedSystemPricingPlansPricingPlan{plan}
	r = ValidateFeedSystemPricingPlans(f, "")
	if !hasError(r.Errors, ErrInvalidValue) || len(r.Errors) != 1 {
		t.Errorf("expected error [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	plan.PlanID = gbfs.NewID("123")
	r = ValidateFeedSystemPricingPlans(f, "")
	if errorCount(r.Errors, ErrRequired) != 5 {
		t.Errorf("expected 5 errors of [%s], got %v", ErrRequired, r.Errors)
		return
	}
	plan.URL = gbfs.NewString("")
	plan.Name = gbfs.NewString("")
	plan.Currency = gbfs.NewString("")
	plan.Price = gbfs.NewPrice(-1)
	plan.IsTaxable = gbfs.NewBoolean(true)
	plan.Description = gbfs.NewString("")
	r = ValidateFeedSystemPricingPlans(f, "")
	if errorCount(r.Errors, ErrInvalidValue) != 5 {
		t.Errorf("expected 5 errors of [%s], got %v", ErrInvalidValue, r.Errors)
		return
	}
	plan.URL = gbfs.NewString("http://localhost/pricing-plan")
	plan.Name = gbfs.NewString("Pricing plan")
	plan.Currency = gbfs.NewString("EUR")
	plan.Price = gbfs.NewPrice(1)
	plan.IsTaxable = gbfs.NewBoolean(false)
	plan.Description = gbfs.NewString("Description")
	r = ValidateFeedSystemPricingPlans(f, "")
	if errorCount(r.Errors, ErrInvalidValue) > 0 {
		t.Errorf("unexpected errors %v", r.Errors)
		return
	}
}
