package gbfs

type (
	// FeedSystemCalendar ...
	FeedSystemCalendar struct {
		FeedCommon
		Data *FeedSystemCalendarData `json:"data"`
	}
	// FeedSystemCalendarData ...
	FeedSystemCalendarData struct {
		Calendars []*FeedSystemCalendarCalendar `json:"calendars"`
	}
	// FeedSystemCalendarCalendar ...
	FeedSystemCalendarCalendar struct {
		StartMonth *int64 `json:"start_month"`
		StartDay   *int64 `json:"start_day"`
		StartYear  *int64 `json:"start_year,omitempty"`
		EndMonth   *int64 `json:"end_month"`
		EndDay     *int64 `json:"end_day"`
		EndYear    *int64 `json:"end_year,omitempty"`
	}
)

// Name ...
func (f *FeedSystemCalendar) Name() string {
	return FeedNameSystemCalendar
}
