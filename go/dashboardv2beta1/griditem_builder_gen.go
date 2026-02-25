// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[GridLayoutItemKind] = (*GridItemBuilder)(nil)

type GridItemBuilder struct {
	internal *GridLayoutItemKind
	errors   cog.BuildErrors
}

func NewGridItemBuilder() *GridItemBuilder {
	resource := NewGridLayoutItemKind()
	builder := &GridItemBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "GridLayoutItem"
	builder.internal.Spec.Element.Kind = "ElementReference"

	return builder
}

func GridItem(name string) *GridItemBuilder {
	builder := NewGridItemBuilder()

	builder.Name(name)

	return builder
}

func (builder *GridItemBuilder) Build() (GridLayoutItemKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return GridLayoutItemKind{}, err
	}

	if len(builder.errors) > 0 {
		return GridLayoutItemKind{}, cog.MakeBuildErrors("dashboardv2beta1.gridItem", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *GridItemBuilder) RecordError(path string, err error) *GridItemBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *GridItemBuilder) X(x int64) *GridItemBuilder {
	builder.internal.Spec.X = x

	return builder
}

func (builder *GridItemBuilder) Y(y int64) *GridItemBuilder {
	builder.internal.Spec.Y = y

	return builder
}

func (builder *GridItemBuilder) Width(width int64) *GridItemBuilder {
	builder.internal.Spec.Width = width

	return builder
}

func (builder *GridItemBuilder) Height(height int64) *GridItemBuilder {
	builder.internal.Spec.Height = height

	return builder
}

func (builder *GridItemBuilder) Repeat(repeat cog.Builder[RepeatOptions]) *GridItemBuilder {
	repeatResource, err := repeat.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Repeat = &repeatResource

	return builder
}

func (builder *GridItemBuilder) Name(name string) *GridItemBuilder {
	builder.internal.Spec.Element.Name = name

	return builder
}
