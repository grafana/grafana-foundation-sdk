// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package rolebinding

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuiltinRoleRef] = (*BuiltinRoleRefBuilder)(nil)

type BuiltinRoleRefBuilder struct {
	internal *BuiltinRoleRef
	errors   cog.BuildErrors
}

func NewBuiltinRoleRefBuilder() *BuiltinRoleRefBuilder {
	resource := NewBuiltinRoleRef()
	builder := &BuiltinRoleRefBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "BuiltinRole"

	return builder
}

func (builder *BuiltinRoleRefBuilder) Build() (BuiltinRoleRef, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuiltinRoleRef{}, err
	}

	if len(builder.errors) > 0 {
		return BuiltinRoleRef{}, cog.MakeBuildErrors("rolebinding.builtinRoleRef", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *BuiltinRoleRefBuilder) Name(name BuiltinRoleRefName) *BuiltinRoleRefBuilder {
	builder.internal.Name = name

	return builder
}
