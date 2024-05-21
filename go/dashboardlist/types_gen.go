// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardlist

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
