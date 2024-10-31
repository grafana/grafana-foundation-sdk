// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RuleGroup] = (*RuleGroupBuilder)(nil)

type RuleGroupBuilder struct {
	internal *RuleGroup
	errors   map[string]cog.BuildErrors
}

func NewRuleGroupBuilder(title string) *RuleGroupBuilder {
	resource := &RuleGroup{}
	builder := &RuleGroupBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Title = &title

	return builder
}

func (builder *RuleGroupBuilder) Build() (RuleGroup, error) {
	if err := builder.internal.Validate(); err != nil {
		return RuleGroup{}, err
	}

	return *builder.internal, nil
}

func (builder *RuleGroupBuilder) FolderUid(folderUid string) *RuleGroupBuilder {
	builder.internal.FolderUid = &folderUid

	return builder
}

// The interval, in seconds, at which all rules in the group are evaluated.
// If a group contains many rules, the rules are evaluated sequentially.
func (builder *RuleGroupBuilder) Interval(interval Duration) *RuleGroupBuilder {
	builder.internal.Interval = &interval

	return builder
}

func (builder *RuleGroupBuilder) Rules(rules []cog.Builder[Rule]) *RuleGroupBuilder {
	rulesResources := make([]Rule, 0, len(rules))
	for _, r1 := range rules {
		rulesDepth1, err := r1.Build()
		if err != nil {
			builder.errors["rules"] = err.(cog.BuildErrors)
			return builder
		}
		rulesResources = append(rulesResources, rulesDepth1)
	}
	builder.internal.Rules = rulesResources

	return builder
}

func (builder *RuleGroupBuilder) WithRule(rule cog.Builder[Rule]) *RuleGroupBuilder {
	ruleResource, err := rule.Build()
	if err != nil {
		builder.errors["rules"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Rules = append(builder.internal.Rules, ruleResource)

	return builder
}

func (builder *RuleGroupBuilder) Title(title string) *RuleGroupBuilder {
	builder.internal.Title = &title

	return builder
}

func (builder *RuleGroupBuilder) applyDefaults() {
}
