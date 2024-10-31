// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package rolebinding

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

type RoleBinding struct {
	// The role we are discussing
	Role BuiltinRoleRefOrCustomRoleRef `json:"role"`
	// The team or user that has the specified role
	Subject RoleBindingSubject `json:"subject"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RoleBinding` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RoleBinding) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "role"
	if fields["role"] != nil {
		if string(fields["role"]) != "null" {

			resource.Role = BuiltinRoleRefOrCustomRoleRef{}
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
	// Field "subject"
	if fields["subject"] != nil {
		if string(fields["subject"]) != "null" {

			resource.Subject = RoleBindingSubject{}
			if err := resource.Subject.UnmarshalJSONStrict(fields["subject"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("subject", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("subject", errors.New("required field is null"))...)

		}
		delete(fields, "subject")
	} else {
		errs = append(errs, cog.MakeBuildErrors("subject", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RoleBinding", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `RoleBinding` objects.
func (resource RoleBinding) Equals(other RoleBinding) bool {
	if !resource.Role.Equals(other.Role) {
		return false
	}
	if !resource.Subject.Equals(other.Subject) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `RoleBinding` fields for violations and returns them.
func (resource RoleBinding) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Role.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("role", err)...)
	}
	if err := resource.Subject.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("subject", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type CustomRoleRef struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CustomRoleRef` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CustomRoleRef) UnmarshalJSONStrict(raw []byte) error {
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
		errs = append(errs, cog.MakeBuildErrors("CustomRoleRef", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `CustomRoleRef` objects.
func (resource CustomRoleRef) Equals(other CustomRoleRef) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Name != other.Name {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `CustomRoleRef` fields for violations and returns them.
func (resource CustomRoleRef) Validate() error {
	return nil
}

type BuiltinRoleRef struct {
	Kind string             `json:"kind"`
	Name BuiltinRoleRefName `json:"name"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BuiltinRoleRef` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *BuiltinRoleRef) UnmarshalJSONStrict(raw []byte) error {
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
		errs = append(errs, cog.MakeBuildErrors("BuiltinRoleRef", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `BuiltinRoleRef` objects.
func (resource BuiltinRoleRef) Equals(other BuiltinRoleRef) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Name != other.Name {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `BuiltinRoleRef` fields for violations and returns them.
func (resource BuiltinRoleRef) Validate() error {
	return nil
}

type RoleBindingSubject struct {
	Kind RoleBindingSubjectKind `json:"kind"`
	// The team/user identifier name
	Name string `json:"name"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RoleBindingSubject` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RoleBindingSubject) UnmarshalJSONStrict(raw []byte) error {
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
		errs = append(errs, cog.MakeBuildErrors("RoleBindingSubject", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `RoleBindingSubject` objects.
func (resource RoleBindingSubject) Equals(other RoleBindingSubject) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Name != other.Name {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `RoleBindingSubject` fields for violations and returns them.
func (resource RoleBindingSubject) Validate() error {
	return nil
}

type BuiltinRoleRefName string

const (
	BuiltinRoleRefNameViewer BuiltinRoleRefName = "viewer"
	BuiltinRoleRefNameEditor BuiltinRoleRefName = "editor"
	BuiltinRoleRefNameAdmin  BuiltinRoleRefName = "admin"
)

type RoleBindingSubjectKind string

const (
	RoleBindingSubjectKindTeam RoleBindingSubjectKind = "Team"
	RoleBindingSubjectKindUser RoleBindingSubjectKind = "User"
)

type BuiltinRoleRefOrCustomRoleRef struct {
	BuiltinRoleRef *BuiltinRoleRef `json:"BuiltinRoleRef,omitempty"`
	CustomRoleRef  *CustomRoleRef  `json:"CustomRoleRef,omitempty"`
}

// MarshalJSON implements a custom JSON marshalling logic to encode `BuiltinRoleRefOrCustomRoleRef` as JSON.
func (resource BuiltinRoleRefOrCustomRoleRef) MarshalJSON() ([]byte, error) {
	if resource.BuiltinRoleRef != nil {
		return json.Marshal(resource.BuiltinRoleRef)
	}
	if resource.CustomRoleRef != nil {
		return json.Marshal(resource.CustomRoleRef)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `BuiltinRoleRefOrCustomRoleRef` from JSON.
func (resource *BuiltinRoleRefOrCustomRoleRef) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["kind"]
	if !found {
		return errors.New("discriminator field 'kind' not found in payload")
	}

	switch discriminator {
	case "BuiltinRole":
		var builtinRoleRef BuiltinRoleRef
		if err := json.Unmarshal(raw, &builtinRoleRef); err != nil {
			return err
		}

		resource.BuiltinRoleRef = &builtinRoleRef
		return nil
	case "Role":
		var customRoleRef CustomRoleRef
		if err := json.Unmarshal(raw, &customRoleRef); err != nil {
			return err
		}

		resource.CustomRoleRef = &customRoleRef
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `kind = %v`", discriminator)
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BuiltinRoleRefOrCustomRoleRef` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *BuiltinRoleRefOrCustomRoleRef) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["kind"]
	if !found {
		return fmt.Errorf("discriminator field 'kind' not found in payload")
	}

	switch discriminator {
	case "BuiltinRole":
		builtinRoleRef := &BuiltinRoleRef{}
		if err := builtinRoleRef.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.BuiltinRoleRef = builtinRoleRef
		return nil
	case "Role":
		customRoleRef := &CustomRoleRef{}
		if err := customRoleRef.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.CustomRoleRef = customRoleRef
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `kind = %v`", discriminator)
}

// Equals tests the equality of two `BuiltinRoleRefOrCustomRoleRef` objects.
func (resource BuiltinRoleRefOrCustomRoleRef) Equals(other BuiltinRoleRefOrCustomRoleRef) bool {
	if resource.BuiltinRoleRef == nil && other.BuiltinRoleRef != nil || resource.BuiltinRoleRef != nil && other.BuiltinRoleRef == nil {
		return false
	}

	if resource.BuiltinRoleRef != nil {
		if !resource.BuiltinRoleRef.Equals(*other.BuiltinRoleRef) {
			return false
		}
	}
	if resource.CustomRoleRef == nil && other.CustomRoleRef != nil || resource.CustomRoleRef != nil && other.CustomRoleRef == nil {
		return false
	}

	if resource.CustomRoleRef != nil {
		if !resource.CustomRoleRef.Equals(*other.CustomRoleRef) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `BuiltinRoleRefOrCustomRoleRef` fields for violations and returns them.
func (resource BuiltinRoleRefOrCustomRoleRef) Validate() error {
	var errs cog.BuildErrors
	if resource.BuiltinRoleRef != nil {
		if err := resource.BuiltinRoleRef.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("BuiltinRoleRef", err)...)
		}
	}
	if resource.CustomRoleRef != nil {
		if err := resource.CustomRoleRef.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("CustomRoleRef", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}
