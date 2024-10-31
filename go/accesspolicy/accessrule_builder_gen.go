// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package accesspolicy

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AccessRule] = (*AccessRuleBuilder)(nil)

type AccessRuleBuilder struct {
	internal *AccessRule
	errors   map[string]cog.BuildErrors
}

func NewAccessRuleBuilder() *AccessRuleBuilder {
	resource := &AccessRule{}
	builder := &AccessRuleBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *AccessRuleBuilder) Build() (AccessRule, error) {
	if err := builder.internal.Validate(); err != nil {
		return AccessRule{}, err
	}

	return *builder.internal, nil
}

// The kind this rule applies to (dashboards, alert, etc)
func (builder *AccessRuleBuilder) Kind(kind string) *AccessRuleBuilder {
	builder.internal.Kind = kind

	return builder
}

// READ, WRITE, CREATE, DELETE, ...
// should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"
func (builder *AccessRuleBuilder) Verb(verb string) *AccessRuleBuilder {
	builder.internal.Verb = verb

	return builder
}

// Specific sub-elements like "alert.rules" or "dashboard.permissions"????
func (builder *AccessRuleBuilder) Target(target string) *AccessRuleBuilder {
	builder.internal.Target = &target

	return builder
}

func (builder *AccessRuleBuilder) applyDefaults() {
	builder.Kind("*")
}
