// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package accesspolicy

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

type AccessPolicy struct {
	// The scope where these policies should apply
	Scope ResourceRef `json:"scope"`
	// The role that must apply this policy
	Role RoleRef `json:"role"`
	// The set of rules to apply.  Note that * is required to modify
	// access policy rules, and that "none" will reject all actions
	Rules []AccessRule `json:"rules"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AccessPolicy` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AccessPolicy) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "scope"
	if fields["scope"] != nil {
		if string(fields["scope"]) != "null" {

			resource.Scope = ResourceRef{}
			if err := resource.Scope.UnmarshalJSONStrict(fields["scope"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("scope", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("scope", errors.New("required field is null"))...)

		}
		delete(fields, "scope")
	} else {
		errs = append(errs, cog.MakeBuildErrors("scope", errors.New("required field is missing from input"))...)
	}
	// Field "role"
	if fields["role"] != nil {
		if string(fields["role"]) != "null" {

			resource.Role = RoleRef{}
			if err := resource.Role.UnmarshalJSONStrict(fields["role"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("role", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("role", errors.New("required field is null"))...)

		}
		delete(fields, "role")
	} else {
		errs = append(errs, cog.MakeBuildErrors("role", errors.New("required field is missing from input"))...)
	}
	// Field "rules"
	if fields["rules"] != nil {
		if string(fields["rules"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["rules"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 AccessRule

				result1 = AccessRule{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("rules["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Rules = append(resource.Rules, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("rules", errors.New("required field is null"))...)

		}
		delete(fields, "rules")
	} else {
		errs = append(errs, cog.MakeBuildErrors("rules", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AccessPolicy", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AccessPolicy` objects.
func (resource AccessPolicy) Equals(other AccessPolicy) bool {
	if !resource.Scope.Equals(other.Scope) {
		return false
	}
	if !resource.Role.Equals(other.Role) {
		return false
	}

	if len(resource.Rules) != len(other.Rules) {
		return false
	}

	for i1 := range resource.Rules {
		if !resource.Rules[i1].Equals(other.Rules[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AccessPolicy` fields for violations and returns them.
func (resource AccessPolicy) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Scope.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("scope", err)...)
	}
	if err := resource.Role.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("role", err)...)
	}

	for i1 := range resource.Rules {
		if err := resource.Rules[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("rules["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type RoleRef struct {
	// Policies can apply to roles, teams, or users
	// Applying policies to individual users is supported, but discouraged
	Kind  RoleRefKind `json:"kind"`
	Name  string      `json:"name"`
	Xname string      `json:"xname"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RoleRef` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RoleRef) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
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
	// Field "xname"
	if fields["xname"] != nil {
		if string(fields["xname"]) != "null" {
			if err := json.Unmarshal(fields["xname"], &resource.Xname); err != nil {
				errs = append(errs, cog.MakeBuildErrors("xname", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("xname", errors.New("required field is null"))...)

		}
		delete(fields, "xname")
	} else {
		errs = append(errs, cog.MakeBuildErrors("xname", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RoleRef", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `RoleRef` objects.
func (resource RoleRef) Equals(other RoleRef) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Name != other.Name {
		return false
	}
	if resource.Xname != other.Xname {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `RoleRef` fields for violations and returns them.
func (resource RoleRef) Validate() error {
	return nil
}

type ResourceRef struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ResourceRef` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ResourceRef) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ResourceRef", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ResourceRef` objects.
func (resource ResourceRef) Equals(other ResourceRef) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Name != other.Name {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ResourceRef` fields for violations and returns them.
func (resource ResourceRef) Validate() error {
	return nil
}

type AccessRule struct {
	// The kind this rule applies to (dashboards, alert, etc)
	Kind string `json:"kind"`
	// READ, WRITE, CREATE, DELETE, ...
	// should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"
	Verb string `json:"verb"`
	// Specific sub-elements like "alert.rules" or "dashboard.permissions"????
	Target *string `json:"target,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AccessRule` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *AccessRule) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "kind"
	if fields["kind"] != nil {
		if string(fields["kind"]) != "null" {
			if err := json.Unmarshal(fields["kind"], &resource.Kind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("kind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is null"))...)

		}
		delete(fields, "kind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("kind", errors.New("required field is missing from input"))...)
	}
	// Field "verb"
	if fields["verb"] != nil {
		if string(fields["verb"]) != "null" {
			if err := json.Unmarshal(fields["verb"], &resource.Verb); err != nil {
				errs = append(errs, cog.MakeBuildErrors("verb", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("verb", errors.New("required field is null"))...)

		}
		delete(fields, "verb")
	} else {
		errs = append(errs, cog.MakeBuildErrors("verb", errors.New("required field is missing from input"))...)
	}
	// Field "target"
	if fields["target"] != nil {
		if string(fields["target"]) != "null" {
			if err := json.Unmarshal(fields["target"], &resource.Target); err != nil {
				errs = append(errs, cog.MakeBuildErrors("target", err)...)
			}

		}
		delete(fields, "target")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("AccessRule", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `AccessRule` objects.
func (resource AccessRule) Equals(other AccessRule) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Verb != other.Verb {
		return false
	}
	if resource.Target == nil && other.Target != nil || resource.Target != nil && other.Target == nil {
		return false
	}

	if resource.Target != nil {
		if *resource.Target != *other.Target {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `AccessRule` fields for violations and returns them.
func (resource AccessRule) Validate() error {
	return nil
}

type RoleRefKind string

const (
	RoleRefKindRole        RoleRefKind = "Role"
	RoleRefKindBuiltinRole RoleRefKind = "BuiltinRole"
	RoleRefKindTeam        RoleRefKind = "Team"
	RoleRefKindUser        RoleRefKind = "User"
)
