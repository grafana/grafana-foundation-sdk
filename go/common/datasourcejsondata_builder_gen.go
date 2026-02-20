// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DataSourceJsonData] = (*DataSourceJsonDataBuilder)(nil)

// TODO docs
type DataSourceJsonDataBuilder struct {
	internal *DataSourceJsonData
	errors   cog.BuildErrors
}

func NewDataSourceJsonDataBuilder() *DataSourceJsonDataBuilder {
	resource := NewDataSourceJsonData()
	builder := &DataSourceJsonDataBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *DataSourceJsonDataBuilder) Build() (DataSourceJsonData, error) {
	if err := builder.internal.Validate(); err != nil {
		return DataSourceJsonData{}, err
	}

	if len(builder.errors) > 0 {
		return DataSourceJsonData{}, cog.MakeBuildErrors("common.dataSourceJsonData", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *DataSourceJsonDataBuilder) RecordError(path string, err error) *DataSourceJsonDataBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *DataSourceJsonDataBuilder) AuthType(authType string) *DataSourceJsonDataBuilder {
	builder.internal.AuthType = &authType

	return builder
}

func (builder *DataSourceJsonDataBuilder) DefaultRegion(defaultRegion string) *DataSourceJsonDataBuilder {
	builder.internal.DefaultRegion = &defaultRegion

	return builder
}

func (builder *DataSourceJsonDataBuilder) Profile(profile string) *DataSourceJsonDataBuilder {
	builder.internal.Profile = &profile

	return builder
}

func (builder *DataSourceJsonDataBuilder) ManageAlerts(manageAlerts bool) *DataSourceJsonDataBuilder {
	builder.internal.ManageAlerts = &manageAlerts

	return builder
}

func (builder *DataSourceJsonDataBuilder) AlertmanagerUid(alertmanagerUid string) *DataSourceJsonDataBuilder {
	builder.internal.AlertmanagerUid = &alertmanagerUid

	return builder
}
