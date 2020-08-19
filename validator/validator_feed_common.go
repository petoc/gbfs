package validator

import "github.com/petoc/gbfs"

// ValidateFeedCommon ...
func ValidateFeedCommon(f gbfs.Feed, version string) *Result {
	r := &Result{
		Feed: f,
	}
	if !validateTimestamp(f.GetLastUpdated()) {
		r.ErrorWSP("last_updated", ErrInvalidValue, "POSIX time")
	}
	if f.GetTTL() < 0 {
		r.ErrorWSP("ttl", ErrInvalidValue, "non-negative integer")
	}
	if f.GetVersion() != "" {
		if !ValidateVersion(f.GetVersion()) {
			r.ErrorW("version", ErrInvalidValue)
			return r
		}
	}
	if version != "" {
		if verGE(version, gbfs.V11) {
			if f.GetVersion() == "" {
				r.ErrorWSP("version", ErrRequired, gbfs.V11)
				return r
			}
			if verLT(f.GetVersion(), version) {
				r.ErrorW("version", ErrInconsistent)
				return r
			}
		}
	}
	return r
}
