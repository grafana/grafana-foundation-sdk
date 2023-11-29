// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package rolebinding

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuiltinRoleRef] = (*BuiltinRoleRefBuilder)(nil)

type BuiltinRoleRefBuilder struct {
	internal *BuiltinRoleRef
	errors   map[string]cog.BuildErrors
}

func NewBuiltinRoleRefBuilder() *BuiltinRoleRefBuilder {
	resource := &BuiltinRoleRef{}
	builder := &BuiltinRoleRefBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Kind = "BuiltinRole"

	return builder
}

func (builder *BuiltinRoleRefBuilder) Build() (BuiltinRoleRef, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("BuiltinRoleRef", err)...)
	}

	if len(errs) != 0 {
		return BuiltinRoleRef{}, errs
	}

	return *builder.internal, nil
}

func (builder *BuiltinRoleRefBuilder) Name(name BuiltinRoleRefName) *BuiltinRoleRefBuilder {
	builder.internal.Name = name

	return builder
}

func (builder *BuiltinRoleRefBuilder) applyDefaults() {
}
