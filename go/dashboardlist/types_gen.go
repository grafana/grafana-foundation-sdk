// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardlist

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type Options struct {
	KeepTime           bool     `json:"keepTime"`
	IncludeVars        bool     `json:"includeVars"`
	ShowStarred        bool     `json:"showStarred"`
	ShowRecentlyViewed bool     `json:"showRecentlyViewed"`
	ShowSearch         bool     `json:"showSearch"`
	ShowHeadings       bool     `json:"showHeadings"`
	ShowFolderNames    bool     `json:"showFolderNames"`
	MaxItems           int64    `json:"maxItems"`
	Query              string   `json:"query"`
	Tags               []string `json:"tags"`
	// folderId is deprecated, and migrated to folderUid on panel init
	FolderId  *int64  `json:"folderId,omitempty"`
	FolderUID *string `json:"folderUID,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Options` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Options) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "keepTime"
	if fields["keepTime"] != nil {
		if string(fields["keepTime"]) != "null" {
			if err := json.Unmarshal(fields["keepTime"], &resource.KeepTime); err != nil {
				errs = append(errs, cog.MakeBuildErrors("keepTime", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("keepTime", errors.New("required field is null"))...)

		}
		delete(fields, "keepTime")
	} else {
		errs = append(errs, cog.MakeBuildErrors("keepTime", errors.New("required field is missing from input"))...)
	}
	// Field "includeVars"
	if fields["includeVars"] != nil {
		if string(fields["includeVars"]) != "null" {
			if err := json.Unmarshal(fields["includeVars"], &resource.IncludeVars); err != nil {
				errs = append(errs, cog.MakeBuildErrors("includeVars", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("includeVars", errors.New("required field is null"))...)

		}
		delete(fields, "includeVars")
	} else {
		errs = append(errs, cog.MakeBuildErrors("includeVars", errors.New("required field is missing from input"))...)
	}
	// Field "showStarred"
	if fields["showStarred"] != nil {
		if string(fields["showStarred"]) != "null" {
			if err := json.Unmarshal(fields["showStarred"], &resource.ShowStarred); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showStarred", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showStarred", errors.New("required field is null"))...)

		}
		delete(fields, "showStarred")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showStarred", errors.New("required field is missing from input"))...)
	}
	// Field "showRecentlyViewed"
	if fields["showRecentlyViewed"] != nil {
		if string(fields["showRecentlyViewed"]) != "null" {
			if err := json.Unmarshal(fields["showRecentlyViewed"], &resource.ShowRecentlyViewed); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showRecentlyViewed", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showRecentlyViewed", errors.New("required field is null"))...)

		}
		delete(fields, "showRecentlyViewed")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showRecentlyViewed", errors.New("required field is missing from input"))...)
	}
	// Field "showSearch"
	if fields["showSearch"] != nil {
		if string(fields["showSearch"]) != "null" {
			if err := json.Unmarshal(fields["showSearch"], &resource.ShowSearch); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showSearch", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showSearch", errors.New("required field is null"))...)

		}
		delete(fields, "showSearch")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showSearch", errors.New("required field is missing from input"))...)
	}
	// Field "showHeadings"
	if fields["showHeadings"] != nil {
		if string(fields["showHeadings"]) != "null" {
			if err := json.Unmarshal(fields["showHeadings"], &resource.ShowHeadings); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showHeadings", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showHeadings", errors.New("required field is null"))...)

		}
		delete(fields, "showHeadings")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showHeadings", errors.New("required field is missing from input"))...)
	}
	// Field "showFolderNames"
	if fields["showFolderNames"] != nil {
		if string(fields["showFolderNames"]) != "null" {
			if err := json.Unmarshal(fields["showFolderNames"], &resource.ShowFolderNames); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showFolderNames", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showFolderNames", errors.New("required field is null"))...)

		}
		delete(fields, "showFolderNames")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showFolderNames", errors.New("required field is missing from input"))...)
	}
	// Field "maxItems"
	if fields["maxItems"] != nil {
		if string(fields["maxItems"]) != "null" {
			if err := json.Unmarshal(fields["maxItems"], &resource.MaxItems); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxItems", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("maxItems", errors.New("required field is null"))...)

		}
		delete(fields, "maxItems")
	} else {
		errs = append(errs, cog.MakeBuildErrors("maxItems", errors.New("required field is missing from input"))...)
	}
	// Field "query"
	if fields["query"] != nil {
		if string(fields["query"]) != "null" {
			if err := json.Unmarshal(fields["query"], &resource.Query); err != nil {
				errs = append(errs, cog.MakeBuildErrors("query", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is null"))...)

		}
		delete(fields, "query")
	} else {
		errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is missing from input"))...)
	}
	// Field "tags"
	if fields["tags"] != nil {
		if string(fields["tags"]) != "null" {

			if err := json.Unmarshal(fields["tags"], &resource.Tags); err != nil {
				errs = append(errs, cog.MakeBuildErrors("tags", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("tags", errors.New("required field is null"))...)

		}
		delete(fields, "tags")
	} else {
		errs = append(errs, cog.MakeBuildErrors("tags", errors.New("required field is missing from input"))...)
	}
	// Field "folderId"
	if fields["folderId"] != nil {
		if string(fields["folderId"]) != "null" {
			if err := json.Unmarshal(fields["folderId"], &resource.FolderId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("folderId", err)...)
			}

		}
		delete(fields, "folderId")

	}
	// Field "folderUID"
	if fields["folderUID"] != nil {
		if string(fields["folderUID"]) != "null" {
			if err := json.Unmarshal(fields["folderUID"], &resource.FolderUID); err != nil {
				errs = append(errs, cog.MakeBuildErrors("folderUID", err)...)
			}

		}
		delete(fields, "folderUID")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Options", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Options` objects.
func (resource Options) Equals(other Options) bool {
	if resource.KeepTime != other.KeepTime {
		return false
	}
	if resource.IncludeVars != other.IncludeVars {
		return false
	}
	if resource.ShowStarred != other.ShowStarred {
		return false
	}
	if resource.ShowRecentlyViewed != other.ShowRecentlyViewed {
		return false
	}
	if resource.ShowSearch != other.ShowSearch {
		return false
	}
	if resource.ShowHeadings != other.ShowHeadings {
		return false
	}
	if resource.ShowFolderNames != other.ShowFolderNames {
		return false
	}
	if resource.MaxItems != other.MaxItems {
		return false
	}
	if resource.Query != other.Query {
		return false
	}

	if len(resource.Tags) != len(other.Tags) {
		return false
	}

	for i1 := range resource.Tags {
		if resource.Tags[i1] != other.Tags[i1] {
			return false
		}
	}
	if resource.FolderId == nil && other.FolderId != nil || resource.FolderId != nil && other.FolderId == nil {
		return false
	}

	if resource.FolderId != nil {
		if *resource.FolderId != *other.FolderId {
			return false
		}
	}
	if resource.FolderUID == nil && other.FolderUID != nil || resource.FolderUID != nil && other.FolderUID == nil {
		return false
	}

	if resource.FolderUID != nil {
		if *resource.FolderUID != *other.FolderUID {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	return nil
}

// VariantConfig returns the configuration related to dashlist panels.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "dashlist",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := &Options{}

			if err := json.Unmarshal(raw, options); err != nil {
				return nil, err
			}

			return options, nil
		},
		StrictOptionsUnmarshaler: func(raw []byte) (any, error) {
			options := &Options{}

			if err := options.UnmarshalJSONStrict(raw); err != nil {
				return nil, err
			}

			return options, nil
		},
		GoConverter: func(inputPanel any) string {
			if panel, ok := inputPanel.(*dashboard.Panel); ok {
				return PanelConverter(*panel)
			}

			return PanelConverter(inputPanel.(dashboard.Panel))
		},
	}
}
