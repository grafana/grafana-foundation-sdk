// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardlist

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
