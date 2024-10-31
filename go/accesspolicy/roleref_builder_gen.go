// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package accesspolicy

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RoleRef] = (*RoleRefBuilder)(nil)

type RoleRefBuilder struct {
	internal *RoleRef
	errors   map[string]cog.BuildErrors
}

func NewRoleRefBuilder() *RoleRefBuilder {
	resource := &RoleRef{}
	builder := &RoleRefBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *RoleRefBuilder) Build() (RoleRef, error) {
	if err := builder.internal.Validate(); err != nil {
		return RoleRef{}, err
	}

	return *builder.internal, nil
}

// Policies can apply to roles, teams, or users
// Applying policies to individual users is supported, but discouraged
func (builder *RoleRefBuilder) Kind(kind RoleRefKind) *RoleRefBuilder {
	builder.internal.Kind = kind

	return builder
}

func (builder *RoleRefBuilder) Name(name string) *RoleRefBuilder {
	builder.internal.Name = name

	return builder
}

func (builder *RoleRefBuilder) Xname(xname string) *RoleRefBuilder {
	builder.internal.Xname = xname

	return builder
}

func (builder *RoleRefBuilder) applyDefaults() {
}
