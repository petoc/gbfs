package validator

import (
	"encoding/json"
	"errors"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/petoc/gbfs"
)

var (
	ErrInvalidInput         = errors.New("invalid input")
	ErrInconsistent         = errors.New("inconsistent")
	ErrRequired             = errors.New("required")
	ErrConflict             = errors.New("conflict")
	ErrEmptyValue           = errors.New("empty value")
	ErrInvalidType          = errors.New("invalid type")
	ErrInvalidValue         = errors.New("invalid value")
	ErrOutOfRange           = errors.New("out of range")
	ErrAvailableFromVersion = errors.New("officially available from version")
	ErrZeroCoordinates      = errors.New("zero coordinates")
)

var (
	regexpVersion         = regexp.MustCompile(`^[\d]+\.[\d]+$`)
	regexpURL             = regexp.MustCompile(`^(https?:\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3]|24\d|25[0-5])(\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-5]))|(\[(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(((www\.)|([a-zA-Z0-9]+([-_\.]?[a-zA-Z0-9])*[a-zA-Z0-9]\.[a-zA-Z0-9]+))?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?(:([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5]))?((\/|\?|#)[^\s]*)?$`)
	regexpEmail           = regexp.MustCompile("^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$")
	regexpTimezone        = regexp.MustCompile(`^[a-zA-Z0-9_\-\+\/]+$`)
	regexpStoreURIAndroid = regexp.MustCompile(`^https:\/\/play.google.com\/store\/apps\/details\?id=[a-zA-Z0-9\._\-]+$`)
	regexpStoreURIIOS     = regexp.MustCompile(`^https:\/\/apps.apple.com(\/[a-z\-]{2,5})?/app\/.*?\/id[0-9]+$`)
	regexpDiscoveryURI    = regexp.MustCompile(`^[a-zA-Z0-9\._\-]+://[a-zA-Z0-9\._\-]*$`)
	regexpYear            = regexp.MustCompile(`^[0-9]{4}$`)
)

type (
	// Result ...
	Result struct {
		Language string    `json:"language,omitempty"`
		URL      string    `json:"url,omitempty"`
		Feed     gbfs.Feed `json:"feed,omitempty"`
		Infos    []error   `json:"infos,omitempty"`
		Warnings []error   `json:"warnings,omitempty"`
		Errors   []error   `json:"errors,omitempty"`
	}
)

// Info Add info
func (r *Result) Info(e error) *Result {
	r.Infos = append(r.Infos, e)
	return r
}

// InfoW Add wrapped info
func (r *Result) InfoW(m string, e error) *Result {
	return r.Info(ErrorWrap(m, e))
}

// InfoWS Add wrapped info with suffix
func (r *Result) InfoWS(m string, e error, s string) *Result {
	return r.Info(ErrorWrapSuffix(m, e, s))
}

// InfoWSP Add wrapped info with suffix in parenthesis
func (r *Result) InfoWSP(m string, e error, s string) *Result {
	return r.InfoWS(m, e, " ("+s+")")
}

// InfoWV Add wrapped info with version info
func (r *Result) InfoWV(m string, e error, s string) *Result {
	return r.InfoWS(m, e, " "+s)
}

// Warning Add warning
func (r *Result) Warning(e error) *Result {
	r.Warnings = append(r.Warnings, e)
	return r
}

// WarningW Add wrapped warning
func (r *Result) WarningW(m string, e error) *Result {
	return r.Warning(ErrorWrap(m, e))
}

// WarningWS Add wrapped warning with suffix
func (r *Result) WarningWS(m string, e error, s string) *Result {
	return r.Warning(ErrorWrapSuffix(m, e, s))
}

// WarningWSP Add wrapped warning with suffix in parenthesis
func (r *Result) WarningWSP(m string, e error, s string) *Result {
	return r.WarningWS(m, e, " ("+s+")")
}

// Error Add error
func (r *Result) Error(e error) *Result {
	r.Errors = append(r.Errors, e)
	return r
}

// ErrorW Add wrapped error
func (r *Result) ErrorW(m string, e error) *Result {
	return r.Error(ErrorWrap(m, e))
}

// ErrorWS Add wrapped error with suffix
func (r *Result) ErrorWS(m string, e error, s string) *Result {
	return r.Error(ErrorWrapSuffix(m, e, s))
}

// ErrorWSP Add wrapped error with suffix in parenthesis
func (r *Result) ErrorWSP(m string, e error, s string) *Result {
	return r.ErrorWS(m, e, " ("+s+")")
}

// HasInfos ...
func (r *Result) HasInfos() bool {
	if r.Infos != nil && len(r.Infos) > 0 {
		return true
	}
	return false
}

// HasWarnings ...
func (r *Result) HasWarnings() bool {
	if r.Warnings != nil && len(r.Warnings) > 0 {
		return true
	}
	return false
}

// HasErrors ...
func (r *Result) HasErrors() bool {
	if r.Errors != nil && len(r.Errors) > 0 {
		return true
	}
	return false
}

// ValidateFeedName ...
func ValidateFeedName(v string) bool {
	return inSlice(v, gbfs.FeedNameAll())
}

// ValidateVersion ...
func ValidateVersion(v string) bool {
	return regexpVersion.MatchString(v)
}

// ValidateFormFactor ...
func ValidateFormFactor(v *string) bool {
	return inSlice(*v, gbfs.FormFactorAll())
}

// ValidatePropulsionType ...
func ValidatePropulsionType(v *string) bool {
	return inSlice(*v, gbfs.PropulsionTypeAll())
}

// ValidateAlertType ...
func ValidateAlertType(v *string) bool {
	return inSlice(*v, gbfs.AlertTypeAll())
}

// ValidateRentalMethod ...
func ValidateRentalMethod(v string) bool {
	return inSlice(v, gbfs.RentalMethodAll())
}

// ValidateUserType ...
func ValidateUserType(v string) bool {
	return inSlice(v, gbfs.UserTypeAll())
}

// ValidateDay ...
func ValidateDay(v string) bool {
	return inSlice(v, gbfs.DayAll())
}

// FeedVersion ...
func FeedVersion(v *string) string {
	if v == nil || *v == "" {
		return gbfs.V10
	}
	return *v
}

func validateURI(v *string) bool {
	if v == nil {
		return false
	}
	u, err := url.ParseRequestURI(*v)
	if err != nil || u.Host == "" {
		return false
	}
	return true
}

func validateURL(v *string) bool {
	if v == nil {
		return false
	}
	if !regexpURL.MatchString(*v) {
		return false
	}
	u, err := url.ParseRequestURI(*v)
	if err != nil || u.Host == "" {
		return false
	}
	return validateURI(v)
}

func validateEmail(v *string) bool {
	if v == nil {
		return false
	}
	if regexpEmail.MatchString(*v) {
		return true
	}
	return false
}

func validateTimestamp(v gbfs.Timestamp) bool {
	if v == 0 || time.Unix(int64(v), 0).IsZero() {
		return false
	}
	return true
}

func validateTimePattern(l string, v *string) bool {
	if v == nil {
		return false
	}
	_, err := time.Parse(l, *v)
	if err != nil {
		return false
	}
	return true
}

func validateDate(v *string) bool {
	return validateTimePattern("2006-01-02", v)
}

func validateTime(v *string) bool {
	return validateTimePattern("15:04:05", v)
}

func validateIntYear(v *int64) bool {
	if v == nil {
		return false
	}
	if regexpYear.MatchString(strconv.FormatInt(*v, 10)) {
		return true
	}
	return false
}

func validateIntMonth(v *int64) bool {
	if v == nil {
		return false
	}
	return *v >= 1 && *v <= 12
}

func validateIntDay(v *int64) bool {
	if v == nil {
		return false
	}
	return *v >= 1 && *v <= 31
}

func validateTimezone(v *string) bool {
	if v == nil {
		return false
	}
	if regexpTimezone.MatchString(*v) {
		return true
	}
	return false
}

func validateStoreURIAndroid(v *string) bool {
	if v == nil {
		return false
	}
	if !regexpStoreURIAndroid.MatchString(*v) {
		return false
	}
	return validateURL(v)
}

func validateStoreURIIOS(v *string) bool {
	if v == nil {
		return false
	}
	if !regexpStoreURIIOS.MatchString(*v) {
		return false
	}
	return validateURL(v)
}

func validateDiscoveryURI(v *string) bool {
	if v == nil {
		return false
	}
	if !regexpDiscoveryURI.MatchString(*v) {
		return false
	}
	return true
}

func validateLatitude(v *float64) bool {
	if v == nil {
		return false
	}
	return *v >= -90 && *v <= 90
}

func validateLongitude(v *float64) bool {
	if v == nil {
		return false
	}
	return *v >= -180 && *v <= 180
}

func verEQ(v1, v2 string) bool {
	return v1 == v2
}

func verLT(v1, v2 string) bool {
	return strings.Compare(v1, v2) < 0
}

func verLE(v1, v2 string) bool {
	return strings.Compare(v1, v2) <= 0
}

func verGT(v1, v2 string) bool {
	return strings.Compare(v1, v2) > 0
}

func verGE(v1, v2 string) bool {
	return strings.Compare(v1, v2) >= 0
}

func hasError(errs []error, err error) bool {
	for _, e := range errs {
		if errors.Is(e, err) {
			return true
		}
	}
	return false
}

func hasErrors(errs []error, errs2 []error) bool {
	for _, e := range errs2 {
		if !hasError(errs, e) {
			return false
		}
	}
	return true
}

func errorCount(errs []error, err error) int {
	m := 0
	for _, e := range errs {
		if errors.Is(e, err) {
			m++
		}
	}
	return m
}

func msgPrefix(s string) string {
	if s != "" {
		s = s + ": "
	}
	return s
}

func msgSuffix(s string) string {
	if s != "" {
		s = ": " + s
	}
	return s
}

// ErrorWrap Wrap error
func ErrorWrap(msg string, err error) error {
	if err != nil {
		msg = msgPrefix(msg) + err.Error()
	}
	return &wrapError{msg, err}
}

// ErrorWrapSuffix Wrap error
func ErrorWrapSuffix(msg string, err error, suffix string) error {
	if err != nil {
		msg = msgPrefix(msg) + err.Error() + suffix
	}
	return &wrapError{msg, err}
}

type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err
}

func (e *wrapError) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.msg)
}

func nilOrEmpty(s1 interface{}) bool {
	return reflect.ValueOf(s1).IsNil() || reflect.DeepEqual(s1, reflect.New(reflect.TypeOf(s1).Elem()).Interface())
}

func nilOrZero(s interface{}) bool {
	return reflect.ValueOf(s).IsNil() || reflect.ValueOf(s).Len() == 0
}

func sliceIndexN(s string, i int) string {
	return s + "[" + strconv.Itoa(i) + "]"
}

func indexInSlice(n string, h []string) int {
	for k, v := range h {
		if n == v {
			return k
		}
	}
	return -1
}

func inSlice(n string, h []string) bool {
	return indexInSlice(n, h) > -1
}
