// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package folder

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

// TODO:
// common metadata will soon support setting the parent folder in the metadata
type Folder struct {
	// Unique folder id. (will be k8s name)
	Uid string `json:"uid"`
	// Folder title
	Title string `json:"title"`
	// Description of the folder.
	Description *string `json:"description,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Folder` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Folder) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "title"
	if fields["title"] != nil {
		if string(fields["title"]) != "null" {
			if err := json.Unmarshal(fields["title"], &resource.Title); err != nil {
				errs = append(errs, cog.MakeBuildErrors("title", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is null"))...)

		}
		delete(fields, "title")
	} else {
		errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is missing from input"))...)
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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Folder", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Folder` objects.
func (resource Folder) Equals(other Folder) bool {
	if resource.Uid != other.Uid {
		return false
	}
	if resource.Title != other.Title {
		return false
	}
	if resource.Description == nil && other.Description != nil || resource.Description != nil && other.Description == nil {
		return false
	}

	if resource.Description != nil {
		if *resource.Description != *other.Description {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Folder` fields for violations and returns them.
func (resource Folder) Validate() error {
	return nil
}
