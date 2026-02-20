// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[LogGroup] = (*LogGroupBuilder)(nil)

type LogGroupBuilder struct {
	internal *LogGroup
	errors   cog.BuildErrors
}

func NewLogGroupBuilder() *LogGroupBuilder {
	resource := NewLogGroup()
	builder := &LogGroupBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *LogGroupBuilder) Build() (LogGroup, error) {
	if err := builder.internal.Validate(); err != nil {
		return LogGroup{}, err
	}

	if len(builder.errors) > 0 {
		return LogGroup{}, cog.MakeBuildErrors("cloudwatch.logGroup", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *LogGroupBuilder) RecordError(path string, err error) *LogGroupBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
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
