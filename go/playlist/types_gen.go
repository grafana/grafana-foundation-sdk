// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package playlist

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

type Playlist struct {
	// Unique playlist identifier. Generated on creation, either by the
	// creator of the playlist of by the application.
	Uid string `json:"uid"`
	// Name of the playlist.
	Name string `json:"name"`
	// Interval sets the time between switching views in a playlist.
	// FIXME: Is this based on a standardized format or what options are available? Can datemath be used?
	Interval string `json:"interval"`
	// The ordered list of items that the playlist will iterate over.
	// FIXME! This should not be optional, but changing it makes the godegen awkward
	Items []PlaylistItem `json:"items,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Playlist` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Playlist) UnmarshalJSONStrict(raw []byte) error {
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
	// Field "interval"
	if fields["interval"] != nil {
		if string(fields["interval"]) != "null" {
			if err := json.Unmarshal(fields["interval"], &resource.Interval); err != nil {
				errs = append(errs, cog.MakeBuildErrors("interval", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("interval", errors.New("required field is null"))...)

		}
		delete(fields, "interval")
	} else {
		errs = append(errs, cog.MakeBuildErrors("interval", errors.New("required field is missing from input"))...)
	}
	// Field "items"
	if fields["items"] != nil {
		if string(fields["items"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["items"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 PlaylistItem

				result1 = PlaylistItem{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("items["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Items = append(resource.Items, result1)
			}

		}
		delete(fields, "items")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Playlist", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Playlist` objects.
func (resource Playlist) Equals(other Playlist) bool {
	if resource.Uid != other.Uid {
		return false
	}
	if resource.Name != other.Name {
		return false
	}
	if resource.Interval != other.Interval {
		return false
	}

	if len(resource.Items) != len(other.Items) {
		return false
	}

	for i1 := range resource.Items {
		if !resource.Items[i1].Equals(other.Items[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Playlist` fields for violations and returns them.
func (resource Playlist) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Items {
		if err := resource.Items[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("items["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type PlaylistItem struct {
	// Type of the item.
	Type PlaylistItemType `json:"type"`
	// Value depends on type and describes the playlist item.
	//
	//  - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This
	//  is not portable as the numerical identifier is non-deterministic between different instances.
	//  Will be replaced by dashboard_by_uid in the future. (deprecated)
	//  - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All
	//  dashboards behind the tag will be added to the playlist.
	//  - dashboard_by_uid: The value is the dashboard UID
	Value string `json:"value"`
	// Title is an unused property -- it will be removed in the future
	Title *string `json:"title,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PlaylistItem` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *PlaylistItem) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {
			if err := json.Unmarshal(fields["value"], &resource.Value); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is null"))...)

		}
		delete(fields, "value")
	} else {
		errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is missing from input"))...)
	}
	// Field "title"
	if fields["title"] != nil {
		if string(fields["title"]) != "null" {
			if err := json.Unmarshal(fields["title"], &resource.Title); err != nil {
				errs = append(errs, cog.MakeBuildErrors("title", err)...)
			}

		}
		delete(fields, "title")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("PlaylistItem", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `PlaylistItem` objects.
func (resource PlaylistItem) Equals(other PlaylistItem) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Value != other.Value {
		return false
	}
	if resource.Title == nil && other.Title != nil || resource.Title != nil && other.Title == nil {
		return false
	}

	if resource.Title != nil {
		if *resource.Title != *other.Title {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `PlaylistItem` fields for violations and returns them.
func (resource PlaylistItem) Validate() error {
	return nil
}

type PlaylistItemType string

const (
	PlaylistItemTypeDashboardByUid PlaylistItemType = "dashboard_by_uid"
	PlaylistItemTypeDashboardById  PlaylistItemType = "dashboard_by_id"
	PlaylistItemTypeDashboardByTag PlaylistItemType = "dashboard_by_tag"
)
