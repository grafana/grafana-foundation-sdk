// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

type Manifest struct {
	ApiVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
	Spec       any      `json:"spec"`
}

// NewManifest creates a new Manifest object.
func NewManifest() *Manifest {
	return &Manifest{
		Metadata: *NewMetadata(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Manifest` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Manifest) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "apiVersion"
	if fields["apiVersion"] != nil {
		if string(fields["apiVersion"]) != "null" {
			if err := json.Unmarshal(fields["apiVersion"], &resource.ApiVersion); err != nil {
				errs = append(errs, cog.MakeBuildErrors("apiVersion", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("apiVersion", errors.New("required field is null"))...)

		}
		delete(fields, "apiVersion")
	} else {
		errs = append(errs, cog.MakeBuildErrors("apiVersion", errors.New("required field is missing from input"))...)
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
	// Field "metadata"
	if fields["metadata"] != nil {
		if string(fields["metadata"]) != "null" {

			resource.Metadata = Metadata{}
			if err := resource.Metadata.UnmarshalJSONStrict(fields["metadata"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metadata", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("metadata", errors.New("required field is null"))...)

		}
		delete(fields, "metadata")
	} else {
		errs = append(errs, cog.MakeBuildErrors("metadata", errors.New("required field is missing from input"))...)
	}
	// Field "spec"
	if fields["spec"] != nil {
		if string(fields["spec"]) != "null" {
			if err := json.Unmarshal(fields["spec"], &resource.Spec); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spec", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is null"))...)

		}
		delete(fields, "spec")
	} else {
		errs = append(errs, cog.MakeBuildErrors("spec", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Manifest", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Manifest` objects.
func (resource Manifest) Equals(other Manifest) bool {
	if resource.ApiVersion != other.ApiVersion {
		return false
	}
	if resource.Kind != other.Kind {
		return false
	}
	if !resource.Metadata.Equals(other.Metadata) {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Spec, other.Spec) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Manifest` fields for violations and returns them.
func (resource Manifest) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Metadata.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("metadata", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Metadata struct {
	Name              string            `json:"name"`
	Namespace         *string           `json:"namespace,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	Annotations       map[string]string `json:"annotations,omitempty"`
	Uid               *string           `json:"uid,omitempty"`
	ResourceVersion   *string           `json:"resourceVersion,omitempty"`
	Generation        *int64            `json:"generation,omitempty"`
	CreationTimestamp *string           `json:"creationTimestamp,omitempty"`
	UpdateTimestamp   *string           `json:"updateTimestamp,omitempty"`
	DeletionTimestamp *string           `json:"deletionTimestamp,omitempty"`
}

// NewMetadata creates a new Metadata object.
func NewMetadata() *Metadata {
	return &Metadata{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Metadata` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Metadata) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "namespace"
	if fields["namespace"] != nil {
		if string(fields["namespace"]) != "null" {
			if err := json.Unmarshal(fields["namespace"], &resource.Namespace); err != nil {
				errs = append(errs, cog.MakeBuildErrors("namespace", err)...)
			}

		}
		delete(fields, "namespace")

	}
	// Field "labels"
	if fields["labels"] != nil {
		if string(fields["labels"]) != "null" {

			if err := json.Unmarshal(fields["labels"], &resource.Labels); err != nil {
				errs = append(errs, cog.MakeBuildErrors("labels", err)...)
			}

		}
		delete(fields, "labels")

	}
	// Field "annotations"
	if fields["annotations"] != nil {
		if string(fields["annotations"]) != "null" {

			if err := json.Unmarshal(fields["annotations"], &resource.Annotations); err != nil {
				errs = append(errs, cog.MakeBuildErrors("annotations", err)...)
			}

		}
		delete(fields, "annotations")

	}
	// Field "uid"
	if fields["uid"] != nil {
		if string(fields["uid"]) != "null" {
			if err := json.Unmarshal(fields["uid"], &resource.Uid); err != nil {
				errs = append(errs, cog.MakeBuildErrors("uid", err)...)
			}

		}
		delete(fields, "uid")

	}
	// Field "resourceVersion"
	if fields["resourceVersion"] != nil {
		if string(fields["resourceVersion"]) != "null" {
			if err := json.Unmarshal(fields["resourceVersion"], &resource.ResourceVersion); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resourceVersion", err)...)
			}

		}
		delete(fields, "resourceVersion")

	}
	// Field "generation"
	if fields["generation"] != nil {
		if string(fields["generation"]) != "null" {
			if err := json.Unmarshal(fields["generation"], &resource.Generation); err != nil {
				errs = append(errs, cog.MakeBuildErrors("generation", err)...)
			}

		}
		delete(fields, "generation")

	}
	// Field "creationTimestamp"
	if fields["creationTimestamp"] != nil {
		if string(fields["creationTimestamp"]) != "null" {
			if err := json.Unmarshal(fields["creationTimestamp"], &resource.CreationTimestamp); err != nil {
				errs = append(errs, cog.MakeBuildErrors("creationTimestamp", err)...)
			}

		}
		delete(fields, "creationTimestamp")

	}
	// Field "updateTimestamp"
	if fields["updateTimestamp"] != nil {
		if string(fields["updateTimestamp"]) != "null" {
			if err := json.Unmarshal(fields["updateTimestamp"], &resource.UpdateTimestamp); err != nil {
				errs = append(errs, cog.MakeBuildErrors("updateTimestamp", err)...)
			}

		}
		delete(fields, "updateTimestamp")

	}
	// Field "deletionTimestamp"
	if fields["deletionTimestamp"] != nil {
		if string(fields["deletionTimestamp"]) != "null" {
			if err := json.Unmarshal(fields["deletionTimestamp"], &resource.DeletionTimestamp); err != nil {
				errs = append(errs, cog.MakeBuildErrors("deletionTimestamp", err)...)
			}

		}
		delete(fields, "deletionTimestamp")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Metadata", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Metadata` objects.
func (resource Metadata) Equals(other Metadata) bool {
	if resource.Name != other.Name {
		return false
	}
	if resource.Namespace == nil && other.Namespace != nil || resource.Namespace != nil && other.Namespace == nil {
		return false
	}

	if resource.Namespace != nil {
		if *resource.Namespace != *other.Namespace {
			return false
		}
	}

	if len(resource.Labels) != len(other.Labels) {
		return false
	}

	for key1 := range resource.Labels {
		if resource.Labels[key1] != other.Labels[key1] {
			return false
		}
	}

	if len(resource.Annotations) != len(other.Annotations) {
		return false
	}

	for key1 := range resource.Annotations {
		if resource.Annotations[key1] != other.Annotations[key1] {
			return false
		}
	}
	if resource.Uid == nil && other.Uid != nil || resource.Uid != nil && other.Uid == nil {
		return false
	}

	if resource.Uid != nil {
		if *resource.Uid != *other.Uid {
			return false
		}
	}
	if resource.ResourceVersion == nil && other.ResourceVersion != nil || resource.ResourceVersion != nil && other.ResourceVersion == nil {
		return false
	}

	if resource.ResourceVersion != nil {
		if *resource.ResourceVersion != *other.ResourceVersion {
			return false
		}
	}
	if resource.Generation == nil && other.Generation != nil || resource.Generation != nil && other.Generation == nil {
		return false
	}

	if resource.Generation != nil {
		if *resource.Generation != *other.Generation {
			return false
		}
	}
	if resource.CreationTimestamp == nil && other.CreationTimestamp != nil || resource.CreationTimestamp != nil && other.CreationTimestamp == nil {
		return false
	}

	if resource.CreationTimestamp != nil {
		if *resource.CreationTimestamp != *other.CreationTimestamp {
			return false
		}
	}
	if resource.UpdateTimestamp == nil && other.UpdateTimestamp != nil || resource.UpdateTimestamp != nil && other.UpdateTimestamp == nil {
		return false
	}

	if resource.UpdateTimestamp != nil {
		if *resource.UpdateTimestamp != *other.UpdateTimestamp {
			return false
		}
	}
	if resource.DeletionTimestamp == nil && other.DeletionTimestamp != nil || resource.DeletionTimestamp != nil && other.DeletionTimestamp == nil {
		return false
	}

	if resource.DeletionTimestamp != nil {
		if *resource.DeletionTimestamp != *other.DeletionTimestamp {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Metadata` fields for violations and returns them.
func (resource Metadata) Validate() error {
	return nil
}
