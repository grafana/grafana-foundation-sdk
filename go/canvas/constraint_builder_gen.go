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
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Constraint", err)...)
	}

	if len(errs) != 0 {
		return Constraint{}, errs
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
