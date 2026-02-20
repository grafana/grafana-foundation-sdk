// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package starsv1alpha1

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

type Resource struct {
	Group string `json:"group"`
	Kind  string `json:"kind"`
	// The set of resources
	// +listType=set
	Names []string `json:"names"`
}

// NewResource creates a new Resource object.
func NewResource() *Resource {
	return &Resource{
		Names: []string{},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Resource` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Resource) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "group"
	if fields["group"] != nil {
		if string(fields["group"]) != "null" {
			if err := json.Unmarshal(fields["group"], &resource.Group); err != nil {
				errs = append(errs, cog.MakeBuildErrors("group", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("group", errors.New("required field is null"))...)

		}
		delete(fields, "group")
	} else {
		errs = append(errs, cog.MakeBuildErrors("group", errors.New("required field is missing from input"))...)
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
	// Field "names"
	if fields["names"] != nil {
		if string(fields["names"]) != "null" {

			if err := json.Unmarshal(fields["names"], &resource.Names); err != nil {
				errs = append(errs, cog.MakeBuildErrors("names", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("names", errors.New("required field is null"))...)

		}
		delete(fields, "names")
	} else {
		errs = append(errs, cog.MakeBuildErrors("names", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Resource", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Resource` objects.
func (resource Resource) Equals(other Resource) bool {
	if resource.Group != other.Group {
		return false
	}
	if resource.Kind != other.Kind {
		return false
	}

	if len(resource.Names) != len(other.Names) {
		return false
	}

	for i1 := range resource.Names {
		if resource.Names[i1] != other.Names[i1] {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Resource` fields for violations and returns them.
func (resource Resource) Validate() error {
	return nil
}

type Stars struct {
	Resource []Resource `json:"resource"`
}

// NewStars creates a new Stars object.
func NewStars() *Stars {
	return &Stars{
		Resource: []Resource{},
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Stars` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Stars) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "resource"
	if fields["resource"] != nil {
		if string(fields["resource"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["resource"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 Resource

				result1 = Resource{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("resource["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Resource = append(resource.Resource, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("resource", errors.New("required field is null"))...)

		}
		delete(fields, "resource")
	} else {
		errs = append(errs, cog.MakeBuildErrors("resource", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Stars", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Stars` objects.
func (resource Stars) Equals(other Stars) bool {

	if len(resource.Resource) != len(other.Resource) {
		return false
	}

	for i1 := range resource.Resource {
		if !resource.Resource[i1].Equals(other.Resource[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Stars` fields for violations and returns them.
func (resource Stars) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Resource {
		if err := resource.Resource[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("resource["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

const StarsKind = "Stars"
