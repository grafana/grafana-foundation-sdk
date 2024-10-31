// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package role

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Role] = (*RoleBuilder)(nil)

type RoleBuilder struct {
	internal *Role
	errors   map[string]cog.BuildErrors
}

func NewRoleBuilder() *RoleBuilder {
	resource := &Role{}
	builder := &RoleBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *RoleBuilder) Build() (Role, error) {
	if err := builder.internal.Validate(); err != nil {
		return Role{}, err
	}

	return *builder.internal, nil
}

// The role identifier `managed:builtins:editor:permissions`
func (builder *RoleBuilder) Name(name string) *RoleBuilder {
	builder.internal.Name = name

	return builder
}

// Optional display
func (builder *RoleBuilder) DisplayName(displayName string) *RoleBuilder {
	builder.internal.DisplayName = &displayName

	return builder
}

// Name of the team.
func (builder *RoleBuilder) GroupName(groupName string) *RoleBuilder {
	builder.internal.GroupName = &groupName

	return builder
}

// Role description
func (builder *RoleBuilder) Description(description string) *RoleBuilder {
	builder.internal.Description = &description

	return builder
}

// Do not show this role
func (builder *RoleBuilder) Hidden(hidden bool) *RoleBuilder {
	builder.internal.Hidden = hidden

	return builder
}

func (builder *RoleBuilder) applyDefaults() {
	builder.Hidden(false)
}
