// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package team

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

type Team struct {
	// Name of the team.
	Name string `json:"name"`
	// Email of the team.
	Email *string `json:"email,omitempty"`
}

func (resource *Team) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("name", errors.New("required field is null"))...)

		}
		delete(fields, "name")
	} else {
		errs = append(errs, cog.MakeBuildErrors("name", errors.New("required field is missing from input"))...)
	}
	// Field "email"
	if fields["email"] != nil {
		if string(fields["email"]) != "null" {
			if err := json.Unmarshal(fields["email"], &resource.Email); err != nil {
				errs = append(errs, cog.MakeBuildErrors("email", err)...)
			}

		}
		delete(fields, "email")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Team", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

func (resource Team) Equals(other Team) bool {
	if resource.Name != other.Name {
		return false
	}
	if resource.Email == nil && other.Email != nil || resource.Email != nil && other.Email == nil {
		return false
	}

	if resource.Email != nil {
		if *resource.Email != *other.Email {
			return false
		}
	}

	return true
}

// Validate checks any constraint that may be defined for this type
// and returns all violations.
func (resource Team) Validate() error {
	return nil
}
