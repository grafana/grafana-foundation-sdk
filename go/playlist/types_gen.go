// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package playlist

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

type PlaylistItemType string

const (
	PlaylistItemTypeDashboardByUid PlaylistItemType = "dashboard_by_uid"
	PlaylistItemTypeDashboardById  PlaylistItemType = "dashboard_by_id"
	PlaylistItemTypeDashboardByTag PlaylistItemType = "dashboard_by_tag"
)
