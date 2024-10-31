// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package publicdashboard

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

type PublicDashboard struct {
	// Unique public dashboard identifier
	Uid string `json:"uid"`
	// Dashboard unique identifier referenced by this public dashboard
	DashboardUid string `json:"dashboardUid"`
	// Unique public access token
	AccessToken *string `json:"accessToken,omitempty"`
	// Flag that indicates if the public dashboard is enabled
	IsEnabled bool `json:"isEnabled"`
	// Flag that indicates if annotations are enabled
	AnnotationsEnabled bool `json:"annotationsEnabled"`
	// Flag that indicates if the time range picker is enabled
	TimeSelectionEnabled bool `json:"timeSelectionEnabled"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PublicDashboard` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *PublicDashboard) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "uid"
	if fields["uid"] != nil {
		if string(fields["uid"]) != "null" {
			if err := json.Unmarshal(fields["uid"], &resource.Uid); err != nil {
				errs = append(errs, cog.MakeBuildErrors("uid", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("uid", errors.New("required field is null"))...)

		}
		delete(fields, "uid")
	} else {
		errs = append(errs, cog.MakeBuildErrors("uid", errors.New("required field is missing from input"))...)
	}
	// Field "dashboardUid"
	if fields["dashboardUid"] != nil {
		if string(fields["dashboardUid"]) != "null" {
			if err := json.Unmarshal(fields["dashboardUid"], &resource.DashboardUid); err != nil {
				errs = append(errs, cog.MakeBuildErrors("dashboardUid", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("dashboardUid", errors.New("required field is null"))...)

		}
		delete(fields, "dashboardUid")
	} else {
		errs = append(errs, cog.MakeBuildErrors("dashboardUid", errors.New("required field is missing from input"))...)
	}
	// Field "accessToken"
	if fields["accessToken"] != nil {
		if string(fields["accessToken"]) != "null" {
			if err := json.Unmarshal(fields["accessToken"], &resource.AccessToken); err != nil {
				errs = append(errs, cog.MakeBuildErrors("accessToken", err)...)
			}

		}
		delete(fields, "accessToken")

	}
	// Field "isEnabled"
	if fields["isEnabled"] != nil {
		if string(fields["isEnabled"]) != "null" {
			if err := json.Unmarshal(fields["isEnabled"], &resource.IsEnabled); err != nil {
				errs = append(errs, cog.MakeBuildErrors("isEnabled", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("isEnabled", errors.New("required field is null"))...)

		}
		delete(fields, "isEnabled")
	} else {
		errs = append(errs, cog.MakeBuildErrors("isEnabled", errors.New("required field is missing from input"))...)
	}
	// Field "annotationsEnabled"
	if fields["annotationsEnabled"] != nil {
		if string(fields["annotationsEnabled"]) != "null" {
			if err := json.Unmarshal(fields["annotationsEnabled"], &resource.AnnotationsEnabled); err != nil {
				errs = append(errs, cog.MakeBuildErrors("annotationsEnabled", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("annotationsEnabled", errors.New("required field is null"))...)

		}
		delete(fields, "annotationsEnabled")
	} else {
		errs = append(errs, cog.MakeBuildErrors("annotationsEnabled", errors.New("required field is missing from input"))...)
	}
	// Field "timeSelectionEnabled"
	if fields["timeSelectionEnabled"] != nil {
		if string(fields["timeSelectionEnabled"]) != "null" {
			if err := json.Unmarshal(fields["timeSelectionEnabled"], &resource.TimeSelectionEnabled); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeSelectionEnabled", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("timeSelectionEnabled", errors.New("required field is null"))...)

		}
		delete(fields, "timeSelectionEnabled")
	} else {
		errs = append(errs, cog.MakeBuildErrors("timeSelectionEnabled", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("PublicDashboard", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `PublicDashboard` objects.
func (resource PublicDashboard) Equals(other PublicDashboard) bool {
	if resource.Uid != other.Uid {
		return false
	}
	if resource.DashboardUid != other.DashboardUid {
		return false
	}
	if resource.AccessToken == nil && other.AccessToken != nil || resource.AccessToken != nil && other.AccessToken == nil {
		return false
	}

	if resource.AccessToken != nil {
		if *resource.AccessToken != *other.AccessToken {
			return false
		}
	}
	if resource.IsEnabled != other.IsEnabled {
		return false
	}
	if resource.AnnotationsEnabled != other.AnnotationsEnabled {
		return false
	}
	if resource.TimeSelectionEnabled != other.TimeSelectionEnabled {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `PublicDashboard` fields for violations and returns them.
func (resource PublicDashboard) Validate() error {
	return nil
}
