// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package playlist

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Playlist] = (*PlaylistBuilder)(nil)

type PlaylistBuilder struct {
	internal *Playlist
	errors   map[string]cog.BuildErrors
}

func NewPlaylistBuilder() *PlaylistBuilder {
	resource := &Playlist{}
	builder := &PlaylistBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *PlaylistBuilder) Build() (Playlist, error) {
	if err := builder.internal.Validate(); err != nil {
		return Playlist{}, err
	}

	return *builder.internal, nil
}

// Unique playlist identifier. Generated on creation, either by the
// creator of the playlist of by the application.
func (builder *PlaylistBuilder) Uid(uid string) *PlaylistBuilder {
	builder.internal.Uid = uid

	return builder
}

// Name of the playlist.
func (builder *PlaylistBuilder) Name(name string) *PlaylistBuilder {
	builder.internal.Name = name

	return builder
}

// Interval sets the time between switching views in a playlist.
// FIXME: Is this based on a standardized format or what options are available? Can datemath be used?
func (builder *PlaylistBuilder) Interval(interval string) *PlaylistBuilder {
	builder.internal.Interval = interval

	return builder
}

// The ordered list of items that the playlist will iterate over.
// FIXME! This should not be optional, but changing it makes the godegen awkward
func (builder *PlaylistBuilder) Items(items []cog.Builder[PlaylistItem]) *PlaylistBuilder {
	itemsResources := make([]PlaylistItem, 0, len(items))
	for _, r1 := range items {
		itemsDepth1, err := r1.Build()
		if err != nil {
			builder.errors["items"] = err.(cog.BuildErrors)
			return builder
		}
		itemsResources = append(itemsResources, itemsDepth1)
	}
	builder.internal.Items = itemsResources

	return builder
}

func (builder *PlaylistBuilder) applyDefaults() {
	builder.Interval("5m")
}
