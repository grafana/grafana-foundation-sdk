// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Constraint] = (*ConstraintBuilder)(nil)

type ConstraintBuilder struct {
	internal *Constraint
	errors   map[string]cog.BuildErrors
}

func NewConstraintBuilder() *ConstraintBuilder {
	resource := &Constraint{}
	builder := &ConstraintBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ConstraintBuilder) Build() (Constraint, error) {
	if err := builder.internal.Validate(); err != nil {
		return Constraint{}, err
	}

	return *builder.internal, nil
}

func (builder *ConstraintBuilder) Horizontal(horizontal HorizontalConstraint) *ConstraintBuilder {
	builder.internal.Horizontal = &horizontal

	return builder
}

func (builder *ConstraintBuilder) Vertical(vertical VerticalConstraint) *ConstraintBuilder {
	builder.internal.Vertical = &vertical

	return builder
}

func (builder *ConstraintBuilder) applyDefaults() {
}
