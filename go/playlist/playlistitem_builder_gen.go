// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package playlist

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[PlaylistItem] = (*PlaylistItemBuilder)(nil)

type PlaylistItemBuilder struct {
	internal *PlaylistItem
	errors   map[string]cog.BuildErrors
}

func NewPlaylistItemBuilder() *PlaylistItemBuilder {
	resource := &PlaylistItem{}
	builder := &PlaylistItemBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *PlaylistItemBuilder) Build() (PlaylistItem, error) {
	if err := builder.internal.Validate(); err != nil {
		return PlaylistItem{}, err
	}

	return *builder.internal, nil
}

// Type of the item.
func (builder *PlaylistItemBuilder) Type(typeArg PlaylistItemType) *PlaylistItemBuilder {
	builder.internal.Type = typeArg

	return builder
}

// Value depends on type and describes the playlist item.
//
//   - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This
//     is not portable as the numerical identifier is non-deterministic between different instances.
//     Will be replaced by dashboard_by_uid in the future. (deprecated)
//   - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All
//     dashboards behind the tag will be added to the playlist.
//   - dashboard_by_uid: The value is the dashboard UID
func (builder *PlaylistItemBuilder) Value(value string) *PlaylistItemBuilder {
	builder.internal.Value = value

	return builder
}

// Title is an unused property -- it will be removed in the future
func (builder *PlaylistItemBuilder) Title(title string) *PlaylistItemBuilder {
	builder.internal.Title = &title

	return builder
}

func (builder *PlaylistItemBuilder) applyDefaults() {
}
