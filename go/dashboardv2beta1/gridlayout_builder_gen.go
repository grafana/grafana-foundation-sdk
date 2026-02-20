// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[GridLayoutKind] = (*GridLayoutBuilder)(nil)

type GridLayoutBuilder struct {
	internal *GridLayoutKind
	errors   cog.BuildErrors
}

func NewGridLayoutBuilder() *GridLayoutBuilder {
	resource := NewGridLayoutKind()
	builder := &GridLayoutBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "GridLayout"

	return builder
}

func (builder *GridLayoutBuilder) Build() (GridLayoutKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return GridLayoutKind{}, err
	}

	if len(builder.errors) > 0 {
		return GridLayoutKind{}, cog.MakeBuildErrors("dashboardv2beta1.gridLayout", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *GridLayoutBuilder) RecordError(path string, err error) *GridLayoutBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *GridLayoutBuilder) Items(items []cog.Builder[GridLayoutItemKind]) *GridLayoutBuilder {
	itemsResources := make([]GridLayoutItemKind, 0, len(items))
	for _, r1 := range items {
		itemsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		itemsResources = append(itemsResources, itemsDepth1)
	}
	builder.internal.Spec.Items = itemsResources

	return builder
}

func (builder *GridLayoutBuilder) Item(item cog.Builder[GridLayoutItemKind]) *GridLayoutBuilder {
	itemResource, err := item.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Items = append(builder.internal.Spec.Items, itemResource)

	return builder
}
