// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package rolebinding

import (
	"encoding/json"
	"errors"
	"fmt"
)

type RoleBinding struct {
	// The role we are discussing
	Role BuiltinRoleRefOrCustomRoleRef `json:"role"`
	// The team or user that has the specified role
	Subject RoleBindingSubject `json:"subject"`
}

func (resource RoleBinding) Equals(other RoleBinding) bool {
	if !resource.Role.Equals(other.Role) {
		return false
	}
	if !resource.Subject.Equals(other.Subject) {
		return false
	}

	return true
}

type CustomRoleRef struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

func (resource CustomRoleRef) Equals(other CustomRoleRef) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Name != other.Name {
		return false
	}

	return true
}

type BuiltinRoleRef struct {
	Kind string             `json:"kind"`
	Name BuiltinRoleRefName `json:"name"`
}

func (resource BuiltinRoleRef) Equals(other BuiltinRoleRef) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Name != other.Name {
		return false
	}

	return true
}

type RoleBindingSubject struct {
	Kind RoleBindingSubjectKind `json:"kind"`
	// The team/user identifier name
	Name string `json:"name"`
}

func (resource RoleBindingSubject) Equals(other RoleBindingSubject) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Name != other.Name {
		return false
	}

	return true
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

func (resource BuiltinRoleRefOrCustomRoleRef) MarshalJSON() ([]byte, error) {
	if resource.BuiltinRoleRef != nil {
		return json.Marshal(resource.BuiltinRoleRef)
	}
	if resource.CustomRoleRef != nil {
		return json.Marshal(resource.CustomRoleRef)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

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
