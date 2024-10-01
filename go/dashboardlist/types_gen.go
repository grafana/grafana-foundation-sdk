// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardlist

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

type Options struct {
	KeepTime           bool     `json:"keepTime"`
	IncludeVars        bool     `json:"includeVars"`
	ShowStarred        bool     `json:"showStarred"`
	ShowRecentlyViewed bool     `json:"showRecentlyViewed"`
	ShowSearch         bool     `json:"showSearch"`
	ShowHeadings       bool     `json:"showHeadings"`
	MaxItems           int64    `json:"maxItems"`
	Query              string   `json:"query"`
	FolderId           *int64   `json:"folderId,omitempty"`
	Tags               []string `json:"tags"`
}

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
	if resource.MaxItems != other.MaxItems {
		return false
	}
	if resource.Query != other.Query {
		return false
	}
	if resource.FolderId == nil && other.FolderId != nil || resource.FolderId != nil && other.FolderId == nil {
		return false
	}

	if resource.FolderId != nil {
		if *resource.FolderId != *other.FolderId {
			return false
		}
	}

	if len(resource.Tags) != len(other.Tags) {
		return false
	}

	for i1 := range resource.Tags {
		if resource.Tags[i1] != other.Tags[i1] {
			return false
		}
	}

	return true
}

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
	}
}
