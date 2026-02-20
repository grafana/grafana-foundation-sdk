// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Constraint] = (*ConstraintBuilder)(nil)

type ConstraintBuilder struct {
	internal *Constraint
	errors   cog.BuildErrors
}

func NewConstraintBuilder() *ConstraintBuilder {
	resource := NewConstraint()
	builder := &ConstraintBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ConstraintBuilder) Build() (Constraint, error) {
	if err := builder.internal.Validate(); err != nil {
		return Constraint{}, err
	}

	if len(builder.errors) > 0 {
		return Constraint{}, cog.MakeBuildErrors("canvas.constraint", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ConstraintBuilder) RecordError(path string, err error) *ConstraintBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ConstraintBuilder) Horizontal(horizontal HorizontalConstraint) *ConstraintBuilder {
	builder.internal.Horizontal = &horizontal

	return builder
}

func (builder *ConstraintBuilder) Vertical(vertical VerticalConstraint) *ConstraintBuilder {
	builder.internal.Vertical = &vertical

	return builder
}
