// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[LogGroup] = (*LogGroupBuilder)(nil)

type LogGroupBuilder struct {
	internal *LogGroup
	errors   map[string]cog.BuildErrors
}

func NewLogGroupBuilder() *LogGroupBuilder {
	resource := &LogGroup{}
	builder := &LogGroupBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *LogGroupBuilder) Build() (LogGroup, error) {
	if err := builder.internal.Validate(); err != nil {
		return LogGroup{}, err
	}

	return *builder.internal, nil
}

// ARN of the log group
func (builder *LogGroupBuilder) Arn(arn string) *LogGroupBuilder {
	builder.internal.Arn = arn

	return builder
}

// Name of the log group
func (builder *LogGroupBuilder) Name(name string) *LogGroupBuilder {
	builder.internal.Name = name

	return builder
}

// AccountId of the log group
func (builder *LogGroupBuilder) AccountId(accountId string) *LogGroupBuilder {
	builder.internal.AccountId = &accountId

	return builder
}

// Label of the log group
func (builder *LogGroupBuilder) AccountLabel(accountLabel string) *LogGroupBuilder {
	builder.internal.AccountLabel = &accountLabel

	return builder
}

func (builder *LogGroupBuilder) applyDefaults() {
}
