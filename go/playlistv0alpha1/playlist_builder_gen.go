// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package playlistv0alpha1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	resource "github.com/grafana/grafana-foundation-sdk/go/resource"
)

var _ cog.Builder[Playlist] = (*PlaylistBuilder)(nil)

type PlaylistBuilder struct {
	internal *Playlist
	errors   cog.BuildErrors
}

func NewPlaylistBuilder(title string) *PlaylistBuilder {
	resource := NewPlaylist()
	builder := &PlaylistBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Title = title

	return builder
}

// Creates a resource manifest from a Playlist.
func Manifest(name string, playlist cog.Builder[Playlist]) *resource.ManifestBuilder {
	builder := resource.NewManifestBuilder()

	builder.ApiVersion("playlist.grafana.app/playlistv0alpha1")

	builder.Kind("Playlist")

	builder.Metadata(resource.Named(name))
	playlistResource, err := playlist.Build()
	if err != nil {
		builder.RecordError("Manifest(playlist)", err)
		return builder
	}
	builder.Spec(playlistResource)

	return builder
}

func (builder *PlaylistBuilder) Build() (Playlist, error) {
	if err := builder.internal.Validate(); err != nil {
		return Playlist{}, err
	}

	if len(builder.errors) > 0 {
		return Playlist{}, cog.MakeBuildErrors("playlistv0alpha1.playlist", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *PlaylistBuilder) RecordError(path string, err error) *PlaylistBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *PlaylistBuilder) Title(title string) *PlaylistBuilder {
	builder.internal.Title = title

	return builder
}

func (builder *PlaylistBuilder) Interval(interval string) *PlaylistBuilder {
	builder.internal.Interval = interval

	return builder
}

func (builder *PlaylistBuilder) Items(items []cog.Builder[Item]) *PlaylistBuilder {
	itemsResources := make([]Item, 0, len(items))
	for _, r1 := range items {
		itemsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		itemsResources = append(itemsResources, itemsDepth1)
	}
	builder.internal.Items = itemsResources

	return builder
}

func (builder *PlaylistBuilder) Item(item cog.Builder[Item]) *PlaylistBuilder {
	itemResource, err := item.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Items = append(builder.internal.Items, itemResource)

	return builder
}
