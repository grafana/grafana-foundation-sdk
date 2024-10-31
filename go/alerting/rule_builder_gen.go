// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"time"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Rule] = (*RuleBuilder)(nil)

type RuleBuilder struct {
	internal *Rule
	errors   map[string]cog.BuildErrors
}

func NewRuleBuilder(title string) *RuleBuilder {
	resource := &Rule{}
	builder := &RuleBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Title = title

	return builder
}

func (builder *RuleBuilder) Build() (Rule, error) {
	if err := builder.internal.Validate(); err != nil {
		return Rule{}, err
	}

	return *builder.internal, nil
}

func (builder *RuleBuilder) Annotations(annotations map[string]string) *RuleBuilder {
	builder.internal.Annotations = annotations

	return builder
}

func (builder *RuleBuilder) Condition(condition string) *RuleBuilder {
	builder.internal.Condition = condition

	return builder
}

func (builder *RuleBuilder) Queries(data []cog.Builder[Query]) *RuleBuilder {
	dataResources := make([]Query, 0, len(data))
	for _, r1 := range data {
		dataDepth1, err := r1.Build()
		if err != nil {
			builder.errors["data"] = err.(cog.BuildErrors)
			return builder
		}
		dataResources = append(dataResources, dataDepth1)
	}
	builder.internal.Data = dataResources

	return builder
}

func (builder *RuleBuilder) WithQuery(data cog.Builder[Query]) *RuleBuilder {
	dataResource, err := data.Build()
	if err != nil {
		builder.errors["data"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Data = append(builder.internal.Data, dataResource)

	return builder
}

func (builder *RuleBuilder) ExecErrState(execErrState RuleExecErrState) *RuleBuilder {
	builder.internal.ExecErrState = execErrState

	return builder
}

func (builder *RuleBuilder) FolderUID(folderUID string) *RuleBuilder {
	builder.internal.FolderUID = folderUID

	return builder
}

// The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.
// Before this time has elapsed, the rule is only considered to be Pending.
func (builder *RuleBuilder) For(forArg string) *RuleBuilder {
	builder.internal.For = forArg

	return builder
}

func (builder *RuleBuilder) Id(id int64) *RuleBuilder {
	builder.internal.Id = &id

	return builder
}

func (builder *RuleBuilder) IsPaused(isPaused bool) *RuleBuilder {
	builder.internal.IsPaused = &isPaused

	return builder
}

func (builder *RuleBuilder) Labels(labels map[string]string) *RuleBuilder {
	builder.internal.Labels = labels

	return builder
}

func (builder *RuleBuilder) NoDataState(noDataState RuleNoDataState) *RuleBuilder {
	builder.internal.NoDataState = noDataState

	return builder
}

func (builder *RuleBuilder) OrgID(orgID int64) *RuleBuilder {
	builder.internal.OrgID = orgID

	return builder
}

func (builder *RuleBuilder) Provenance(provenance Provenance) *RuleBuilder {
	builder.internal.Provenance = &provenance

	return builder
}

func (builder *RuleBuilder) RuleGroup(ruleGroup string) *RuleBuilder {
	builder.internal.RuleGroup = ruleGroup

	return builder
}

func (builder *RuleBuilder) Title(title string) *RuleBuilder {
	builder.internal.Title = title

	return builder
}

func (builder *RuleBuilder) Uid(uid string) *RuleBuilder {
	builder.internal.Uid = &uid

	return builder
}

func (builder *RuleBuilder) Updated(updated time.Time) *RuleBuilder {
	builder.internal.Updated = &updated

	return builder
}

func (builder *RuleBuilder) applyDefaults() {
}
