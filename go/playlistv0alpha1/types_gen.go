// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package playlistv0alpha1

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

type Item struct {
	// type of the item.
	Type ItemType `json:"type"`
	// Value depends on type and describes the playlist item.
	//  - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This
	//  is not portable as the numerical identifier is non-deterministic between different instances.
	//  Will be replaced by dashboard_by_uid in the future. (deprecated)
	//  - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All
	//  dashboards behind the tag will be added to the playlist.
	//  - dashboard_by_uid: The value is the dashboard UID
	Value string `json:"value"`
}

// NewItem creates a new Item object.
func NewItem() *Item {
	return &Item{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Item` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Item) UnmarshalJSONStrict(raw []byte) error {
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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Item", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Item` objects.
func (resource Item) Equals(other Item) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Value != other.Value {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Item` fields for violations and returns them.
func (resource Item) Validate() error {
	return nil
}

type Playlist struct {
	Title    string `json:"title"`
	Interval string `json:"interval"`
	Items    []Item `json:"items"`
}

// NewPlaylist creates a new Playlist object.
func NewPlaylist() *Playlist {
	return &Playlist{
		Items: []Item{},
	}
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
				var result1 Item

				result1 = Item{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("items["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Items = append(resource.Items, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("items", errors.New("required field is null"))...)

		}
		delete(fields, "items")
	} else {
		errs = append(errs, cog.MakeBuildErrors("items", errors.New("required field is missing from input"))...)
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
	if resource.Title != other.Title {
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

type ItemType string

const (
	ItemTypeDashboardByTag ItemType = "dashboard_by_tag"
	ItemTypeDashboardByUid ItemType = "dashboard_by_uid"
	ItemTypeDashboardById  ItemType = "dashboard_by_id"
)
