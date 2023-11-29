// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package rolebinding

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CustomRoleRef] = (*CustomRoleRefBuilder)(nil)

type CustomRoleRefBuilder struct {
	internal *CustomRoleRef
	errors   map[string]cog.BuildErrors
}

func NewCustomRoleRefBuilder() *CustomRoleRefBuilder {
	resource := &CustomRoleRef{}
	builder := &CustomRoleRefBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Kind = "Role"

	return builder
}

func (builder *CustomRoleRefBuilder) Build() (CustomRoleRef, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("CustomRoleRef", err)...)
	}

	if len(errs) != 0 {
		return CustomRoleRef{}, errs
	}

	return *builder.internal, nil
}

func (builder *CustomRoleRefBuilder) Name(name string) *CustomRoleRefBuilder {
	builder.internal.Name = name

	return builder
}

func (builder *CustomRoleRefBuilder) applyDefaults() {
}
