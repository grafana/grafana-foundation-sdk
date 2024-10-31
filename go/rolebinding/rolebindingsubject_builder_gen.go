// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package rolebinding

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RoleBindingSubject] = (*RoleBindingSubjectBuilder)(nil)

type RoleBindingSubjectBuilder struct {
	internal *RoleBindingSubject
	errors   map[string]cog.BuildErrors
}

func NewRoleBindingSubjectBuilder() *RoleBindingSubjectBuilder {
	resource := &RoleBindingSubject{}
	builder := &RoleBindingSubjectBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *RoleBindingSubjectBuilder) Build() (RoleBindingSubject, error) {
	if err := builder.internal.Validate(); err != nil {
		return RoleBindingSubject{}, err
	}

	return *builder.internal, nil
}

func (builder *RoleBindingSubjectBuilder) Kind(kind RoleBindingSubjectKind) *RoleBindingSubjectBuilder {
	builder.internal.Kind = kind

	return builder
}

// The team/user identifier name
func (builder *RoleBindingSubjectBuilder) Name(name string) *RoleBindingSubjectBuilder {
	builder.internal.Name = name

	return builder
}

func (builder *RoleBindingSubjectBuilder) applyDefaults() {
}
