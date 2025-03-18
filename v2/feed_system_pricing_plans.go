package gbfs

type (
	// FeedSystemPricingPlans ...
	FeedSystemPricingPlans struct {
		FeedCommon
		Data *FeedSystemPricingPlansData `json:"data"`
	}
	// FeedSystemPricingPlansData ...
	FeedSystemPricingPlansData struct {
		Plans []*FeedSystemPricingPlansPricingPlan `json:"plans"`
	}
	// FeedSystemPricingPlansPricingPlan ...
	FeedSystemPricingPlansPricingPlan struct {
		PlanID      *ID      `json:"plan_id"`
		URL         *string  `json:"url,omitempty"`
		Name        *string  `json:"name"`
		Currency    *string  `json:"currency"`
		Price       *Price   `json:"price"`
		IsTaxable   *Boolean `json:"is_taxable"`
		Description *string  `json:"description"`
	}
)

// Name ...
func (f *FeedSystemPricingPlans) Name() string {
	return FeedNameSystemPricingPlans
}
