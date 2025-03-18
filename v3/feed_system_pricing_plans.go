package gbfs

type (
	FeedSystemPricingPlans struct {
		FeedCommon
		Data *FeedSystemPricingPlansData `json:"data"`
	}
	FeedSystemPricingPlansData struct {
		Plans []*FeedSystemPricingPlansPricingPlan `json:"plans"`
	}
	FeedSystemPricingPlansPricingPlan struct {
		PlanID        *ID                `json:"plan_id"`
		URL           *string            `json:"url,omitempty"`
		Name          []*LocalizedString `json:"name"`
		Currency      *string            `json:"currency"`
		Price         *Price             `json:"price"`
		IsTaxable     *Boolean           `json:"is_taxable"`
		Description   []*LocalizedString `json:"description"`
		PerKmPricing  *PerUnitPricing    `json:"per_km_pricing,omitempty"`
		PerMinPricing *PerUnitPricing    `json:"per_min_pricing,omitempty"`
		SurgePricing  *Boolean           `json:"surge_pricing,omitempty"`
	}
)

func (f *FeedSystemPricingPlans) Name() string {
	return FeedNameSystemPricingPlans
}
