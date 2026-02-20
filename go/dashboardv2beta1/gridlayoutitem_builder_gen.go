// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[GridLayoutItemKind] = (*GridLayoutItemBuilder)(nil)

type GridLayoutItemBuilder struct {
	internal *GridLayoutItemKind
	errors   cog.BuildErrors
}

func NewGridLayoutItemBuilder(name string) *GridLayoutItemBuilder {
	resource := NewGridLayoutItemKind()
	builder := &GridLayoutItemBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "GridLayoutItem"
	builder.internal.Spec.Element.Kind = "ElementReference"
	builder.internal.Spec.Element.Name = name

	return builder
}

func (builder *GridLayoutItemBuilder) Build() (GridLayoutItemKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return GridLayoutItemKind{}, err
	}

	if len(builder.errors) > 0 {
		return GridLayoutItemKind{}, cog.MakeBuildErrors("dashboardv2beta1.gridLayoutItem", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *GridLayoutItemBuilder) RecordError(path string, err error) *GridLayoutItemBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *GridLayoutItemBuilder) X(x int64) *GridLayoutItemBuilder {
	builder.internal.Spec.X = x

	return builder
}

func (builder *GridLayoutItemBuilder) Y(y int64) *GridLayoutItemBuilder {
	builder.internal.Spec.Y = y

	return builder
}

func (builder *GridLayoutItemBuilder) Width(width int64) *GridLayoutItemBuilder {
	builder.internal.Spec.Width = width

	return builder
}

func (builder *GridLayoutItemBuilder) Height(height int64) *GridLayoutItemBuilder {
	builder.internal.Spec.Height = height

	return builder
}

func (builder *GridLayoutItemBuilder) Repeat(repeat cog.Builder[RepeatOptions]) *GridLayoutItemBuilder {
	repeatResource, err := repeat.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Repeat = &repeatResource

	return builder
}

func (builder *GridLayoutItemBuilder) Element(name string) *GridLayoutItemBuilder {
	builder.internal.Spec.Element.Name = name

	return builder
}
