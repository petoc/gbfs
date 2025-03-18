package gbfs

import (
	"encoding/json"
	"strconv"
	"strings"
	"sync"
	"time"
)

type (
	Feed interface {
		Name() string
		GetLastUpdated() Timestamp
		SetLastUpdated(Timestamp) Feed
		GetTTL() int
		SetTTL(int) Feed
		GetVersion() string
		SetVersion(string) Feed
		GetData() any
		SetData(any) Feed
		Expired() bool
	}
	FeedCommon struct {
		sync.RWMutex
		LastUpdated *Timestamp `json:"last_updated"`
		TTL         *int       `json:"ttl"`
		Version     *string    `json:"version,omitempty"`
		Data        any        `json:"data"`
	}
	RentalURIs struct {
		Android *string `json:"android,omitempty"`
		IOS     *string `json:"ios,omitempty"`
		Web     *string `json:"web,omitempty"`
	}
	RentalApps struct {
		Android *RentalApp `json:"android,omitempty"`
		IOS     *RentalApp `json:"ios,omitempty"`
	}
	RentalApp struct {
		StoreURI     *string `json:"store_uri,omitempty"`
		DiscoveryURI *string `json:"discovery_uri,omitempty"`
	}
	LocalizedString struct {
		Text     string `json:"text"`
		Language string `json:"language"`
	}
	VehicleTypeCapacity struct {
		VehicleTypeID *ID    `json:"vehicle_type_id"`
		Count         *int64 `json:"count"`
	}
	VehicleTypesCapacity struct {
		VehicleTypeIDs []*ID  `json:"vehicle_type_ids"`
		Count          *int64 `json:"count"`
	}
	PerUnitPricing struct {
		Start    *int64   `json:"start"`
		Rate     *float64 `json:"rate"`
		Interval *int64   `json:"interval"`
		End      *int64   `json:"end,omitempty"`
	}
)

func NewInt64(v int64) *int64 {
	return &v
}

func NewFloat64(v float64) *float64 {
	return &v
}

func NewString(v string) *string {
	return &v
}

func NewLocalizedString(t, l string) *LocalizedString {
	return &LocalizedString{
		Text:     t,
		Language: l,
	}
}

func NewVehicleTypeCapacity(v string, c int64) *VehicleTypeCapacity {
	return &VehicleTypeCapacity{
		VehicleTypeID: NewID(v),
		Count:         NewInt64(c),
	}
}

func NewVehicleTypesCapacity(v []string, c int64) *VehicleTypesCapacity {
	ids := []*ID{}
	for _, vt := range v {
		ids = append(ids, NewID(vt))
	}
	return &VehicleTypesCapacity{
		VehicleTypeIDs: ids,
		Count:          NewInt64(c),
	}
}

type Boolean bool

func (t *Boolean) UnmarshalJSON(b []byte) error {
	switch v := strings.ToLower(strings.Trim(string(b), `"`)); v {
	case "1", "true":
		*t = true
	default:
		*t = false
	}
	return nil
}

func NewBoolean(v bool) *Boolean {
	t := Boolean(v)
	return &t
}

type Timestamp string

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	*t = Timestamp(strings.Trim(string(b), `"`))
	return nil
}

func (t Timestamp) Time() (time.Time, error) {
	return time.Parse(time.RFC3339, string(t))
}

func NewTimestamp(v string) *Timestamp {
	t := Timestamp(v)
	return &t
}

type ID string

func (t *ID) UnmarshalJSON(b []byte) error {
	*t = ID(strings.Trim(string(b), `"`))
	return nil
}

func NewID(v string) *ID {
	t := ID(v)
	return &t
}

type Price struct {
	Float64 float64
	OldType string
}

func (p *Price) UnmarshalJSON(b []byte) error {
	tv := strings.Trim(string(b), `"`)
	f, err := strconv.ParseFloat(tv, 64)
	if err != nil {
		return err
	}
	p.Float64 = f
	if tv != string(b) {
		p.OldType = "string"
	}
	return nil
}

func (p *Price) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Float64)
}

func (p *Price) String() string {
	return strconv.FormatFloat(p.Float64, 'f', -1, 64)
}

func NewPrice(v float64) *Price {
	t := Price{
		Float64: v,
	}
	return &t
}

type Coordinate struct {
	Float64 float64
	OldType string
}

func (p *Coordinate) UnmarshalJSON(b []byte) error {
	tv := strings.Trim(string(b), `"`)
	f, err := strconv.ParseFloat(tv, 64)
	if err != nil {
		return err
	}
	p.Float64 = f
	if tv != string(b) {
		p.OldType = "string"
	}
	return nil
}

func (p *Coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Float64)
}

func (p *Coordinate) String() string {
	return strconv.FormatFloat(p.Float64, 'f', -1, 64)
}

func NewCoordinate(v float64) *Coordinate {
	t := Coordinate{
		Float64: v,
	}
	return &t
}

const (
	DateFormat = "2006-01-02"
	TimeFormat = "15:04:05"

	V30 string = "3.0"

	FeedNameGbfs               = "gbfs"
	FeedNameGbfsVersions       = "gbfs_versions"
	FeedNameGeofencingZones    = "geofencing_zones"
	FeedNameManifest           = "manifest"
	FeedNameStationInformation = "station_information"
	FeedNameStationStatus      = "station_status"
	FeedNameSystemAlerts       = "system_alerts"
	FeedNameSystemInformation  = "system_information"
	FeedNameSystemPricingPlans = "system_pricing_plans"
	FeedNameSystemRegions      = "system_regions"
	FeedNameVehicleStatus      = "vehicle_status"
	FeedNameVehicleTypes       = "vehicle_types"
)

const (
	FormFactorBicycle         = "bicycle"
	FormFactorCar             = "car"
	FormFactorCargoBicycle    = "cargo_bicycle"
	FormFactorMoped           = "moped"
	FormFactorOther           = "other"
	FormFactorScooterSeated   = "scooter_seated"
	FormFactorScooterStanding = "scooter_standing"
)

func FormFactorAll() []string {
	return []string{
		FormFactorBicycle,
		FormFactorCar,
		FormFactorCargoBicycle,
		FormFactorMoped,
		FormFactorOther,
		FormFactorScooterStanding,
		FormFactorScooterSeated,
	}
}

const (
	PropulsionTypeCombustion       = "combustion"
	PropulsionTypeCombustionDiesel = "combustion_diesel"
	PropulsionTypeElectric         = "electric"
	PropulsionTypeElectricAssist   = "electric_assist"
	PropulsionTypeHuman            = "human"
	PropulsionTypeHybrid           = "hybrid"
	PropulsionTypeHydrogenFuelCell = "hydrogen_fuel_cell"
	PropulsionTypePlugInHybrid     = "plug_in_hybrid"
)

func PropulsionTypeAll() []string {
	return []string{
		PropulsionTypeCombustion,
		PropulsionTypeCombustionDiesel,
		PropulsionTypeElectric,
		PropulsionTypeElectricAssist,
		PropulsionTypeHuman,
		PropulsionTypeHybrid,
		PropulsionTypeHydrogenFuelCell,
		PropulsionTypePlugInHybrid,
	}
}

const (
	VehicleAccessoryAirConditioning = "air_conditioning"
	VehicleAccessoryAutomatic       = "automatic"
	VehicleAccessoryManual          = "manual"
	VehicleAccessoryConvertible     = "convertible"
	VehicleAccessoryCruiseControl   = "cruise_control"
	VehicleAccessoryDoors2          = "doors_2"
	VehicleAccessoryDoors3          = "doors_3"
	VehicleAccessoryDoors4          = "doors_4"
	VehicleAccessoryDoors5          = "doors_5"
	VehicleAccessoryNavigation      = "navigation"
)

func VehicleAccessoryAll() []string {
	return []string{
		VehicleAccessoryAirConditioning,
		VehicleAccessoryAutomatic,
		VehicleAccessoryManual,
		VehicleAccessoryConvertible,
		VehicleAccessoryCruiseControl,
		VehicleAccessoryDoors2,
		VehicleAccessoryDoors3,
		VehicleAccessoryDoors4,
		VehicleAccessoryDoors5,
		VehicleAccessoryNavigation,
	}
}

const (
	VehicleEquipmentChildSeatA  = "child_seat_a"
	VehicleEquipmentChildSeatB  = "child_seat_b"
	VehicleEquipmentChildSeatC  = "child_seat_c"
	VehicleEquipmentSnowChains  = "snow_chains"
	VehicleEquipmentWinterTires = "winter_tires"
)

func VehicleEquipmentAll() []string {
	return []string{
		VehicleEquipmentChildSeatA,
		VehicleEquipmentChildSeatB,
		VehicleEquipmentChildSeatC,
		VehicleEquipmentSnowChains,
		VehicleEquipmentWinterTires,
	}
}

const (
	ReturnConstraintFreeFloating     = "free_floating"
	ReturnConstraintRoundtripStation = "roundtrip_station"
	ReturnConstraintAnyStation       = "any_station"
	ReturnConstraintHybrid           = "hybrid"
)

func ReturnConstraintAll() []string {
	return []string{
		ReturnConstraintFreeFloating,
		ReturnConstraintRoundtripStation,
		ReturnConstraintAnyStation,
		ReturnConstraintHybrid,
	}
}

const (
	AlertTypeSystemClosure  = "system_closure"
	AlertTypeStationClosure = "station_closure"
	AlertTypeStationMove    = "station_move"
	AlertTypeOther          = "other"
)

func AlertTypeAll() []string {
	return []string{
		AlertTypeSystemClosure,
		AlertTypeStationClosure,
		AlertTypeStationMove,
		AlertTypeOther,
	}
}

const (
	RentalMethodKey           = "key"
	RentalMethodCreditCard    = "creditcard"
	RentalMethodPayPass       = "paypass"
	RentalMethodApplePay      = "applepay"
	RentalMethodAndroidPay    = "androidpay"
	RentalMethodTransitCard   = "transitcard"
	RentalMethodAccountNumber = "accountnumber"
	RentalMethodPhone         = "phone"
)

func RentalMethodAll() []string {
	return []string{
		RentalMethodKey,
		RentalMethodCreditCard,
		RentalMethodPayPass,
		RentalMethodApplePay,
		RentalMethodAndroidPay,
		RentalMethodTransitCard,
		RentalMethodAccountNumber,
		RentalMethodPhone,
	}
}

const (
	ParkingTypeParkingLot         = "parking_lot"
	ParkingTypeStreetParking      = "street_parking"
	ParkingTypeUndergroundParking = "underground_parking"
	ParkingTypeSidewalkParking    = "sidewalk_parking"
	ParkingTypeOther              = "other"
)

func ParkingTypeAll() []string {
	return []string{
		ParkingTypeParkingLot,
		ParkingTypeStreetParking,
		ParkingTypeUndergroundParking,
		ParkingTypeSidewalkParking,
		ParkingTypeOther,
	}
}

type GeoJSONGeometry struct {
	Type        string `json:"type"`
	Coordinates any    `json:"coordinates"`
	Properties  any    `json:"properties,omitempty"`
}

type GeoJSONFeature struct {
	Type       string           `json:"type"`
	Geometry   *GeoJSONGeometry `json:"geometry"`
	Properties any              `json:"properties,omitempty"`
}

type GeoJSONFeatureCollection struct {
	Type     string            `json:"type"`
	Features []*GeoJSONFeature `json:"features"`
}

func NewGeoJSONFeatureCollection(features []*GeoJSONFeature) *GeoJSONFeatureCollection {
	return &GeoJSONFeatureCollection{
		Type:     "FeatureCollection",
		Features: features,
	}
}

func NewGeoJSONFeature(geometry *GeoJSONGeometry, properties any) *GeoJSONFeature {
	return &GeoJSONFeature{
		Type:       "Feature",
		Geometry:   geometry,
		Properties: properties,
	}
}

func NewGeoJSONGeometryMultiPolygon(coordinates any, properties any) *GeoJSONGeometry {
	return &GeoJSONGeometry{
		Type:        "MultiPolygon",
		Coordinates: coordinates,
		Properties:  properties,
	}
}

func FeedNameAll() []string {
	return []string{
		FeedNameGbfs,
		FeedNameGbfsVersions,
		FeedNameGeofencingZones,
		FeedNameManifest,
		FeedNameStationInformation,
		FeedNameStationStatus,
		FeedNameSystemAlerts,
		FeedNameSystemInformation,
		FeedNameSystemPricingPlans,
		FeedNameSystemRegions,
		FeedNameVehicleStatus,
		FeedNameVehicleTypes,
	}
}

func FeedStruct(name string) Feed {
	switch name {
	case FeedNameGbfs:
		return &FeedGbfs{}
	case FeedNameGbfsVersions:
		return &FeedGbfsVersions{}
	case FeedNameGeofencingZones:
		return &FeedGeofencingZones{}
	case FeedNameManifest:
		return &FeedManifest{}
	case FeedNameStationInformation:
		return &FeedStationInformation{}
	case FeedNameStationStatus:
		return &FeedStationStatus{}
	case FeedNameSystemAlerts:
		return &FeedSystemAlerts{}
	case FeedNameSystemInformation:
		return &FeedSystemInformation{}
	case FeedNameSystemPricingPlans:
		return &FeedSystemPricingPlans{}
	case FeedNameSystemRegions:
		return &FeedSystemRegions{}
	case FeedNameVehicleStatus:
		return &FeedVehicleStatus{}
	case FeedNameVehicleTypes:
		return &FeedVehicleTypes{}
	}
	return nil
}

func (s *FeedCommon) Name() string {
	return ""
}

func (s *FeedCommon) GetLastUpdated() Timestamp {
	if s.LastUpdated == nil {
		return Timestamp("2000-01-01T00:00:00Z")
	}
	return *s.LastUpdated
}

func (s *FeedCommon) SetLastUpdated(v Timestamp) Feed {
	s.LastUpdated = &v
	return s
}

func (s *FeedCommon) GetTTL() int {
	if s.TTL == nil {
		return 0
	}
	return *s.TTL
}

func (s *FeedCommon) SetTTL(v int) Feed {
	s.TTL = &v
	return s
}

func (s *FeedCommon) GetVersion() string {
	if s.Version == nil {
		return ""
	}
	return *s.Version
}

func (s *FeedCommon) SetVersion(v string) Feed {
	s.Version = &v
	return s
}

func (s *FeedCommon) GetData() any {
	return s.Data
}

func (s *FeedCommon) SetData(v any) Feed {
	s.Data = v
	return s
}

func (s *FeedCommon) Expired() bool {
	if s.TTL == nil || *s.TTL == 0 {
		return false
	}
	if s.LastUpdated == nil {
		return true
	}
	lastUpdated, err := s.LastUpdated.Time()
	if err != nil {
		return true
	}
	return lastUpdated.Add(time.Duration(*s.TTL) * time.Second).Before(time.Now())
}
