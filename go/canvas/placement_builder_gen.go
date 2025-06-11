// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Placement] = (*PlacementBuilder)(nil)

type PlacementBuilder struct {
	internal *Placement
	errors   cog.BuildErrors
}

func NewPlacementBuilder() *PlacementBuilder {
	resource := NewPlacement()
	builder := &PlacementBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *PlacementBuilder) Build() (Placement, error) {
	if err := builder.internal.Validate(); err != nil {
		return Placement{}, err
	}

	if len(builder.errors) > 0 {
		return Placement{}, cog.MakeBuildErrors("canvas.placement", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *PlacementBuilder) Top(top float64) *PlacementBuilder {
	builder.internal.Top = &top

	return builder
}

func (builder *PlacementBuilder) Left(left float64) *PlacementBuilder {
	builder.internal.Left = &left

	return builder
}

func (builder *PlacementBuilder) Right(right float64) *PlacementBuilder {
	builder.internal.Right = &right

	return builder
}

func (builder *PlacementBuilder) Bottom(bottom float64) *PlacementBuilder {
	builder.internal.Bottom = &bottom

	return builder
}

func (builder *PlacementBuilder) Width(width float64) *PlacementBuilder {
	builder.internal.Width = &width

	return builder
}

func (builder *PlacementBuilder) Height(height float64) *PlacementBuilder {
	builder.internal.Height = &height

	return builder
}

func (builder *PlacementBuilder) Rotation(rotation float64) *PlacementBuilder {
	builder.internal.Rotation = &rotation

	return builder
}
