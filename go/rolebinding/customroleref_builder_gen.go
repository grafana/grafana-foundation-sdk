// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package rolebinding

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CustomRoleRef] = (*CustomRoleRefBuilder)(nil)

type CustomRoleRefBuilder struct {
	internal *CustomRoleRef
	errors   cog.BuildErrors
}

func NewCustomRoleRefBuilder() *CustomRoleRefBuilder {
	resource := NewCustomRoleRef()
	builder := &CustomRoleRefBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "Role"

	return builder
}

func (builder *CustomRoleRefBuilder) Build() (CustomRoleRef, error) {
	if err := builder.internal.Validate(); err != nil {
		return CustomRoleRef{}, err
	}

	if len(builder.errors) > 0 {
		return CustomRoleRef{}, cog.MakeBuildErrors("rolebinding.customRoleRef", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *CustomRoleRefBuilder) Name(name string) *CustomRoleRefBuilder {
	builder.internal.Name = name

	return builder
}
