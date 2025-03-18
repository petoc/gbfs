package gbfs

type (
	FeedVehicleTypes struct {
		FeedCommon
		Data *FeedVehicleTypesData `json:"data"`
	}
	FeedVehicleTypesData struct {
		VehicleTypes []*FeedVehicleTypesVehicleType `json:"vehicle_types"`
	}
	FeedVehicleTypesVehicleType struct {
		VehicleTypeID        *ID                `json:"vehicle_type_id"`
		FormFactor           *string            `json:"form_factor"`
		RiderCapacity        *int64             `json:"rider_capacity,omitempty"`
		CargoVolumeCapacity  *int64             `json:"cargo_volume_capacity,omitempty"` // l
		CargoLoadCapacity    *int64             `json:"cargo_load_capacity,omitempty"`   // kg
		PropulsionType       *string            `json:"propulsion_type"`
		EcoLabels            []*EcoLabel        `json:"eco_labels,omitempty"`
		MaxRangeMeters       *float64           `json:"max_range_meters,omitempty"`
		Name                 []*LocalizedString `json:"name,omitempty"`
		VehicleAccessories   []string           `json:"vehicle_accessories,omitempty"`
		GCO2Km               *int64             `json:"g_CO2_km,omitempty"`
		VehicleImage         *string            `json:"vehicle_image,omitempty"` // jpg, png
		Make                 []*LocalizedString `json:"make,omitempty"`
		Model                []*LocalizedString `json:"model,omitempty"`
		Color                *string            `json:"color,omitempty"`
		Description          []*LocalizedString `json:"description,omitempty"`
		WheelCount           *int64             `json:"wheel_count,omitempty"`
		MaxPermittedSpeed    *int64             `json:"max_permitted_speed,omitempty"`
		RatedPower           *int64             `json:"rated_power,omitempty"`
		DefaultReserveTime   *int64             `json:"default_reserve_time,omitempty"`
		ReturnConstraint     *string            `json:"return_constraint,omitempty"`
		VehicleAssets        []*VehicleAsset    `json:"vehicle_assets,omitempty"`
		DefaultPricingPlanID *ID                `json:"default_pricing_plan_id,omitempty"`
		PricingPlanIDs       []*ID              `json:"pricing_plan_ids,omitempty"`
	}
	EcoLabel struct {
		CountryCode *string `json:"country_code"`
		EcoSticker  *string `json:"eco_sticker"`
	}
	VehicleAsset struct {
		IconURL          *string `json:"icon_url"`
		IconURLDark      *string `json:"icon_url_dark,omitempty"`
		IconLastModified *string `json:"icon_last_modified"`
	}
)

func (f *FeedVehicleTypes) Name() string {
	return FeedNameVehicleTypes
}
