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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Team` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
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

// Equals tests the equality of two `Team` objects.
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

// Validate checks all the validation constraints that may be defined on `Team` fields for violations and returns them.
func (resource Team) Validate() error {
	return nil
}
