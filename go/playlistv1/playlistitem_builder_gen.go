// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package playlistv1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[PlaylistItem] = (*PlaylistItemBuilder)(nil)

// Shared item definition for all versions
type PlaylistItemBuilder struct {
	internal *PlaylistItem
	errors   cog.BuildErrors
}

func NewPlaylistItemBuilder() *PlaylistItemBuilder {
	resource := NewPlaylistItem()
	builder := &PlaylistItemBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *PlaylistItemBuilder) Build() (PlaylistItem, error) {
	if err := builder.internal.Validate(); err != nil {
		return PlaylistItem{}, err
	}

	if len(builder.errors) > 0 {
		return PlaylistItem{}, cog.MakeBuildErrors("playlistv1.playlistItem", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *PlaylistItemBuilder) RecordError(path string, err error) *PlaylistItemBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// type of the item.
func (builder *PlaylistItemBuilder) Type(typeArg PlaylistItemType) *PlaylistItemBuilder {
	builder.internal.Type = typeArg

	return builder
}

// Value depends on type and describes the playlist item.
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
