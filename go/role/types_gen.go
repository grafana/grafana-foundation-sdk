// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package role

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

type Role struct {
	// The role identifier `managed:builtins:editor:permissions`
	Name string `json:"name"`
	// Optional display
	DisplayName *string `json:"displayName,omitempty"`
	// Name of the team.
	GroupName *string `json:"groupName,omitempty"`
	// Role description
	Description *string `json:"description,omitempty"`
	// Do not show this role
	Hidden bool `json:"hidden"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Role` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Role) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "displayName"
	if fields["displayName"] != nil {
		if string(fields["displayName"]) != "null" {
			if err := json.Unmarshal(fields["displayName"], &resource.DisplayName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("displayName", err)...)
			}

		}
		delete(fields, "displayName")

	}
	// Field "groupName"
	if fields["groupName"] != nil {
		if string(fields["groupName"]) != "null" {
			if err := json.Unmarshal(fields["groupName"], &resource.GroupName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("groupName", err)...)
			}

		}
		delete(fields, "groupName")

	}
	// Field "description"
	if fields["description"] != nil {
		if string(fields["description"]) != "null" {
			if err := json.Unmarshal(fields["description"], &resource.Description); err != nil {
				errs = append(errs, cog.MakeBuildErrors("description", err)...)
			}

		}
		delete(fields, "description")

	}
	// Field "hidden"
	if fields["hidden"] != nil {
		if string(fields["hidden"]) != "null" {
			if err := json.Unmarshal(fields["hidden"], &resource.Hidden); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hidden", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("hidden", errors.New("required field is null"))...)

		}
		delete(fields, "hidden")
	} else {
		errs = append(errs, cog.MakeBuildErrors("hidden", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Role", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Role` objects.
func (resource Role) Equals(other Role) bool {
	if resource.Name != other.Name {
		return false
	}
	if resource.DisplayName == nil && other.DisplayName != nil || resource.DisplayName != nil && other.DisplayName == nil {
		return false
	}

	if resource.DisplayName != nil {
		if *resource.DisplayName != *other.DisplayName {
			return false
		}
	}
	if resource.GroupName == nil && other.GroupName != nil || resource.GroupName != nil && other.GroupName == nil {
		return false
	}

	if resource.GroupName != nil {
		if *resource.GroupName != *other.GroupName {
			return false
		}
	}
	if resource.Description == nil && other.Description != nil || resource.Description != nil && other.Description == nil {
		return false
	}

	if resource.Description != nil {
		if *resource.Description != *other.Description {
			return false
		}
	}
	if resource.Hidden != other.Hidden {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Role` fields for violations and returns them.
func (resource Role) Validate() error {
	return nil
}
