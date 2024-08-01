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
	ShowFolderNames    bool     `json:"showFolderNames"`
	MaxItems           int64    `json:"maxItems"`
	Query              string   `json:"query"`
	Tags               []string `json:"tags"`
	// folderId is deprecated, and migrated to folderUid on panel init
	FolderId  *int64  `json:"folderId,omitempty"`
	FolderUID *string `json:"folderUID,omitempty"`
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "dashlist",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := Options{}

			if err := json.Unmarshal(raw, &options); err != nil {
				return nil, err
			}

			return options, nil
		},
	}
}
