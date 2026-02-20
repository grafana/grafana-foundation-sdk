// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package playlistv0alpha1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Item] = (*ItemBuilder)(nil)

type ItemBuilder struct {
	internal *Item
	errors   cog.BuildErrors
}

func NewItemBuilder() *ItemBuilder {
	resource := NewItem()
	builder := &ItemBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ItemBuilder) Build() (Item, error) {
	if err := builder.internal.Validate(); err != nil {
		return Item{}, err
	}

	if len(builder.errors) > 0 {
		return Item{}, cog.MakeBuildErrors("playlistv0alpha1.item", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ItemBuilder) RecordError(path string, err error) *ItemBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// type of the item.
func (builder *ItemBuilder) Type(typeArg ItemType) *ItemBuilder {
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
func (builder *ItemBuilder) Value(value string) *ItemBuilder {
	builder.internal.Value = value

	return builder
}
